package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/josepht96/learning/go/serviceMonitor/registry"
)

func main() {
	http.Handle("/services", &registry.RegistryService{})
	ctx, cancel := context.WithCancel(context.Background())
	var server http.Server
	server.Addr = fmt.Sprintf("%v", registry.ServerPort)

	go func() {
		fmt.Println("Starting registry service...")
		log.Println(server.ListenAndServe())
		cancel()
	}()

	// go func() {
	// 	fmt.Printf("Registry service started. Press key to stop...\n")
	// 	runnning := true
	// 	for runnning == true {
	// 		var s string
	// 		fmt.Scanln(&s)
	// 		if s != "" {
	// 			server.Shutdown(ctx)
	// 			cancel()
	// 			runnning = false
	// 		}
	// 	}

	// }()
	<-ctx.Done()
	fmt.Println("Shutting down the registry service...")

}
