package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := "hello from test-server"
		w.Write([]byte(response))
	})
	log.Printf("server is listening at %s", "localhost:8080\n")
	http.ListenAndServe(":8080", nil)
}
