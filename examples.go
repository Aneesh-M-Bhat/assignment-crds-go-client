package main

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func getExampleTask() *unstructured.Unstructured {
	task := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "tekton.dev/v1beta1",
			"kind":       "Task",
			"metadata": map[string]interface{}{
				"name": "hello",
			},
			"spec": map[string]interface{}{
				"steps": []map[string]interface{}{
					{
						"name":  "echo",
						"image": "alpine",
						"script": "#!/bin/sh\n" +
							"echo \"Hello World\"",
					},
				},
			},
		},
	}

	task.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1beta1",
		Kind:    "Task",
	})
	return task
}

func getExampleTaskRun() *unstructured.Unstructured {
	taskrun := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "tekton.dev/v1beta1",
			"kind":       "TaskRun",
			"metadata": map[string]interface{}{
				"name": "hello-task-run",
			},
			"spec": map[string]interface{}{
				"taskRef": map[string]interface{}{
					"name": "hello",
				},
			},
		},
	}

	taskrun.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1beta1",
		Kind:    "TaskRun",
	})
	return taskrun
}

func getExampleTask2() *unstructured.Unstructured {
	task := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "tekton.dev/v1beta1",
			"kind":       "Task",
			"metadata": map[string]interface{}{
				"name": "goodbye",
			},
			"spec": map[string]interface{}{
				"params": []map[string]interface{}{
					{
						"name": "username",
						"type": "string",
					},
				},
				"steps": []map[string]interface{}{
					{
						"name":   "goodbye",
						"image":  "ubuntu",
						"script": "#!/bin/bash\n" + "echo \"Goodbye $(params.username)!\"",
					},
				},
			},
		},
	}

	task.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1beta1",
		Kind:    "Task",
	})
	return task
}

func getExamplePipeline() *unstructured.Unstructured {
	pipeline := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "tekton.dev/v1beta1",
			"kind":       "Pipeline",
			"metadata": map[string]interface{}{
				"name": "hello-goodbye",
			},
			"spec": map[string]interface{}{
				"params": []map[string]interface{}{
					{
						"name": "username",
						"type": "string",
					},
				},
				"tasks": []map[string]interface{}{
					{
						"name": "hello",
						"taskRef": map[string]interface{}{
							"name": "hello",
						},
					},
					{
						"name": "goodbye",
						"runAfter": []interface{}{
							"hello",
						},
						"taskRef": map[string]interface{}{
							"name": "goodbye",
						},
						"params": []map[string]interface{}{
							{
								"name":  "username",
								"value": "$(params.username)",
							},
						},
					},
				},
			},
		},
	}

	pipeline.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1beta1",
		Kind:    "Pipeline",
	})
	return pipeline
}

func getExamplePipelineRun() *unstructured.Unstructured {
	pr := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "tekton.dev/v1beta1",
			"kind":       "PipelineRun",
			"metadata": map[string]interface{}{
				"name": "hello-goodbye-run",
			},
			"spec": map[string]interface{}{
				"pipelineRef": map[string]interface{}{
					"name": "hello-goodbye",
				},
				"params": []map[string]interface{}{
					{
						"name":  "username",
						"value": "Tekton",
					},
				},
			},
		},
	}

	pr.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1beta1",
		Kind:    "PipelineRun",
	})
	return pr
}

func getExampleClusterTask() *unstructured.Unstructured {
	task := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "tekton.dev/v1beta1",
			"kind":       "ClusterTask",
			"metadata": map[string]interface{}{
				"name": "example-cluster-task",
			},
			"spec": map[string]interface{}{
				"steps": []map[string]interface{}{
					{
						"name":   "step1",
						"image":  "alpine",
						"script": "#!/bin/sh\necho \"This is step 1\"",
					},
					{
						"name":   "step2",
						"image":  "alpine",
						"script": "#!/bin/sh\necho \"This is step 2\"",
					},
				},
			},
		},
	}

	task.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1beta1",
		Kind:    "ClusterTask",
	})
	return task
}

func getExampleCustomRun() *unstructured.Unstructured {
	cr := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "tekton.dev/v1beta1",
			"kind":       "CustomRun",
			"metadata": map[string]interface{}{
				"name": "my-example-run",
			},
			"spec": map[string]interface{}{
				"customRef": map[string]interface{}{
					"apiVersion": "example.dev/v1beta1",
					"kind":       "Example",
					"name":       "an-existing-example-task",
				},
			},
		},
	}

	cr.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1beta1",
		Kind:    "ClusterTask",
	})
	return cr
}
