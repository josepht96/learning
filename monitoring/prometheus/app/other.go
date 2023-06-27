package main

import (
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main2() {
	counter := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "successful_documentum_tests",
		Help: "A counter that tracks successful documentum tests",
	})

	prometheus.MustRegister(counter)
	// Start the HTTP server to expose the metrics.
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			println("Failed to start HTTP server:", err)
			os.Exit(1)
		}
	}()
	for {
		// code to test documentum
		counter.Inc()
		time.Sleep(10 * time.Second)
	}
	// switch r.Method {
	// case http.MethodGet:
	// 	for name, values := range r.Header {
	// 		// Loop over all values for the name.
	// 		for _, value := range values {
	// 			fmt.Println(name, value)
	// 		}
	// 	}
	// 	(w).Header().Set("Access-Control-Allow-Origin", "*")
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusOK)

	// 	dt := time.Now()
	// 	data := Response{
	// 		Status:     "OK",
	// 		StatusCode: 200,
	// 		Body: Message{
	// 			Message: "hello from hello-server",
	// 			Time:    dt.String(),
	// 		},
	// 	}
	// 	fmt.Println(data)
	// 	json.NewEncoder(w).Encode(data)
	// 	fmt.Println("---------------------------------------------------------")
	// case http.MethodPost:
	// 	w.WriteHeader(http.StatusNotFound)
	// case http.MethodPut:
	// 	w.WriteHeader(http.StatusServiceUnavailable)
	// }
}
