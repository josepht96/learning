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
	m := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10)
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.Mutex) {
			s, ok := queryCache(id, m)
			if ok {
				fmt.Println("From cache")
				fmt.Println(s)
			} else if !ok {
				fmt.Printf("id (%v) not found in cache\n", id)
			}
			wg.Done()
		}(id, wg, m)

		go func(id int, wg *sync.WaitGroup, m *sync.Mutex) {
			s, ok := queryDatabase(id, m)
			if ok {
				fmt.Println("from database")
				fmt.Println(s)
			} else if !ok {
				fmt.Printf("id (%v) not found in database\n", id)
			}

			wg.Done()
		}(id, wg, m)
	}

	wg.Wait()
}

func queryCache(id int, m *sync.Mutex) (Song, bool) {
	s, ok := cache[id]
	return s, ok
}

func queryDatabase(id int, m *sync.Mutex) (Song, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, s := range songs {
		if s.ID == id {
			cache[id] = s
			return s, true
		}
	}
	return Song{}, false
}
