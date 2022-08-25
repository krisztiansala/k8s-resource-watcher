provider "helm" {
  kubernetes {
    config_path    = "~/.kube/config"
    config_context = "k3d-${var.cluster_name}"
  }
}

terraform {
  required_providers {
    k3d = {
      source  = "pvotal-tech/k3d"
      version = "0.0.6"
    }
  }
}

provider "k3d" {}

resource "k3d_cluster" "mycluster" {
  name = var.cluster_name

  kubeconfig {
    update_default_kubeconfig = true
    switch_current_context    = true
  }
  # Uncomment these options if having issues with ephemeral storage
  # k3s {
  #   extra_args {
  #     arg          = "--kubelet-arg=eviction-hard=imagefs.available<1%,nodefs.available<1%"
  #     node_filters = ["server:0"]
  #   }
  # }
}

resource "helm_release" "k8s-resource-watcher" {
  name      = "k8s-resource-watcher"
  chart     = "${path.root}/../../helm"
  namespace = "default"
  depends_on = [
    k3d_cluster.mycluster
  ]
}