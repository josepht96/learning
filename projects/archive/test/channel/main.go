package main

import (
	"fmt"
	"sync"
)

//returnValue wont return until its value has been read through the channel
func returnValue(v int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	c <- v
}

func main() {
	wg := new(sync.WaitGroup)
	c := make(chan int)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go returnValue(i, c, wg)
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	total := 0
	for value := range c {
		fmt.Println(value)
		total = total + value
	}

	fmt.Printf("sum %v\n", total)
}
