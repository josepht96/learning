package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func data(w http.ResponseWriter, r *http.Request) {
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()

	// The event loop
	i := true
	for i {
		err = conn.WriteMessage(1, []byte("test"))
		if err != nil {
			log.Println("Error during message writing:", err)
			break
		}
	}

}
func main() {
	RegisterControllers()
	fmt.Println("Listening at: http://localhost:4040/data")
	err := http.ListenAndServe(":4040", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func RegisterControllers() {
	fmt.Println("Generating handlers...")
	http.Handle("/", http.FileServer(http.Dir("./websocket")))
	http.HandleFunc("/data", data)

}
