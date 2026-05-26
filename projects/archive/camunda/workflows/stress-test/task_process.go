package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/camunda/zeebe/clients/go/v8/pkg/entities"
	"github.com/camunda/zeebe/clients/go/v8/pkg/worker"
)

func handleStressTestJobSlow(client worker.JobClient, job entities.Job) {
	time.Sleep(time.Duration(rand.Intn(TASK_INTERVAL)) * time.Millisecond)
	jobKey := job.GetKey()

	_, err := job.GetCustomHeadersAsMap()
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

	variables["testId"] = "stress_test"
	request, err := client.NewCompleteJobCommand().JobKey(jobKey).VariablesFromMap(variables)
	if err != nil {
		// failed to set the updated variables
		failJob(client, job)
		return
	}

	ctx := context.Background()
	_, err = request.Send(ctx)
	if err != nil {
		log.Println(err)
	}

	log.Println("Complete job", jobKey, "of type", job.Type)
}

func handleStressTestJobFast(client worker.JobClient, job entities.Job) {
	time.Sleep(100 * time.Millisecond)
	jobKey := job.GetKey()

	_, err := job.GetCustomHeadersAsMap()
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

	variables["testId"] = "stress_test"
	request, err := client.NewCompleteJobCommand().JobKey(jobKey).VariablesFromMap(variables)
	if err != nil {
		// failed to set the updated variables
		failJob(client, job)
		return
	}

	ctx := context.Background()
	_, err = request.Send(ctx)
	if err != nil {
		log.Println(err)
	}

	log.Println("Complete job", jobKey, "of type", job.Type)
}
