package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWT secret key - In production, use environment variable
const jwtSecret = "your-secret-key-change-in-production"

// CSRF token store (in production, use Redis or similar)
var csrfTokens = struct {
	sync.RWMutex
	tokens map[string]time.Time
}{tokens: make(map[string]time.Time)}

// Hardcoded credentials for demo purposes
const (
	validUsername = "admin"
	validPassword = "password"
)

// Claims structure for JWT
type Claims struct {
	Username  string `json:"username"`
	TokenType string `json:"token_type"` // "access" or "refresh"
	jwt.RegisteredClaims
}

// Login request structure
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login response structure
type LoginResponse struct {
	CsrfToken string `json:"csrf_token"`
	Message   string `json:"message"`
}

// Refresh response structure
type RefreshResponse struct {
	Message string `json:"message"`
}

// Protected response structure
type ProtectedResponse struct {
	Message  string `json:"message"`
	Username string `json:"username"`
	IssuedAt string `json:"issued_at"`
	ExpiresAt string `json:"expires_at"`
}

// Error response structure
type ErrorResponse struct {
	Error string `json:"error"`
}

// Generate a random CSRF token
func generateCSRFToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	token := base64.URLEncoding.EncodeToString(b)

	// Store token with expiration
	csrfTokens.Lock()
	csrfTokens.tokens[token] = time.Now().Add(24 * time.Hour)
	csrfTokens.Unlock()

	// Clean up expired tokens
	go cleanupExpiredCSRFTokens()

	return token, nil
}

// Validate CSRF token
func validateCSRFToken(token string) bool {
	csrfTokens.RLock()
	expiry, exists := csrfTokens.tokens[token]
	csrfTokens.RUnlock()

	if !exists {
		return false
	}

	if time.Now().After(expiry) {
		// Token expired, remove it
		csrfTokens.Lock()
		delete(csrfTokens.tokens, token)
		csrfTokens.Unlock()
		return false
	}

	return true
}

// Clean up expired CSRF tokens
func cleanupExpiredCSRFTokens() {
	csrfTokens.Lock()
	defer csrfTokens.Unlock()

	now := time.Now()
	for token, expiry := range csrfTokens.tokens {
		if now.After(expiry) {
			delete(csrfTokens.tokens, token)
		}
	}
}

// Set security headers including CORS and CSP
func setSecurityHeaders(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")

	// Whitelist specific origins (localhost for development)
	allowedOrigins := map[string]bool{
		"http://localhost:8080": true,
	}

	if allowedOrigins[origin] {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}

	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")

	// Content Security Policy
	w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'")

	// Additional security headers
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
}

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Serve index.html at root
	http.HandleFunc("/", serveIndex)

	// API endpoints
	http.HandleFunc("/api/login", loginHandler)
	http.HandleFunc("/api/refresh", csrfMiddleware(refreshHandler))
	http.HandleFunc("/api/protected", csrfMiddleware(jwtMiddleware(protectedHandler)))

	// Start server
	port := ":8080"
	fmt.Printf("Server starting on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// serveIndex serves the login page
func serveIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "./static/index.html")
		return
	}
	http.NotFound(w, r)
}

