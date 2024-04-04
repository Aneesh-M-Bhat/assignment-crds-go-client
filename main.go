package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func getClient() *dynamic.DynamicClient {
	var kubeconfig = "./config.yaml"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %v", err)
	}

	// Create a dynamic client
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating dynamic client: %v", err)
	} else {
		log.Println("Dynamic client created")
	}
	return client
}

func taskWithRuns(client *dynamic.DynamicClient) {
	createResource(client, getExampleTask(), "v1beta1", "tasks")
	createResource(client, getExampleTaskRun(), "v1beta1", "taskruns")
	listResources(client, "v1beta1", "taskruns")
	listResources(client, "v1beta1", "tasks")
	deleteResource(client, "hello-task-run", "v1beta1", "taskruns")
	deleteResource(client, "hello", "v1beta1", "tasks")
}

func pipelineWithRuns(client *dynamic.DynamicClient) {
	createResource(client, getExampleTask(), "v1beta1", "tasks")
	createResource(client, getExampleTask2(), "v1beta1", "tasks")
	createResource(client, getExamplePipeline(), "v1beta1", "pipelines")
	createResource(client, getExamplePipelineRun(), "v1beta1", "pipelineruns")
	listResources(client, "v1beta1", "pipelines")
	listResources(client, "v1beta1", "pipelineruns")
	deleteResource(client, "hello-goodbye-run", "v1beta1", "pipelineruns")
	deleteResource(client, "hello-goodbye", "v1beta1", "pipelines")
	deleteResource(client, "hello", "v1beta1", "tasks")
	deleteResource(client, "goodbye", "v1beta1", "tasks")
}

func customRun(client *dynamic.DynamicClient) {
	createResource(client, getExampleCustomRun(), "v1beta1", "customruns")
	listResources(client, "v1beta1", "customruns")
	deleteResource(client, "my-example-run", "v1beta1", "customruns")
}

func clusterTask(client *dynamic.DynamicClient) {
	createClusterTask(client, getExampleClusterTask(), "v1beta1", "clustertasks")
	listClusterTasks(client, "v1beta1", "clustertasks")
	deleteClusterTask(client, "example-cluster-task", "v1beta1", "clustertasks")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	client := getClient()

	for {
		fmt.Println("Select demo number (1-4) or 0 to exit:")
		fmt.Println("1. Task & TaskRun Demo")
		fmt.Println("2. Pipeline & PipelineRuns Demo")
		fmt.Println("3. CustomRuns")
		fmt.Println("4. ClusterTasks")

		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1]

		option, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch option {
		case 0:
			fmt.Println("Exiting...")
			return
		case 1:
			taskWithRuns(client)
		case 2:
			pipelineWithRuns(client)
		case 3:
			customRun(client)
		case 4:
			clusterTask(client)
		default:
			fmt.Println("Invalid Option.")
		}
	}
}
