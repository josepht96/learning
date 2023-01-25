package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/golang/protobuf/proto"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		data := Person{
			name: "Joe",
			id:   123,
		}
		data, err := proto.Marshal(&data)
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}
	})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		log.Printf("server is listening at %s", "http://localhost:8080")
		http.ListenAndServe(":8080", nil)
	}()
	wg.Wait()
}
