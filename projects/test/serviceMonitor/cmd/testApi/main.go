package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	registerControllers()
	fmt.Println("Listening at: http://localhost:4041/endpoint")
	err := http.ListenAndServe(":4041", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func registerControllers() {
	fmt.Println("Generating handlers...")
	http.HandleFunc("/endpoint", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("test endpoint reached\n"))
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	})
}
