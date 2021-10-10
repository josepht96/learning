package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	registerControllers()
	fmt.Println("Listening at: http://localhost:4040/")
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		err := http.ListenAndServe(":4040", nil)
		if err != nil {
			log.Fatal(err)
		}
		cancel()
	}()
	<-ctx.Done()
	fmt.Println("Shutting down the Sender...")

}
func registerControllers() {
	fmt.Println("Generating handlers...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTime(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	})
}
func getTime(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format("2006-01-02 15:04:05.000000000")
	tString := fmt.Sprintf("container a: %s\n", t)
	fmt.Printf("Sending response at: %s\n", tString)
	w.Write([]byte(tString))
}
