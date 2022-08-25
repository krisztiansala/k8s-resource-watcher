package main

import (
	"fmt"
	"net/http"

	"github.com/krisztiansala/k8s-resource-watcher/internal/kube"
	log "github.com/sirupsen/logrus"
)

type resourceHandler struct {
	kubeClient *kube.KubeClient
}

func (h *resourceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	labelSelector := r.URL.Query().Get("pod-label")
	result, err := h.kubeClient.GetContainerResources(labelSelector)
	if err != nil {
		log.Errorln(err)
	}
	w.Write(result)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}
