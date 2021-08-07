package generator

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

//Test func
func Test() {
	i := 0
	go func() {
		time.Sleep(1 * time.Second)
		for i < 1 {
			t := time.Now().Format("2006-01-02 15:04:05.000000000")
			fmt.Println(t)
			_, err := http.Post("http://localhost:4000/log", "text", bytes.NewBufferString(t))
			if err != nil {
				fmt.Printf("There was an err: %v\n", err)
			}
			time.Sleep(5 * time.Second)
		}
	}()
	return
}
