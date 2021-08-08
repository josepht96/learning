package middleware

import (
	"fmt"
	"net/http"
	"time"
)

//Handler handles CORS and any additional middleware level calls
func Handler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "text") //application/json
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")
		start := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Printf("thread finished; %v\n", time.Since(start))
	})
}
