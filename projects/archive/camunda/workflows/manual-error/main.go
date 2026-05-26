package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/camunda/zeebe/clients/go/v8/pkg/entities"
	"github.com/camunda/zeebe/clients/go/v8/pkg/worker"
	"github.com/camunda/zeebe/clients/go/v8/pkg/zbc"
)

func main() {
	zbClient, err := zbc.NewClient(&zbc.ClientConfig{
		UsePlaintextConnection: true,
	})

	if err != nil {
		panic(err)
	}

	// deploy workflow
	ctx := context.Background()
	response, err := zbClient.NewDeployResourceCommand().AddResourceFile("test-workflow.bpmn").Send(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.String())
	time.Sleep(2 * time.Second)

	zbClient.NewJobWorker().JobType("sys-task-1").Handler(handleJobSysTask1).Open()
	// zbClient.NewJobWorker().JobType("user-task-1").Handler(handleJobSysTask2).Open()
	zbClient.NewJobWorker().JobType("sys-task-2").Handler(handleJobSysTaskFail).Open()

	// create a new workflow instance
	variables := make(map[string]interface{})
	i := 0
	for {
		fmt.Println("Creating new instance of sample-workflow-manual")
		variables["orderId"] = strconv.Itoa(i)
		request, err := zbClient.NewCreateInstanceCommand().BPMNProcessId("sample-workflow-manual").LatestVersion().VariablesFromMap(variables)
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
		time.Sleep(15 * time.Second)
	}
}

func handleJobSysTask1(client worker.JobClient, job entities.Job) {
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

	log.Printf("Processing %s, %s: %v", headers["method"], job.Type, jobKey)

	variables["sampleValue"] = 1.00
	request, err := client.NewCompleteJobCommand().JobKey(jobKey).VariablesFromMap(variables)
	if err != nil {
		// failed to set the updated variables
		failJob(client, job)
		return
	}

	ctx := context.Background()
	_, err = request.Send(ctx)
	if err != nil {
		panic(err)
	}
	log.Printf("Completed %s, %s: %v", headers["method"], job.Type, jobKey)
}

func handleJobSysTaskFail(client worker.JobClient, job entities.Job) {
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

	log.Printf("Processing %s, %s: %v", headers["method"], job.Type, jobKey)

	variables["sampleValue"] = 1.00
	request, err := client.NewCompleteJobCommand().JobKey(jobKey).VariablesFromMap(variables)
	if err == nil {
		// failed to set the updated variables
		failJob(client, job)
		return
	}

	ctx := context.Background()
	_, err = request.Send(ctx)
	if err != nil {
		panic(err)
	}
	log.Printf("Completed %s, %s: %v", headers["method"], job.Type, jobKey)
}

func failJob(client worker.JobClient, job entities.Job) {
	log.Printf("Failed to complete job %s: %v", job.Type, job.GetKey())

	ctx := context.Background()
	_, err := client.NewFailJobCommand().JobKey(job.GetKey()).Retries(job.Retries - 1).Send(ctx)
	if err != nil {
		panic(err)
	}
}
