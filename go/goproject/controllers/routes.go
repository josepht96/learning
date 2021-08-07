package controllers

import (
	"net/http"

	"github.com/josepht96/learning/go/goproject/foundation/middleware"
)

//RegisterControllers lists routes
func RegisterControllers() {
	http.Handle("/foo", middleware.Handler(newFooController()))
	http.Handle("/foo/", middleware.Handler(newFooController()))
}
