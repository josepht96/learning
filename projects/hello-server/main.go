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

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		dt := time.Now()
		data := Response{
			Status:     "OK",
			StatusCode: 200,
			Body: Message{
				Message: "hello from test-server",
				Time:    dt.String(),
			},
		}
		fmt.Println(data)
		json.NewEncoder(w).Encode(data)
	})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		log.Printf("server is listening at %s", "http://localhost:8080")
		http.ListenAndServe(":8080", nil)
	}()
	go func() {
		log.Printf("server is listening at %s", "http://localhost:8081")
		http.ListenAndServe(":8081", nil)
	}()
	wg.Wait()
}
