package main

import (
	"context"
	"fmt"
	"log"

	"github.com/josepht96/learning/go/serviceMonitor/logger"
	"github.com/josepht96/learning/go/serviceMonitor/registry"
	"github.com/josepht96/learning/go/serviceMonitor/service"
)

func main() {
	logger.Run("./cmd/logs/serviceMonitor.log")

	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	//Create registry object
	var reg registry.Registration
	reg.ServiceName = registry.LogService
	reg.ServiceURL = serviceAddress
	reg.Host = host
	reg.Port = port

	ctx, err := service.Start(context.Background(), reg, logger.RegisterHandlers)
	if err != nil {
		log.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("Shuttting down")
}
