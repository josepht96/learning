package controllers

import (
	"net/http"
	"regexp"
)

type allergyintoleranceC struct {
	ID *regexp.Regexp
}

func (api *allergyintoleranceC) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/allergyintolerance" {
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("allergyintolerance route hit"))
		}
	} else if r.URL.Path == "/allergyintolerance/" {
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("allergyintolerance/ route hit"))
		}
	}

}
func newAlrgIntolController() *allergyintoleranceC {
	return &allergyintoleranceC{
		ID: regexp.MustCompile(`^/allergyintolerance/(\d+)/?`),
	}
}
