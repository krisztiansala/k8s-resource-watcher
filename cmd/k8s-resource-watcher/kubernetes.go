package main

import (
	"context"
	"encoding/json"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ContainerResources struct {
	PodName       string `json:"pod_name"`
	ContainerName string `json:"container_name"`
	Namespace     string `json:"namespace"`
	MemoryReq     string `json:"mem_req"`
	MemoryLimit   string `json:"mem_limit"`
	CPUReq        string `json:"cpu_req"`
	CPULimit      string `json:"cpu_limit"`
}

func GetContainerResources(clientset *kubernetes.Clientset, labelSelector string) ([]byte, error) {
	pods, err := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{
		LabelSelector: labelSelector,
	})
	if err != nil {
		return []byte{}, fmt.Errorf("failed to get pods: %v", err)
	}
	var containers []ContainerResources
	for _, pod := range pods.Items {
		for _, container := range pod.Spec.Containers {
			containerRes := ContainerResources{}
			containerRes.PodName = pod.GetName()
			containerRes.Namespace = pod.ObjectMeta.Namespace
			containerRes.ContainerName = container.Name
			containerRes.CPUReq = container.Resources.Requests.Cpu().String()
			containerRes.CPULimit = container.Resources.Limits.Cpu().String()
			containerRes.MemoryReq = container.Resources.Requests.Memory().String()
			containerRes.MemoryLimit = container.Resources.Limits.Memory().String()
			containers = append(containers, containerRes)
		}
	}
	result, err := json.Marshal(containers)
	if err != nil {
		return []byte{}, fmt.Errorf("error marshalling to json: %v", err)
	}
	return result, nil
}
