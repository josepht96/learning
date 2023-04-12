package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/camunda/zeebe/clients/go/v8/pkg/entities"
	"github.com/camunda/zeebe/clients/go/v8/pkg/worker"
	"github.com/camunda/zeebe/clients/go/v8/pkg/zbc"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Listening on http://localhost:8080")
	go http.ListenAndServe(":8080", nil)

	zbClient, err := zbc.NewClient(&zbc.ClientConfig{})

	if err != nil {
		panic(err)
	}

	// deploy workflow
	ctx := context.Background()
	response, err := zbClient.NewDeployResourceCommand().AddResourceFile("order-process-4.bpmn").Send(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.String())

	zbClient.NewJobWorker().JobType("payment-service").Handler(handleJobPayment).Open()
	zbClient.NewJobWorker().JobType("shipment-service").Handler(handleJobShipment).Open()

	// create a new workflow instance
	variables := make(map[string]interface{})
	i := 0
	for {
		fmt.Println("Creating new instance...")
		variables["orderId"] = strconv.Itoa(i)
		request, err := zbClient.NewCreateInstanceCommand().BPMNProcessId("order-process-4-manual").LatestVersion().VariablesFromMap(variables)
		if err != nil {
			panic(err)
		}
		result, err := request.Send(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Println(result.String())
		i++
		fmt.Println("Sleeping...")
		time.Sleep(30 * time.Second)
	}
}

func handleJobPayment(client worker.JobClient, job entities.Job) {
	jobKey := job.GetKey()

	headers, err := job.GetCustomHeadersAsMap()
	if err != nil {
		// failed to handle job as we require the custom job headers
		failJob(client, job)
		return
	}

	variables, err := job.GetVariablesAsMap()
	if err != nil {
		// failed to handle job as we require the variables
		failJob(client, job)
		return
	}

	variables["totalPrice"] = 46.50
	request, err := client.NewCompleteJobCommand().JobKey(jobKey).VariablesFromMap(variables)
	if err != nil {
		// failed to set the updated variables
		failJob(client, job)
		return
	}

	log.Println("Complete job", jobKey, "of type", job.Type)
	log.Println("Payment processed:", headers["method"])

	ctx := context.Background()
	_, err = request.Send(ctx)
	if err != nil {
		panic(err)
	}

	log.Println("Successfully completed job")
}

func handleJobShipment(client worker.JobClient, job entities.Job) {
	time.Sleep(15 * time.Second)
	jobKey := job.GetKey()

	headers, err := job.GetCustomHeadersAsMap()
	if err != nil {
		// failed to handle job as we require the custom job headers
		failJob(client, job)
		return
	}

	variables, err := job.GetVariablesAsMap()
	if err != nil {
		// failed to handle job as we require the variables
		failJob(client, job)
		return
	}

	variables["totalPrice"] = 46.50
	request, err := client.NewCompleteJobCommand().JobKey(jobKey).VariablesFromMap(variables)
	if err != nil {
		// failed to set the updated variables
		failJob(client, job)
		return
	}

	log.Println("Complete job", jobKey, "of type", job.Type)
	log.Println("Shipment processed:", headers["method"])

	ctx := context.Background()
	_, err = request.Send(ctx)
	if err != nil {
		panic(err)
	}

	log.Println("Successfully completed job")
}
func failJob(client worker.JobClient, job entities.Job) {
	log.Println("Failed to complete job", job.GetKey())

	ctx := context.Background()
	_, err := client.NewFailJobCommand().JobKey(job.GetKey()).Retries(job.Retries - 1).Send(ctx)
	if err != nil {
		panic(err)
	}
}
