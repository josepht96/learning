// client.go
package main
 
import (
	"fmt"
    "log"
    "os"
    "os/signal"
    "time"
 
    "github.com/gorilla/websocket"
)
 
var done chan interface{}
var interrupt chan os.Signal
type connObj struct {
	connected bool
	conn *websocket.Conn
}
func receiveHandler(conn *websocket.Conn) {
    defer close(done)
	for {
		//time.Sleep(2 * time.Second)
        _, msg, err := conn.ReadMessage()
        if err != nil {
            log.Println("Error in receive:", err)
			return
        }
        log.Printf("Received: %s\n", msg)
    }
}
func establishConn(c *connObj) {
	socketUrl := "ws://localhost:8080" + "/socket"
	var err error
    c.conn, _, err = websocket.DefaultDialer.Dial(socketUrl, nil)
    if err != nil {
        fmt.Printf("Error connecting to Websocket Server:%s\n", err)
		c.connected = false
    }
    defer c.conn.Close()
	if (c.connected) {
		go receiveHandler(c.conn)
	}
}
func main() {
	c := connObj {
		connected: false,
		conn: nil,
	}
    done = make(chan interface{}) // Channel to indicate that the receiverHandler is done
    interrupt = make(chan os.Signal) // Channel to listen for interrupt signal to terminate gracefully
    signal.Notify(interrupt, os.Interrupt) // Notify the interrupt channel for SIGINT
	establishConn(&c)
    for {
		fmt.Println("waiting for case match...")
		if (!c.connected) {
			time.Sleep(2 * time.Second)
			fmt.Println("Attempting to connect to websocket...")
			establishConn(&c)
		}
        select {
        // case <-time.After(time.Duration(2) * time.Millisecond * 1000):
        //     // Send an echo packet every second
        //     err := conn.WriteMessage(websocket.TextMessage, []byte(""))
        //     if err != nil {
        //         log.Println("Error during writing to websocket:", err)
        //         return
        //     }

        case <-interrupt:
            // We received a SIGINT (Ctrl + C). Terminate gracefully...
            log.Println("Received SIGINT interrupt signal. Closing all pending connections")
 
            // Close our websocket connection
            err := c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
            if err != nil {
                log.Println("Error during closing websocket:", err)
                return
            }
 
            select {
            case <-done:
                log.Println("Receiver Channel Closed! Exiting....")
            case <-time.After(time.Duration(1) * time.Second):
                log.Println("Timeout in closing receiving channel. Exiting....")
            }
            return
		}
	}
}