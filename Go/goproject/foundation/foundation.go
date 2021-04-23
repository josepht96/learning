package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type fooHandler struct {
	Message string `json:"message1"`
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/foo" {
		switch r.Method {
		case http.MethodGet:
			data, _ := json.Marshal(f)
			w.Write([]byte(data))
		case http.MethodPost:
			s := fooHandler{}
			bodyBytes, err := ioutil.ReadAll(r.Body)
			err = json.Unmarshal(bodyBytes, &s)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			data, _ := json.Marshal(s)
			w.Write([]byte(data))
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		switch r.Method {
		case http.MethodGet:
			f.Message = "not /foo route"
			data, _ := json.Marshal(f)
			w.Write([]byte(data))
		case http.MethodPost:
			s := fooHandler{}
			bodyBytes, err := ioutil.ReadAll(r.Body)
			err = json.Unmarshal(bodyBytes, &s)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			s.Message = "also not /foo route"
			data, _ := json.Marshal(s)
			w.Write([]byte(data))
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}

}

func middleWareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Before handler; middleware startup")
		start := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Printf("Middleware finished; %v\n", time.Since(start))
	})
}
func main() {
	http.Handle("/foo", middleWareHandler(&fooHandler{Message: "hello world"}))
	http.Handle("/foo/", middleWareHandler(&fooHandler{Message: "hello worlds"}))
	fmt.Println("Listening at: localhost:5000")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
