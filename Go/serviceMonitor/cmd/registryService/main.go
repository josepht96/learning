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
	server.Addr = fmt.Sprintf("localhost%v", registry.ServerPort)

	go func() {
		log.Println(server.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Printf("Registry service started. Press key to stop...\n")
		var s string
		fmt.Scanln(&s)
		server.Shutdown(ctx)
		cancel()
	}()
	<-ctx.Done()
	fmt.Println("Shutting down the registry service...")

}
