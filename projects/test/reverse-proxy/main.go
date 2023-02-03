package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {

	// define origin server URL
	originServerURL, err := url.Parse("http://127.0.0.1:8085")
	if err != nil {
		log.Fatal("invalid origin server URL")
	}

	reverseProxy := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("[reverse proxy server] received request at: %s\n", time.Now())

		// set req Host, URL and Request URI to forward a request to the origin server
		fmt.Println(req.Host)
		req.Host = originServerURL.Host
		req.URL.Host = originServerURL.Host
		req.URL.Scheme = originServerURL.Scheme
		req.RequestURI = ""

		// save the response from the origin server
		originServerResponse, err := http.DefaultClient.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(w, err)
			return
		}

		// return response to the client
		w.WriteHeader(http.StatusOK)
		io.Copy(w, originServerResponse.Body)
	})

	log.Fatal(http.ListenAndServe(":8080", reverseProxy))
}
