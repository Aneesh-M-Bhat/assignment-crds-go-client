# assignment-crds-go-client

To Setup CRDs in cluster

```
kubectl apply --filename https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml && kubectl get pods --namespace tekton-pipelines --watch
```

To Create a config file in project directory

```
cat ~/.kube/config >> config.yaml
```

To run the go project

```
go run *.go
```
