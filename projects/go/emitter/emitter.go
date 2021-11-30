package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type foo struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

func apiCall(randomInt int) {
	s := foo{}
	url := fmt.Sprintf("http://localhost:5000/foo/%v", randomInt)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bodyBytes, &s)
	fmt.Printf("Name: %v: ID: %v\n", s.Name, s.ID)
}
func main() {
	fmt.Println("Emitter starting up...")
	stop := false
	for stop == false {
		randomInt := rand.Intn(100 - 1)
		apiCall(randomInt)
		time.Sleep(5 * time.Second)
	}
}
