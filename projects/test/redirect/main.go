package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var port = 8080

func main() {
	// start and end page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/ reached")
		w.Write([]byte("hello!"))
	})
	// login page redirects to auth provider
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/login reached")
		http.Redirect(w, r, "http://localhost:8080/auth-provider", http.StatusTemporaryRedirect)
	})
	// auth provider redirects to callback
	http.HandleFunc("/auth-provider", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/auth-provider reached")
		time.Sleep(5 * time.Second)
		http.Redirect(w, r, "http://localhost:8080/callback", http.StatusTemporaryRedirect)
	})
	//callback redirects back to root
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/callback reached")
		http.Redirect(w, r, "http://localhost:8080/", http.StatusPermanentRedirect)
	})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		log.Printf("server is listening at http://localhost:%d", port)
		http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	}()
	wg.Wait()
}
