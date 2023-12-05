package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

type Response struct {
	Status     string  `json:"status"`
	StatusCode int     `json:"statusCode"`
	Body       Message `json:"body"`
}

type Message struct {
	Message string `json:"message"`
}

var port = 8080
var targetService = os.Getenv("TARGET_SERVICE")

func handleRoot(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		log.Printf("%s:%s", k, v)
	}
	req, _ := http.NewRequest("GET", fmt.Sprintf("http://%s", targetService), nil)
	if _, ok := r.Header["X-Request-Id"]; ok {
		log.Printf("X-Request-Id header == %s\n", r.Header["X-Request-Id"][0])
		req.Header.Add("X-Request-Id", r.Header["X-Request-Id"][0])
		req.Header.Add("X-B3-Traceid", r.Header["X-B3-Traceid"][0])
		req.Header.Add("X-B3-Spanid", r.Header["X-B3-Spanid"][0])
		req.Header.Add("X-B3-Parentspanid", r.Header["X-B3-Parentspanid"][0])
		req.Header.Add("X-B3-Sampled", r.Header["X-B3-Sampled"][0])
	} else {
		log.Println("no X-Request-Id header")
	}

	httpClient := &http.Client{Transport: http.DefaultTransport}
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	log.Println(resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := Response{
		Status:     "OK",
		StatusCode: 200,
		Body: Message{
			Message: "connected to jaeger tracing test",
		},
	}
	json.NewEncoder(w).Encode(data)
}

// receive request, forward headers to next service
func main() {
	http.HandleFunc("/", handleRoot)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		log.Printf("server is listening at http://localhost:%d", port)
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
	wg.Wait()
}
