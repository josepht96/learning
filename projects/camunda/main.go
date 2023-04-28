package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/camunda/zeebe/clients/go/v8/pkg/entities"
	"github.com/camunda/zeebe/clients/go/v8/pkg/worker"
	"github.com/camunda/zeebe/clients/go/v8/pkg/zbc"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var TASK_INTERVAL int = 1000

func main() {
	os.Setenv("ZEEBE_ADDRESS", "192.168.86.23:26500")
	os.Setenv("ZEEBE_HOST", "192.168.86.23")
	os.Setenv("ZEEBE_PORT", "26500")

	http.Handle("/metrics", promhttp.Handler())
	log.Println("Listening on http://localhost:8080")
	go http.ListenAndServe(":8080", nil)

	process_interval, err := strconv.Atoi(os.Getenv("PROCESS_INTERVAL"))
	if err != nil {
		log.Printf("error with enviroment variable 'PROCESS_INTERVAL': %s", err)
		process_interval = 1000
	}
	task_interval, err := strconv.Atoi(os.Getenv("TASK_INTERVAL"))
	if err != nil {
		log.Printf("error with environment variable 'TASK_INTERVAL': %s", err)
	} else {
		TASK_INTERVAL = task_interval
	}

	zbClient, err := zbc.NewClient(&zbc.ClientConfig{
		UsePlaintextConnection: true,
	})

	if err != nil {
		log.Panic(err)
	}

	// deploy workflow
	ctx := context.Background()
	response, err := zbClient.NewDeployResourceCommand().AddResourceFile("stress-test-process.bpmn").Send(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.String())

	zbClient.NewJobWorker().JobType("stress-test-1").Handler(handleStressTestJobFast).Open()
	zbClient.NewJobWorker().JobType("stress-test-2").Handler(handleStressTestJobSlow).Open()
	zbClient.NewJobWorker().JobType("stress-test-3").Handler(handleStressTestJobFast).Open()
	zbClient.NewJobWorker().JobType("stress-test-4").Handler(handleStressTestJobSlow).Open()
	zbClient.NewJobWorker().JobType("stress-test-5").Handler(handleStressTestJobFast).Open()
	zbClient.NewJobWorker().JobType("stress-test-6").Handler(handleStressTestJobSlow).Open()
	zbClient.NewJobWorker().JobType("stress-test-7").Handler(handleStressTestJobFast).Open()
	zbClient.NewJobWorker().JobType("stress-test-8").Handler(handleStressTestJobSlow).Open()
	zbClient.NewJobWorker().JobType("stress-test-9").Handler(handleStressTestJobFast).Open()
	zbClient.NewJobWorker().JobType("stress-test-10").Handler(handleStressTestJobSlow).Open()

	// create a new workflow instance
	variables := make(map[string]interface{})
	i := 0
	for {
		log.Println("Creating new stress test instance...")
		variables["testId"] = strconv.Itoa(i)
		request, err := zbClient.NewCreateInstanceCommand().BPMNProcessId("stress-test-process").LatestVersion().VariablesFromMap(variables)
		if err != nil {
			log.Println(err)
		}
		result, err := request.Send(ctx)
		if err != nil {
			log.Println(err)
		}
		log.Println(result.String())
		i++
		time.Sleep(time.Duration(process_interval) * time.Millisecond)
	}
}

func failJob(client worker.JobClient, job entities.Job) {
	log.Println("Failed to complete job", job.GetKey())

	ctx := context.Background()
	_, err := client.NewFailJobCommand().JobKey(job.GetKey()).Retries(job.Retries - 1).Send(ctx)
	if err != nil {
		log.Panic(err)
	}
}
