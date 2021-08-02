package main

import (
	"fmt"
	"sync"
)

func channel() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)
	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		for msg := range ch {
			fmt.Println(msg)
		}
		wg.Done()
	}(ch, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		sum := 0
		for sum < 10 {
			ch <- rnd.Intn(10)
			sum++
		}
		close(ch)
		wg.Done()
	}(ch, wg)
	wg.Wait()

}
