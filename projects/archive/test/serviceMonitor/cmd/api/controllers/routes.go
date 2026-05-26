package controllers

import (
	"fmt"
	"net/http"

	"github.com/josepht96/learning/go/serviceMonitor/cmd/api/middleware"
)

//RegisterControllers lists routes
func RegisterControllers() {
	fmt.Println("Generating handlers...")
	http.Handle("/allergyintolerance", middleware.Handler(newAlrgIntolController()))
	http.Handle("/allergyintolerance/", middleware.Handler(newAlrgIntolController()))
	http.Handle("/condition", middleware.Handler(newConditionController()))
	http.Handle("/condition/", middleware.Handler(newConditionController()))

}
