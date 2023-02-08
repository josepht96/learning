package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
)

var endpoint = &oauth2.Endpoint{
	AuthURL:  os.Getenv("AUTH_URL"),
	TokenURL: os.Getenv("TOKEN_URL"),
}
var config = &oauth2.Config{
	RedirectURL:  os.Getenv("REDIRECT_URL"),
	ClientID:     os.Getenv("CLIENT_ID"),
	ClientSecret: os.Getenv("CLIENT_SECRET"),
	Endpoint:     *endpoint,
}

type authHandler struct {
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/auth/" {
		switch r.Method {
		case http.MethodGet:
			h.login(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else if r.URL.Path == "/auth/callback/" {
		switch r.Method {
		case http.MethodGet:
			h.callback(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		switch r.Method {
		case http.MethodGet:
			h.notFound(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (h *authHandler) login(w http.ResponseWriter, r *http.Request) {
	url := config.AuthCodeURL("test-state")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *authHandler) callback(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	state := r.Form.Get("state")
	if state != "test-state" {
		http.Error(w, "State invalid", http.StatusBadRequest)
		return
	}
	code := r.Form.Get("code")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}

	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	targetToken = token
	//pkg.ParseJWT(token.AccessToken)
	http.Redirect(w, r, "http://localhost:8080/target/", http.StatusPermanentRedirect)
}

func (h *authHandler) notFound(w http.ResponseWriter, r *http.Request) {
	dt := time.Now()
	data := response{
		Status:     "NotFound",
		StatusCode: 404,
		Body: message{
			Message: "404 - auth page not found",
			Time:    dt.String(),
		},
	}
	fmt.Println(data)
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(data)
}

func newAuthHandler() *authHandler {
	return &authHandler{}
}
