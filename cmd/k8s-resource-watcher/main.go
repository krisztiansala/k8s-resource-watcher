package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/krisztiansala/k8s-resource-watcher/internal/logging"
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

const (
	// listenAddress = "0.0.0.0"
	listenAddress = "127.0.0.1"
	port          = 8000
)

func main() {
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	log.Printf("Using kubeconfig file: %s", kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/container-resources", &resourceHandler{clientset: clientset})
	http.HandleFunc("/", RootHandler)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", listenAddress, port),
		Handler: logging.WithLogging(http.DefaultServeMux),
	}
	log.Printf("Listening on %s:%d", listenAddress, port)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Error initializing server: ", err)
		}
	}()

	<-ctx.Done()
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	log.Info("Server shutting down gracefully...")

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forcefully shut down: ", err)
	}

}
