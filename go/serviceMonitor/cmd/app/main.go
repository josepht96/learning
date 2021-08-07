package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i < 1 {
		t := time.Now().Format("2006-01-02 15:04:05.000000000")
		fmt.Println(t)
		time.Sleep(5 * time.Second)
	}

}
