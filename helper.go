package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

func prompt() {
	fmt.Printf("-> Press Return key to continue.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println()
}

func createResource(client *dynamic.DynamicClient, obj *unstructured.Unstructured, version string, resource string) {
	result, err := client.Resource(schema.GroupVersionResource{
		Group:    "tekton.dev",
		Version:  version,
		Resource: resource,
	}).Namespace("default").Create(context.TODO(), obj, metav1.CreateOptions{})

	if err != nil {
		fmt.Println("Error creating custom resource", err)
	} else {
		fmt.Println("Custom resource created successfully ", result)
	}
	prompt()

}

func listResources(client *dynamic.DynamicClient, version string, resource string) {
	fmt.Printf("Listing all %s \n", resource)

	result, err := client.Resource(schema.GroupVersionResource{
		Group:    "tekton.dev",
		Version:  version,
		Resource: resource,
	}).Namespace("default").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err)
	}
	for _, d := range result.Items {
		fmt.Printf(" * %s \n", d.GetName())
	}
	prompt()

}

func deleteResource(client *dynamic.DynamicClient, name string, version string, resource string) {
	err := client.Resource(schema.GroupVersionResource{
		Group:    "tekton.dev",
		Version:  version,
		Resource: resource,
	}).Namespace("default").Delete(context.TODO(), name, metav1.DeleteOptions{})

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Custom Resource %s (%s) has been deleted successfully \n", name, resource)
	}
	prompt()

}

func createClusterTask(client *dynamic.DynamicClient, obj *unstructured.Unstructured, version string, resource string) {
	result, err := client.Resource(schema.GroupVersionResource{
		Group:    "tekton.dev",
		Version:  version,
		Resource: resource,
	}).Create(context.TODO(), obj, metav1.CreateOptions{})

	if err != nil {
		fmt.Println("Error creating custom resource", err)
	} else {
		fmt.Println("Custom resource created successfully ", result)
	}
	prompt()

}

func listClusterTasks(client *dynamic.DynamicClient, version string, resource string) {
	fmt.Printf("Listing all %s \n", resource)

	result, err := client.Resource(schema.GroupVersionResource{
		Group:    "tekton.dev",
		Version:  version,
		Resource: resource,
	}).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err)
	}
	for _, d := range result.Items {
		fmt.Printf(" * %s \n", d.GetName())
	}
	prompt()

}

func deleteClusterTask(client *dynamic.DynamicClient, name string, version string, resource string) {
	err := client.Resource(schema.GroupVersionResource{
		Group:    "tekton.dev",
		Version:  version,
		Resource: resource,
	}).Delete(context.TODO(), name, metav1.DeleteOptions{})

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Custom Resource %s (%s) has been deleted successfully \n", name, resource)
	}
	prompt()

}
