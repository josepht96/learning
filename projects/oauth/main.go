package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/josepht96/learning/oauth/controllers"
)

var Port = 8080

func main() {
	controllers.RegisterControllers()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		log.Printf("server is listening at http://localhost:%d", Port)
		http.ListenAndServe(fmt.Sprintf(":%d", Port), nil)
	}()
	wg.Wait()
}
