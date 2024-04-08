package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type URL struct {
	Link string `json:"url"`
}

func main() {
	// Open and read the JSON file
	file, err := os.Open("urls.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var urls []URL
	// Decode the JSON data into a slice of URL structs
	err = json.NewDecoder(file).Decode(&urls)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	for {
		// Loop through each URL
		for _, url := range urls {
			// Make HTTP GET request to the URL
			resp, err := http.Get(url.Link)
			if err != nil {
				fmt.Printf("Error fetching %s: %s\n", url.Link, err)
				continue
			}
			defer resp.Body.Close()

			// Print the URL and its response code
			fmt.Printf("URL: %s, Response Code: %d\n", url.Link, resp.StatusCode)

			// Pause for 1 second before making the next request
			time.Sleep(1 * time.Second)
		}

		// Pause for 30 seconds before starting the next iteration
		fmt.Println("...")
		time.Sleep(10 * time.Second)
	}
}
