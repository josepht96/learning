# CKAD

## Architecture
Installing Kubernetes you install:
API Server - frontend for k8s, cli use it to communicate with cluster
etcd - key value store to store data about the cluster
kubelet - agent that runs on each node, responsible for maintaining running containers on nodes
Scheduler - schedule workloads to nodes
Controller - notice and respond when things are off desired state
Container Runtime - underlying software used to run containers

