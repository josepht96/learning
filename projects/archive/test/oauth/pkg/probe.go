package pkg

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var line = "--------------------------------------------------------------------"

type Header struct {
	Key   string
	Value string
}

func sendRequest(url string, headers []Header) (string, error) {
	start := time.Now()
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}
	fmt.Println(req)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", body)
	fmt.Println(line)

	return string(body), nil
}

func Probe(url string, headers []Header) (string, error) {
	return sendRequest(url, headers)
}
