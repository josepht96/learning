package main

import (
	"fmt"
    "log"
    "net/http"
    "time"
 
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

func socketHandler(w http.ResponseWriter, r *http.Request) {
    // Upgrade our raw HTTP connection to a websocket based one
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Print("Error during connection upgrade:", err)
        return
    }
    defer conn.Close()
 
    // The event loop
    for {
        time.Sleep(2 * time.Second)
        message := getTime()
		messageType := websocket.TextMessage
        if err != nil {
            log.Println("Error during message reading:", err)
            break
        }
        log.Printf("%s", message)
        err = conn.WriteMessage(messageType, message)
        if err != nil {
            log.Println("Error during message writing:", err)
            break
        }
    }
}
 
func getTime() ([]byte) {
	t := time.Now().Format("2006-01-02 15:04:05.000000000")
	return []byte(fmt.Sprintf("time: %s", t))
	//w.Write([]byte(tString))
}
 
func main() {
    http.HandleFunc("/socket", socketHandler)
	fmt.Println("server starting at ws://localhost:8080/socket")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
