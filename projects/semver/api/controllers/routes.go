package controllers

import (
	"net/http"

	"github.com/josepht96/learning/projects/semver/api/middleware"
)

//RegisterControllers lists routes
func RegisterControllers() {
	http.Handle("/", middleware.Handler(newDefaultHandler()))
	http.Handle("/user/", middleware.Handler(newUserHandler()))
	http.Handle("/project/", middleware.Handler(newProjectHandler()))
	http.Handle("/projects", middleware.Handler(newProjectsHandler()))
}
