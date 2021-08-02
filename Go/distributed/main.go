package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Song{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	cacheCh := make(chan Song)
	dbCh := make(chan Song)

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10)
		//time.Sleep(2 * time.Second)
		fmt.Println(id)
		wg.Add(3)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, cacheCh chan<- Song) {
			if s, ok := queryCache(id, m); ok {
				cacheCh <- s
			}
			fmt.Println("in cache")
			wg.Done()
		}(id, wg, m, cacheCh)

		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, dbCh chan<- Song) {
			if s, ok := queryDatabase(id, m); ok {
				dbCh <- s
			}
			fmt.Println("in database")
			wg.Done()
		}(id, wg, m, dbCh)

		go func(cacheCh, dbCh <-chan Song) {

			select {
			case s := <-cacheCh:
				fmt.Println("from cache")
				fmt.Println(s)
				<-dbCh
			case s := <-dbCh:
				fmt.Println("from database")
				fmt.Println(s)
			}
			fmt.Println("in message queue")
			wg.Done()
		}(cacheCh, dbCh)
		wg.Wait()

	}
	//channel()
}

func queryCache(id int, m *sync.RWMutex) (Song, bool) {
	m.RLock()
	s, ok := cache[id]
	m.RUnlock()
	return s, ok
}

func queryDatabase(id int, m *sync.RWMutex) (Song, bool) {
	for _, s := range songs {
		if s.ID == id {
			m.Lock()
			cache[id] = s
			m.Unlock()
			return s, true
		}
	}
	return Song{}, false
}
