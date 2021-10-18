# Kubernetes

# Kubernetes Overview 
Hardware -> OS -> Docker engine -> container
Containers have their own NICs, mounts, processes but they share OS kernel, as opposed to VMs which have their own kernel. OS is a combination of kernel and software. Linux is the kernel, software is what differentiates Ubuntu and RHEL or any other Linux variant. That's why you can run Ubuntu container on CentOS host because they both have linux kernel. 
To enable container scaling and networking, orchestration is needed, k8s, docker swarm, Mesos

# K8s architecture - All kubernetes clusters have:
API Server - frontend for users (command line)
etcd - key store, store data used to manage the cluster
Scheduler - distributing work across multiple nodes
Controller - Noticing and responding when nodes or endpoints go down, decides to bring up new pods/containers
Container runtime - container engine (docker)
kubelet - agent that runs on each node in cluster
Master vs Worker Nodes: 
master - kube-apiserver, etcd, controller, scheduler
worker - container runtime, kubelet
kubectl cli tool - used to deploy and manage applications on a cluster
kubectl cluster-info, kubectl get nodes

# cmds:
minikube -
	start
	ip
	service <name> --url
	
kubectl -
	get pods
	get deployments
run nginx --image nginx (name can be anything, image needs to be existing image
hosted somewhere)
describe pod nginx
get pods -o wide
	port-forward service/<sn> hostport:container port
	get services <sn>
	delete service <sn>
	delete deployment <deployment name>
	create deployment nginx --image=nginx
	(apply and create work the same when creating new def, apply will throw warning though
first time)
	create -f pod-definition.yml
	replace -f <file>
	scale --replicas=6 replicaset myapp-replicaset (wont update file)
	edit replicaset myapp-replicaset
	get all (list all objects)
	 kubectl rollout status deployment.apps/myapp-deployment
	rollout history
	rollout undo

Edit changes to the yaml file will be applied on save.

# Minikube - single node kubernetes cluster, packaged in iso image
Pods are the smallest manageable unit in k8s, usually pods and containers have 1:1 relationship. If more instances of the container are needed, you add additional pods with the new instances of the container, don't add to the existing pod. Helper containers are fine to run on the same pod if direct container to container communication is needed. Containers created on the same node have access to the same network namespace. 
kubectl get pods: READY = running containers/total containers

# Replication controllers and ReplicaSets’
Replication controller ensures the specified number of pods are running at all times (even if that’s 1). Replica Set is replacing Replication Controller. One difference is ReplicaSet has a selector feature. If you create a pod that matches the lab, then delete a pod from the replicaset pods, a new one won't be created (cuz 3 still match the tag). As well, if you create a new pod outside a replica set with a single matching tag, it will delete the pod.


# Networking
each node has an IP address, each pod has its own IP address as well. All created pods are defaulted to the same network. So by default pods will all contain the same network IP address, pods might contain the same IP addresses as well. Need 3rd party networking utility to make this easier. Flannel is what I used for the local cluster. 
Services 
NodePort service can handle external requests to cluster/node. Listens on a cluster port (ip:30008-32767) and forwards requests to a node. 
- NodePort: 
	TargetPort: port on the pod where application is running
	Port: port on service itself, needs match the TargetPod
	NodePort: External port on cluster that forwards to services
Connect to node port via cluster IP:NodePort
NodePort randomly chooses pods to distribute workload. Since the Service selects pods based on tags, not IPs, if a node or pod goes down and another is spawned, the Service can pick the new instance up simply because of tag. 
- ClusterIP
Used to group layers of the application together. Provides a single interface for pods to access another group of pods. Other pods can access the ClusterIP by its IP or by the service’s name.
- Load balancer 
NodePort’s help with load balancing over a single node, load balancer helps with load balancing over an entire cluster. 
Cluster Architecture
Master node - control plane. Manage, plan, schedule, monitor nodes
kube-scheduler - determines right nodes to place containers on, based on system specs and other requirements.
Controllers: 
	Controller-Manager
	Node-Controller: Add new nodes, handle node failures
Replication-Controller: Ensure desired number of containers are running at all
times
kube-apiserver is responsible for orchestrating all operations within the cluster. Exposes API for users to apply commands.
Install docker on all nodes including master, if you want to run control plane with containers
Every node has a kubelet - agent that listens for instructions for kube-apiserver
periodically communicates with KAS for status updates
Nodes also have kube-proxy service. Necessary rules on the nodes for the containers to communicate with each other if needed. 

# ETCD
Distributed key value store. Most dbs are tabular format, ETCD has to have unique key values that map to a value. Stores info about the cluster concerning any aspect of the cluster. add new nodes will trigger an update in ETCD server. Only once its updated in ETCD is the change complete. 

# Controller Manager
A controller is a service that continuously monitors the state of a resource - will detect things like server failures, container failures, etc…

# Kubelet
# Kube-Proxy
Pod-network - all pods can communicate with each other
Kube-proxy is a process that runs on each node, its job is to look for new services and creates a rule on each node to forward traffic on those services. Forwards requests from service to pod
kubectl get pods -n kube-system

# Pods


# Networking
Devices hace NIC (each has a unique MAC address) which communicate over a switch. A router connects two or more networks together (internet to local network in most cases). Router is just a device on the network. 0.0.0.0 in route table means the device doesnt require a gateway to access, ie its on the same network as the current device. 
