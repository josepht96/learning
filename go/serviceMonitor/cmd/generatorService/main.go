package main

import (
	"context"
	"fmt"
	"log"

	"github.com/josepht96/learning/go/serviceMonitor/generator"
	"github.com/josepht96/learning/go/serviceMonitor/registry"
	"github.com/josepht96/learning/go/serviceMonitor/service"
)

func main() {
	host, port := "localhost", "4040"
	serviceAddress := fmt.Sprintf("https://%v:%v", host, port)

	//Create registry object
	var reg registry.Registration
	reg.ServiceName = registry.GeneratorService
	reg.ServiceURL = serviceAddress
	reg.Host = host
	reg.Port = port

	//Start service server
	ctx, err := service.Start(context.Background(), reg, generator.Test)
	if err != nil {
		log.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("Shuttting down")
}
