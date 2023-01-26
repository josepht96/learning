package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type response struct {
	Status     string  `json:"status"`
	StatusCode int     `json:"statusCode"`
	Body       message `json:"body"`
}

type message struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

type rootHandler struct {
}

func (h *rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		switch r.Method {
		case http.MethodGet:
			h.defaultResponse(w, r)
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

func (h *rootHandler) defaultResponse(w http.ResponseWriter, r *http.Request) {
	dt := time.Now()
	data := response{
		Status:     "OK",
		StatusCode: 200,
		Body: message{
			Message: "hello from test-server",
			Time:    dt.String(),
		},
	}
	fmt.Println(data)
	json.NewEncoder(w).Encode(data)
}

func (h *rootHandler) notFound(w http.ResponseWriter, r *http.Request) {
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

func newRootHandler() *rootHandler {
	return &rootHandler{}
}
