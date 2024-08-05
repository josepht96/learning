package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
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

	client := CreateClient()
	DeployModel(client)

	// log.Println(response.String())

	// client.NewJobWorker().JobType("payment-service").Handler(handleJobPayment).Open()
	// client.NewJobWorker().JobType("shipment-service").Handler(handleJobShipment).Open()

	// // create a new workflow instance
	// variables := make(map[string]interface{})
	// i := 0
	// for {
	// 	fmt.Println("Creating new instance...")
	// 	variables["orderId"] = strconv.Itoa(i)
	// 	request, err := client.NewCreateInstanceCommand().BPMNProcessId("test-process-manual").LatestVersion().VariablesFromMap(variables)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	result, err := request.Send(ctx)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(result.String())
	// 	i++
	// 	fmt.Println("Sleeping...")
	// 	time.Sleep(30 * time.Second)
	// }
}

func DeployModel(client zbc.Client) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFn()

	response, err := client.NewDeployResourceCommand().AddResourceFile("test-workflow-manual.bpmn").Send(ctx)
	if err != nil {
		panic(err)
	}
	log.Println(response)
	// if len(response.GetDeployments()) < 0 {
	// 	panic(errors.New("failed to deploy send-email model; nothing was deployed"))
	// }
}

func CreateClient() zbc.Client {
	config := zbc.ClientConfig{
		UsePlaintextConnection: true,
	}

	client, err := zbc.NewClient(&config)
	if err != nil {
		panic(err)
	}

	return client
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