// loginHandler handles login requests
func loginHandler(w http.ResponseWriter, r *http.Request) {
	setSecurityHeaders(w, r)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var loginReq LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate credentials
	if loginReq.Username != validUsername || loginReq.Password != validPassword {
		sendError(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Create access token (30 seconds)
	accessTokenExp := time.Now().Add(30 * time.Second)
	accessClaims := &Claims{
		Username:  loginReq.Username,
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessTokenExp),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString([]byte(jwtSecret))
	if err != nil {
		sendError(w, "Error generating access token", http.StatusInternalServerError)
		return
	}

	// Create refresh token (7 days)
	refreshTokenExp := time.Now().Add(7 * 24 * time.Hour)
	refreshClaims := &Claims{
		Username:  loginReq.Username,
		TokenType: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshTokenExp),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(jwtSecret))
	if err != nil {
		sendError(w, "Error generating refresh token", http.StatusInternalServerError)
		return
	}

	// Set access token as HttpOnly cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessTokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // Set to true in production with HTTPS
		SameSite: http.SameSiteStrictMode,
		MaxAge:   30, // 30 seconds
	})

	// Set refresh token as HttpOnly cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshTokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // Set to true in production with HTTPS
		SameSite: http.SameSiteStrictMode,
		MaxAge:   7 * 24 * 60 * 60, // 7 days
	})

	// Generate CSRF token
	csrfToken, err := generateCSRFToken()
	if err != nil {
		sendError(w, "Error generating CSRF token", http.StatusInternalServerError)
		return
	}

	// Send response with CSRF token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse{
		CsrfToken: csrfToken,
		Message:   "Login successful",
	})
}

// refreshHandler handles token refresh requests
func refreshHandler(w http.ResponseWriter, r *http.Request) {
	setSecurityHeaders(w, r)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get refresh token from cookie
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		sendError(w, "Refresh token not found", http.StatusUnauthorized)
		return
	}

	// Parse and validate refresh token
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		sendError(w, "Invalid or expired refresh token", http.StatusUnauthorized)
		return
	}

	// Verify this is a refresh token
	if claims.TokenType != "refresh" {
		sendError(w, "Invalid token type", http.StatusUnauthorized)
		return
	}

	// Generate new access token
	accessTokenExp := time.Now().Add(30 * time.Second)
	accessClaims := &Claims{
		Username:  claims.Username,
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessTokenExp),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString([]byte(jwtSecret))
	if err != nil {
		sendError(w, "Error generating access token", http.StatusInternalServerError)
		return
	}

	// Set new access token as HttpOnly cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessTokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // Set to true in production with HTTPS
		SameSite: http.SameSiteStrictMode,
		MaxAge:   30, // 30 seconds
	})

	// Send success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(RefreshResponse{
		Message: "Token refreshed successfully",
	})
}

// csrfMiddleware validates CSRF token
func csrfMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setSecurityHeaders(w, r)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Get CSRF token from header
		csrfToken := r.Header.Get("X-CSRF-Token")
		if csrfToken == "" {
			sendError(w, "CSRF token required", http.StatusForbidden)
			return
		}

		// Validate CSRF token
		if !validateCSRFToken(csrfToken) {
			sendError(w, "Invalid or expired CSRF token", http.StatusForbidden)
			return
		}

		next(w, r)
	}
}

// jwtMiddleware validates JWT token
func jwtMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setSecurityHeaders(w, r)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Get token from cookie
		cookie, err := r.Cookie("access_token")
		if err != nil {
			sendError(w, "Access token not found", http.StatusUnauthorized)
			return
		}

		// Parse and validate token
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			sendError(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Verify this is an access token
		if claims.TokenType != "access" {
			sendError(w, "Invalid token type", http.StatusUnauthorized)
			return
		}

		// Store claims in request context for handler to use
		r = r.WithContext(r.Context())
		r.Header.Set("X-Username", claims.Username)
		r.Header.Set("X-Token-String", cookie.Value)

		next(w, r)
	}
}

// protectedHandler handles protected endpoint
func protectedHandler(w http.ResponseWriter, r *http.Request) {
	// Get username from middleware
	username := r.Header.Get("X-Username")

	// Parse token to get claims for response
	tokenString := r.Header.Get("X-Token-String")

	claims := &Claims{}
	jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	response := ProtectedResponse{
		Message:   fmt.Sprintf("Hello, %s! You have successfully accessed the protected endpoint.", username),
		Username:  username,
		IssuedAt:  claims.IssuedAt.Time.Format(time.RFC3339),
		ExpiresAt: claims.ExpiresAt.Time.Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// sendError sends a JSON error response
func sendError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}
