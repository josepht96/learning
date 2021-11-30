package controllers

import (
	//"encoding/json"
	//"fmt"
	//"io/ioutil"
	"net/http"
	//"strconv"
	"github.com/josepht96/learning/projects/semver/api/models"
)

type projectsHandler struct {
}

func (p *projectsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/projects" {
		switch r.Method {
		case http.MethodGet:
			getProjects(w, r)
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

func getProjects(w http.ResponseWriter, r *http.Request) {
	models.GetProjects()
	w.WriteHeader(http.StatusOK)
}

func newProjectsHandler() *projectsHandler {
	return &projectsHandler{
	}
}
