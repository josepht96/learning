package controllers

import (
	"fmt"
	"net/http"
)

type defaultHandler struct {
}

func (d *defaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		switch r.Method {
		case http.MethodGet:
			defaultResponse(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		switch r.Method {
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}

}

func defaultResponse(w http.ResponseWriter, r *http.Request) {
	rMsg := fmt.Sprintf("Accessing the default route...\nRoutes are:\n /inflation\n")
	w.Write([]byte(rMsg))
}

func newDefaultHandler() *defaultHandler {
	return &defaultHandler{
	}
}