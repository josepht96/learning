package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var line = "--------------------------------------------------------------------"

func main() {
	url := os.Getenv("URL")
	if url == "" {
		log.Fatal("error: environment var 'URL' not set")
	}
	for {
		go func() {
			start := time.Now()
			resp, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
			duration := time.Since(start)
			fmt.Printf("url: %s\n", url)
			fmt.Printf("latency: %v\n", duration)
			// print response code
			fmt.Printf("response code: %d\n", resp.StatusCode)
			// print headers
			fmt.Printf("headers:\n")
			for key, value := range resp.Header {
				fmt.Printf("\t%s: %s\n", key, value)
			}
			// print response body
			fmt.Printf("body:\n")
			// body, err := io.ReadAll(resp.Body)
			// if err != nil {
			// 	log.Fatalln(err)
			// }
			fmt.Printf("\t%v\n", resp.Body)
			fmt.Println(line)
		}()
		time.Sleep(5 * time.Second)
	}
}
