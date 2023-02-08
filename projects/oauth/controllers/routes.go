package controllers

import (
	"fmt"
	"net/http"

	"github.com/josepht96/learning/projects/oauth/middleware"
)

//RegisterControllers lists routes
func RegisterControllers() {
	fmt.Println("-- RegisterControllers --")
	http.Handle("/", middleware.Handler(newRootHandler()))
	http.Handle("/auth/", middleware.Handler(newAuthHandler()))
	http.Handle("/target/", middleware.Handler(newTargetHandler()))
}
