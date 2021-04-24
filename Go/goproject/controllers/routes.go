package controllers

import (
	"net/http"

	"github.com/josepht96/learning/go/goproject/middleware"
)

//RegisterControllers lists routes
func RegisterControllers() {
	http.Handle("/foo", middleware.Handler(newFooController()))
	http.Handle("/foo/", middleware.Handler(newFooController()))
}
