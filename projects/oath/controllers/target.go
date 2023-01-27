package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/josepht96/learning/oauth/pkg"
	"golang.org/x/oauth2"
)

var login_url = "http://localhost:8080/auth/"
var targetToken = &oauth2.Token{}

type ResponseObject struct {
	Token      *oauth2.Token
	TargetBody string
}
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
		headers := []pkg.Header{}
		authHeader := pkg.Header{
			Key:   "Authorization",
			Value: fmt.Sprintf("Bearer %s", targetToken.AccessToken),
		}
		headers = append(headers, authHeader)
		body, err := pkg.Probe("http://localhost:30001", headers)
		resp := ResponseObject{
			Token:      targetToken,
			TargetBody: body,
		}
		e := json.NewEncoder(w)
		e.SetIndent("", "  ")
		e.Encode(resp)
		if err != nil {
			log.Fatal(err)
		}
		targetToken.AccessToken = ""
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
