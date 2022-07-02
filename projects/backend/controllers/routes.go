package controllers

import (
	"net/http"

	"github.com/josepht96/learning/projects/backend/middleware"
)

//RegisterControllers lists routes
func RegisterControllers() {
	http.Handle("/", middleware.Handler(newDefaultHandler()))
	http.Handle("/inflation", middleware.Handler(newInflationHandler()))
}
