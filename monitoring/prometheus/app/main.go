package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/yaml.v3"
)

var port = 8080

type Response struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type Endpoints struct {
	Endpoints map[string]Endpoint `yaml:"endpoints"`
}

type Endpoint struct {
	Url  string `yaml:"url"`
	Auth string `yaml:"auth"`
}

type rootHandler struct{}

// getData parses filename for endpoints
func getData(filename string) (*Endpoints, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	data := &Endpoints{}
	err = yaml.Unmarshal(buf, data)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %w", filename, err)
	}

	return data, err
}

// ServeHTTP handles requests to "/"
func (root *rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		(w).Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		data := Response{
			Status:     "OK",
			StatusCode: 200,
			Message:    "Navigate to /metrics",
		}
		json.NewEncoder(w).Encode(data)
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func createMetrics(data *Endpoints) *prometheus.CounterVec {
	counter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "endpoint_requests",
			Help: "A counter that tracks requests to endpoints",
		},
		[]string{"endpoint"},
	)
	prometheus.MustRegister(counter)

	return counter
}

// calls executes endpoint requests
func calls(data *Endpoints, counter prometheus.CounterVec) {
	for key, value := range data.Endpoints {
		fmt.Printf("Request: %s\n", key)
		resp, err := http.Get(value.Url)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(resp.StatusCode)
		counter.With(prometheus.Labels{"endpoint": fmt.Sprintf("%s", key)}).Inc()
	}
}

// main runs the program
func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/", &rootHandler{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		log.Printf("server is listening at http://localhost:%d", port)
		http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	}()
	data, err := getData("config.yaml")
	if err != nil {
		log.Fatal("Error getting yaml data")
	}
	counter := createMetrics(data)
	go calls(data, *counter)
	wg.Wait()
}
