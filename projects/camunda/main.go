package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/camunda/zeebe/clients/go/v8/pkg/entities"
	"github.com/camunda/zeebe/clients/go/v8/pkg/worker"
	"github.com/camunda/zeebe/clients/go/v8/pkg/zbc"
)

const ZeebeAddr = "zeebe.camunda.svc.cluster.local:26500"

/*
Sample application that connects to a cluster on Camunda Cloud, or a locally deployed cluster.

When connecting to a cluster in Camunda Cloud, this application assumes that the following
environment variables are set:

ZEEBE_ADDRESS
ZEEBE_CLIENT_ID
ZEEBE_CLIENT_SECRET
ZEEBE_AUTHORIZATION_SERVER_URL

Hint: When you create client credentials in Camunda Cloud you have the option
to download a file with the lines above filled out for you.

When connecting to a local cluster or node, this application assumes default port and no
authentication or encryption enabled.
*/

func main() {
	gatewayAddr := os.Getenv("ZEEBE_ADDRESS")
	plainText := true

	if gatewayAddr == "" {
		gatewayAddr = ZeebeAddr
		plainText = true
	}

	if os.Getenv("ZEEBE_CLIENT_ID") == "" {
		log.Fatal("ZEEBE_CLIENT_ID not set")
	}
	if os.Getenv("ZEEBE_CLIENT_SECRET") == "" {
		log.Fatal("ZEEBE_CLIENT_SECRET not set")
	}
	if os.Getenv("ZEEBE_AUTHORIZATION_SERVER_URL") == "" {
		log.Fatal("ZEEBE_AUTHORIZATION_SERVER_URL not set")
	}

	credsProvider, err := zbc.NewOAuthCredentialsProvider(&zbc.OAuthProviderConfig{
		ClientID:               os.Getenv("ZEEBE_CLIENT_ID"),
		ClientSecret:           os.Getenv("ZEEBE_CLIENT_SECRET"),
		Audience:               gatewayAddr,
		AuthorizationServerURL: os.Getenv("ZEEBE_AUTHORIZATION_SERVER_URL"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(gatewayAddr)
	zbClient, err := zbc.NewClient(&zbc.ClientConfig{
		GatewayAddress:         gatewayAddr,
		UsePlaintextConnection: plainText,
		CredentialsProvider:    credsProvider,
	})

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

	zbClient.NewJobWorker().JobType("payment-service").Handler(handleJob).Open()

	// create a new workflow instance
	variables := make(map[string]interface{})
	i := 0
	for {
		fmt.Println("Creating new instance...")
		i++
		variables["orderId"] = strconv.Itoa(i)
		request, err := zbClient.NewCreateInstanceCommand().BPMNProcessId("order-process-4").LatestVersion().VariablesFromMap(variables)
		if err != nil {
			panic(err)
		}
		result, err := request.Send(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Println(result.String())
		fmt.Println("Sleeping...")
		time.Sleep(5 * time.Second)
	}
}

func handleJob(client worker.JobClient, job entities.Job) {
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
	log.Println("Collect money using payment method:", headers["method"])

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
