package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

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
		data := Response{
			Status:     "OK",
			StatusCode: 200,
			Message:    "Navigate to /metrics",
		}
		json.NewEncoder(w).Encode(data)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func createMetrics() *prometheus.CounterVec {
	counter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "endpoint_requests",
			Help: "A counter that tracks requests to endpoints",
		},
		[]string{"endpoint", "code"},
	)
	prometheus.MustRegister(counter)

	return counter
}

// calls executes endpoint requests
func calls(data *Endpoints, counter prometheus.CounterVec) {
	for {
		for key, value := range data.Endpoints {
			resp, err := http.Get(value.Url)
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("request: %s, code: %v\n", key, resp.StatusCode)
			counter.With(prometheus.Labels{"endpoint": key, "code": "200"}).Inc()
		}
		time.Sleep(5 * time.Second)
	}
}

// Register handlers, run server
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
	counter := createMetrics()
	go calls(data, *counter)
	wg.Wait()
}
