package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Response struct {
	Status     string  `json:"status"`
	StatusCode int     `json:"statusCode"`
	Body       Message `json:"body"`
}

type Message struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

var port = 8080

// Define the upgrader with options
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Error while upgrading connection:", err)
			return
		}
		defer conn.Close()

		fmt.Println("Client connected")

		for name, values := range r.Header {
			// Loop over all values for the name.
			for _, value := range values {
				fmt.Println(name, value)
			}
		}
		log.Println(r)
		log.Printf("Host: %s\n", r.Host)
		log.Printf("Method: %s\n", r.Method)
		log.Printf("Proto: %s\n", r.Proto)
		log.Printf("RemoteAddr: %s\n", r.RemoteAddr)
		log.Printf("RequestURI: %s\n", r.RequestURI)
		log.Printf("URL: %s\n", r.URL)
		log.Printf("Body: %s\n", r.Body)
		log.Printf("TLS: %v\n", r.TLS)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		dt := time.Now()
		data := Response{
			Status:     "OK",
			StatusCode: 200,
			Body: Message{
				Message: "hello from game-server",
				Time:    dt.String(),
			},
		}
		fmt.Println(data)
		messageType, msg, err := conn.ReadMessage()
		log.Printf("messageType: %s\n", messageType)
		log.Printf("msg: %s\n", msg)
		// json.NewEncoder(w).Encode(data)
		for {
			err = conn.WriteJSON(data)
			if err != nil {
				log.Println("Error while writing message:", err)
			}
			time.Sleep(5 * time.Second)
		}

		fmt.Println("---------------------------------------------------------")
	})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		log.Printf("server is listening at http://localhost:%d", port)
		http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	}()
	wg.Wait()
}
