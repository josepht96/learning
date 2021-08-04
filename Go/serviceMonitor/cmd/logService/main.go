package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/josepht96/learning/go/serviceMonitor/log"
	"github.com/josepht96/learning/go/serviceMonitor/registry"
	"github.com/josepht96/learning/go/serviceMonitor/service"
)

func main() {
	log.Run("./serviceMonitor.log")

	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var reg registry.Registration
	reg.ServiceName = registry.LogService
	reg.ServiceURL = serviceAddress

	ctx, err := service.Start(context.Background(),
		reg, host, port, log.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("Shuttting down")
}
