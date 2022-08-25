# k8s-resource-watcher

This repository contains a simple web server implemented in Go, that can retrieve the resource requests and limits of containers in a Kubernetes cluster.  
It can be run locally - in that case the authentication is done by accessing the `~/.kube/config` file (and taking the credentials from there).  
It can be deployed to a Kubernetes cluster as well, in that case it relies on the deployment's service account for authentication.
## Running the app locally

Before running, first create a build of the application. This can be done with the `make build` command (or `make all` to run the unit tests befgore the build).

The resulting binary can be found in the `bin/` directory and can be run like: `bin/k8s-resource-watcher`.

The environment - `dev` or `prod` - can be specified using the `ENV` environmental variable. This will decide which authentication method will be used and on which address the server is listening.

## Deploying to Kubernetes

A basic version of the application can be deployed with the `kubectl` command, following the instructions from the `k8s` folder.

Using the Helm chart from the `helm` directory hovewer provides more configuration options and it also creates a service account (which provides the necessary permissions for accessing pod details in the cluster) and a HPA (to scale the deployment if the load increases).

To deploy the chart, run the following command from the `helm` directory: `helm upgrade --install k8s-resource-watcher --values values.yaml .`

The default values can be customized by changing them in the `helm/values.yaml` file.

The created service is of type ClusterIP, so it won't be available publicly. That's why the `make portforward` command can be run if we wish to access the application locally.

## Create local cluster

To create a local Kubernetes cluster and deploy the application on it with one command, run `make tf_local_apply`. This will spin up a new `k3d` cluster (needs `Docker`) and installs the Helm chart on it.
Destroy the above created cluster with `make tf_local_destroy`.
