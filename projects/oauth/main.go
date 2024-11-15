package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	oauthURL     = "https://auth-service.com/oauth/token" // Replace with actual OAuth token URL
	clientID     = "your-client-id"
	clientSecret = "your-client-secret"
	username     = "your-username"
	password     = "your-password"
)

// TokenResponse represents the JSON structure of the token response
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// getOAuthToken retrieves the OAuth token using the password grant type
func getOAuthToken() (string, error) {
	data := map[string]string{
		"grant_type":    "password",
		"client_id":     clientID,
		"client_secret": clientSecret,
		"username":      username,
		"password":      password,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to encode JSON: %v", err)
	}

	resp, err := http.Post(oauthURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("failed to retrieve token: %d, %s", resp.StatusCode, string(body))
	}

	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	return tokenResp.AccessToken, nil
}

// makeAuthorizedRequest sends a GET request to the test endpoint with the token
func makeAuthorizedRequest(token string) error {
	req, err := http.NewRequest("GET", "https://my-service.com/test", nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	fmt.Printf("Response Code: %d\n", resp.StatusCode)
	return nil
}

func main() {
	token, err := getOAuthToken()
	if err != nil {
		fmt.Printf("Error retrieving token: %v\n", err)
		return
	}

	if err := makeAuthorizedRequest(token); err != nil {
		fmt.Printf("Error making authorized request: %v\n", err)
	}
}
