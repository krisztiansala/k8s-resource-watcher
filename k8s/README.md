Use these resources to quickly deploy the service to a cluster (and you don't have Helm installed).  
The deployment will use the default service account, which will need viewer permission to the cluster resources to access the necessary data.  
Run the following command to assign the necessary permission:  
`kubectl create clusterrolebinding default-view --clusterrole=view --serviceaccount=default:default`