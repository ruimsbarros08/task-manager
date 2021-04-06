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

### Apply resources
```shell script
$ kubectl apply -f k8s/namespace.yml
$ kubectl apply -f k8s/redis.yml
```

### Setup db
```shell
$ kubectl apply -f k8s/mysql.yml
$ kubectl exec -n task-manager mysql-0 -- mysql -uroot --execute="CREATE DATABASE task_manager;"
```

### Apply api and consumer
```shell script
$ kubectl apply -f k8s/api.yml
$ kubectl apply -f k8s/consumer.yml
```

### Expose api
```shell script
$ kubectl port-forward service/api 8080:8080 -n task-manager
```

### Test
```shell
$  docker run --rm --env CGO_ENABLED=0 task-manager:latest go test ./... 
```
