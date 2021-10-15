package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	i := 0
	for i < 1 {
		time.Sleep(2 * time.Second)
		resp, err := http.Get("http://sender-service:80/")
		if err != nil {
			fmt.Printf("There was an error: %v\n", err)
			continue
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("There was an error: %v\n", err)
		}
		fmt.Printf("Response received at: %s\n", string(body))
	}
}
