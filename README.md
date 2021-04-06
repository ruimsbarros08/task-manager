# Task Manager

## Development minikube environment

### Requirements
- [Docker](https://www.docker.com/products/docker-desktop)
- [Kubernetes](https://kubernetes.io/docs/tasks/tools/)
- [Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/)

### Build image
```shell script
$ docker build . -t task-manager:latest
```

### Start minikube
```shell script
$ eval $(minikube -p minikube docker-env)
$ minikube start --mount-string="/path/to/project:/app" --mount 
```

### Setup db
```shell
$ kubectl exec -n task-manager mysql-0 -- mysql -uroot --execute="CREATE DATABASE task_manager;"
```

### Apply resources
```shell script
$ kubectl apply -k k8s/base
```

### Expose api
```shell script
$ kubectl port-forward service/api 8080:8080 -n task-manager
```

### Test
```shell
$  docker run --rm --env CGO_ENABLED=0 task-manager:latest go test ./... 
```
