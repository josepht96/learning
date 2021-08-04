package service

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/josepht96/learning/go/serviceMonitor/registry"
)

//Start func
func Start(ctx context.Context, reg registry.Registration, host, port string,
	registerHandlers func()) (context.Context, error) {
	fmt.Println("in Start")
	registerHandlers()
	ctx = startService(ctx, reg.ServiceName, host, port)
	err := registry.RegisterService(reg)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}

func startService(ctx context.Context, serviceName registry.ServiceName,
	host, port string) context.Context {
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
		err := registry.ShutDownService(fmt.Sprintf("http://%v:%v", host, port))
		if err != nil {

		}
		server.Shutdown(ctx)
		cancel()
	}()
	return ctx
}
