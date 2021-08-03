package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

//Start func
func Start(ctx context.Context, serviceName, host, port string,
	registerHandlers func()) (context.Context, error) {
	fmt.Println("in Start")
	registerHandlers()
	ctx = startService(ctx, serviceName, host, port)

	return ctx, nil
}

func startService(ctx context.Context, serviceName, host, port string) context.Context {
	fmt.Println("in startService")
	ctx, cancel := context.WithCancel(ctx)

	var server http.Server
	server.Addr = ":" + port

	go func() {
		log.Println(server.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Printf("%v started. Press key to stop. \n", serviceName)
		var s string
		fmt.Scanln(&s)
		server.Shutdown(ctx)
		cancel()
	}()
	return ctx
}
