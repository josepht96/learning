package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

var login_url = "http://localhost:8080/auth/"
var targetToken = &oauth2.Token{}

type TargetHandler struct {
}

func (h *TargetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/target/" {
		switch r.Method {
		case http.MethodGet:
			h.response(w, r)
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

func (h *TargetHandler) response(w http.ResponseWriter, r *http.Request) {
	if targetToken.AccessToken == "" {
		http.Redirect(w, r, login_url, http.StatusTemporaryRedirect)
	} else {
		e := json.NewEncoder(w)
		e.SetIndent("", "  ")
		e.Encode(targetToken)
	}
}

func (h *TargetHandler) notFound(w http.ResponseWriter, r *http.Request) {
	dt := time.Now()
	data := response{
		Status:     "NotFound",
		StatusCode: 404,
		Body: message{
			Message: "404 - page not found",
			Time:    dt.String(),
		},
	}
	fmt.Println(data)
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(data)
}

func newTargetHandler() *TargetHandler {
	return &TargetHandler{}
}
