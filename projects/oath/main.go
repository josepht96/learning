package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/josepht96/learning/oauth/controllers"
)

func main() {
	controllers.RegisterControllers()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		log.Printf("server is listening at %s", "http://localhost:8080")
		http.ListenAndServe(":8080", nil)
	}()
	wg.Wait()
}
