package controllers

import (
	"fmt"
	"net/http"
	"time"
)

func middleWareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Before handler; middleware startup")
		start := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Printf("Middleware finished; %v\n", time.Since(start))
	})
}
func RegisterControllers() {
	http.Handle("/foo", middleWareHandler(NewFooController()))
	http.Handle("/foo/", middleWareHandler(NewFooController()))
}
