package service

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/josepht96/learning/go/serviceMonitor/registry"
)

//Start func
func Start(ctx context.Context, reg registry.Registration,
	registerHandlers func()) (context.Context, error) {
	fmt.Println("in Start")
	registerHandlers()
	ctx = startService(ctx, reg)
	err := registry.RegisterService(reg)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}

func startService(ctx context.Context, reg registry.Registration) context.Context {
	fmt.Println("in startService")
	ctx, cancel := context.WithCancel(ctx)

	var server http.Server
	server.Addr = fmt.Sprintf("%v:%v", reg.Host, reg.Port)
	go func() {
		log.Println(server.ListenAndServe())
		cancel()
	}()

	// go func() {
	// 	fmt.Printf("%v started. Press key to stop. \n", reg.ServiceURL)
	// 	var s string
	// 	fmt.Scanln(&s)
	// 	println(s)
	// 	err := registry.ShutDownService(fmt.Sprintf("%v", reg.ServiceURL))
	// 	if err != nil {

	// 	}
	// 	server.Shutdown(ctx)
	// 	cancel()
	// }()
	return ctx
}
