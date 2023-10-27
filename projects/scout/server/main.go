package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Response struct {
	Status     string  `json:"status"`
	StatusCode int     `json:"statusCode"`
	Body       Message `json:"body"`
}

type Message struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

// handlers handles the endpoint for remote scout instances
func handlers() {
	// http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		dt := time.Now()
		data := Response{
			Status:     "OK",
			StatusCode: 200,
			Body: Message{
				Message: "connected to scout",
				Time:    dt.String(),
			},
		}
		json.NewEncoder(w).Encode(data)
		log.Printf("host: %s", r.Host)
		log.Printf("requestURI: %s", r.RequestURI)
		log.Printf("contentLength: %d", r.ContentLength)
		log.Printf("method: %s", r.Method)
		log.Printf("protocol: %s", r.Proto)
		log.Printf("remoteAddr: %s", r.RemoteAddr)

		log.Println("request headers:")
		for key, values := range r.Header {
			for _, value := range values {
				log.Printf("\t%s: %s", key, value)
			}
		}
		log.Println("---------------------------------------------------------")
	})
}

// main initializes handlers, invokes the probe mechanism, and starts a server
func main() {
	go handlers()
	var port = 8080
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		log.Printf("server is listening at http://localhost:%d", port)
		http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	}()
	wg.Wait()
}
