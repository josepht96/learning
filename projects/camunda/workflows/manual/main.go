package main

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/camunda/zeebe/clients/go/v8/pkg/entities"
	"github.com/camunda/zeebe/clients/go/v8/pkg/worker"
	"github.com/camunda/zeebe/clients/go/v8/pkg/zbc"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Listening on http://localhost:8089")
	go http.ListenAndServe(":8089", nil)

	ctx, cancelFn := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFn()

	client := createClient()
	deployModel(client, ctx)
	var wg sync.WaitGroup
	wg.Add(4)
	go createNewProcessInstance(client)
	go activateAndProcessJob(client, "test-workflow-service-task-1")
	go activateAndProcessJob(client, "test-workflow-service-task-2")
	go activateAndProcessJob(client, "test-workflow-service-task-3")

	wg.Wait()
}

func createClient() zbc.Client {
	credentials, err := zbc.NewOAuthCredentialsProvider(&zbc.OAuthProviderConfig{
		Audience: "zeebe-api",
	})
	if err != nil {
		panic(err)
	}

	config := zbc.ClientConfig{
		CredentialsProvider:    credentials,
		UsePlaintextConnection: true,
	}

	client, err := zbc.NewClient(&config)
	if err != nil {
		panic(err)
	}

	return client
}

func deployModel(client zbc.Client, ctx context.Context) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFn()

	response, err := client.NewDeployResourceCommand().AddResourceFile("test-workflow-manual.bpmn").TenantId("tenant1").Send(ctx)
	if err != nil {
		panic(err)
	}
	log.Printf("Deploying model: %s", response)
	// if len(response.GetDeployments()) < 0 {
	// 	panic(errors.New("failed to deploy send-email model; nothing was deployed"))
	// }
}

func createNewProcessInstance(client zbc.Client) {
	variables := make(map[string]interface{})
	for {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		log.Println("Creating new process instance...")
		variables["testVar1"] = "123"

		request, err := client.NewCreateInstanceCommand().
			BPMNProcessId("test-workflow").
			LatestVersion().
			TenantId("tenant1").
			VariablesFromMap(variables)

		if err != nil {
			panic(err)
		}
		response, err := request.Send(ctx)
		if err != nil {
			panic(err)
		}
		log.Printf("Response: %s", response.String())
		time.Sleep(5 * time.Second)
	}
}

func activateAndProcessJob(client zbc.Client, jobType string) {
	for {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

		// Activate a job
		response, err := client.NewActivateJobsCommand().
			JobType(jobType).
			MaxJobsToActivate(1).
			TenantIds("tenant1").
			WorkerName("test-workflow-service-task-1").
			Send(ctx)

		if err != nil {
			log.Printf("failed to activate jobs: %v", err)
			continue
		}

		if len(response) == 0 {
			log.Println("No jobs available")
			continue
		}

		job := response[0]
		log.Printf("Activated job: %v", job.ProcessInstanceKey)

		// Report job completion
		_, err = client.NewCompleteJobCommand().JobKey(job.Key).Send(ctx)
		if err != nil {
			log.Printf("failed to report job completion: %v", err)
		}

		log.Printf("Job %d completed successfully", job.Key)
		time.Sleep(5 * time.Second)
	}
}

func failJob(client worker.JobClient, job entities.Job) {
	log.Println("Failed to complete job", job.GetKey())

	ctx := context.Background()
	_, err := client.NewFailJobCommand().JobKey(job.GetKey()).Retries(job.Retries - 1).Send(ctx)
	if err != nil {
		panic(err)
	}
}
