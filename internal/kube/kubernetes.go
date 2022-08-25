package kube

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
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

type KubeClient struct {
	Clientset *kubernetes.Clientset
}

func NewKubeClient(env string) *KubeClient {
	k := new(KubeClient)
	var config *restclient.Config
	var err error
	if env == "dev" {
		kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
		log.Printf("Using kubeconfig file: %s", kubeconfig)
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		config, err = rest.InClusterConfig()
		if err != nil {
			log.Fatal(err)
		}
	}

	k.Clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	return k
}

func (k *KubeClient) GetContainerResources(labelSelector string) ([]byte, error) {
	pods, err := k.Clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{
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
