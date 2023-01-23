package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func getWithGoRoutines(result chan bool) {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 1; i <= 60; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := http.Get("http://localhost:8080")
			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()
		}()
	}
	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("time with go routines: %v\n", duration)
	result <- true
}

func getWithoutGoRoutines(result chan bool) {
	start := time.Now()
	for i := 1; i <= 60; i++ {
		resp, err := http.Get("http://localhost:8080")
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
	}
	duration := time.Since(start)
	fmt.Printf("time without go routines: %v\n", duration)
	time.Sleep(5 * time.Second)
	result <- true
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	result := make(chan bool, 2)
	go getWithGoRoutines(result)
	go getWithoutGoRoutines(result)
	<-result
	wg.Done()
	fmt.Println("first result returned...")
	<-result
	wg.Done()
	fmt.Println("second result returned...")

	wg.Wait()
}
