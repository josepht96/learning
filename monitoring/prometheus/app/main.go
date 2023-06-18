package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
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

var port = 8081

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			for name, values := range r.Header {
				// Loop over all values for the name.
				for _, value := range values {
					fmt.Println(name, value)
				}
			}
			(w).Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			dt := time.Now()
			data := Response{
				Status:     "OK",
				StatusCode: 200,
				Body: Message{
					Message: "hello from hello-server",
					Time:    dt.String(),
				},
			}
			fmt.Println(data)
			json.NewEncoder(w).Encode(data)
			fmt.Println("---------------------------------------------------------")
		case http.MethodPost:
			w.WriteHeader(http.StatusNotFound)
		case http.MethodPut:
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		log.Printf("server is listening at http://localhost:%d", port)
		http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	}()
	wg.Wait()
}
