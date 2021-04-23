package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/josepht96/learning/go/goproject/controllers"
)

func middleWareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Before handler; middleware startup")
		start := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Printf("Middleware finished; %v\n", time.Since(start))
	})
}
func main() {
	controllers.RegisterControllers()
	fmt.Println("Listening at: localhost:5000")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
