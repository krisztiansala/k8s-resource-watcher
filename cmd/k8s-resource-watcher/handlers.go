package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
)

type resourceHandler struct {
	clientset *kubernetes.Clientset
}

func (h *resourceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	labelSelector := r.URL.Query().Get("pod-label")
	result, err := GetContainerResources(h.clientset, labelSelector)
	if err != nil {
		log.Errorln(err)
	}
	w.Write(result)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}
