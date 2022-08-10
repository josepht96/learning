# CKA

[API reference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#-strong-workloads-apis-strong-)
Control plane:
    Scheduler
    etcd
    API server
    Controller manager
    kubelet
    kube-proxy

# etcd
Kubernetes uses etcd as a key-value database store. It stores the configuration of the Kubernetes cluster in etcd.

It also stores the actual state of the system and the desired state of the system in etcd.

It then uses etcd’s watch functionality to monitor changes to either of these two things. If they diverge, Kubernetes makes changes to reconcile the actual state and the desired state.

# kube-controller-manager
Control plane component that runs controller processes.

Logically, each controller is a separate process, but to reduce complexity, they are all compiled into a single binary and run in a single process.

Some types of these controllers are:

Node controller: Responsible for noticing and responding when nodes go down.
Job controller: Watches for Job objects that represent one-off tasks, then creates Pods to run those tasks to completion.
Endpoints controller: Populates the Endpoints object (that is, joins Services & Pods).
Service Account & Token controllers: Create default accounts and API access tokens for new namespaces.


# kubelet 
An agent that runs on each node in the cluster. It makes sure that containers are running in a Pod.

The kubelet takes a set of PodSpecs that are provided through various mechanisms and ensures that the containers described in those PodSpecs are running and healthy. The kubelet doesn't manage containers which were not created by Kubernetes.

# kube-proxy
kube-proxy is a network proxy that runs on each node in your cluster, implementing part of the Kubernetes Service concept.

kube-proxy maintains network rules on nodes. These network rules allow network communication to your Pods from network sessions inside or outside of your cluster.

kube-proxy uses the operating system packet filtering layer if there is one and it's available. Otherwise, kube-proxy forwards the traffic itself.

## kubeadm

provision machines
designate nodes
install container runtimes on all nodes
install kubeadm on all nodes
initialize master server
ensure network connectivity
implement pod network
join worker nodes to master nodes

## Managing host OS

When a node goes offline, control plane determines how long before evicting pods on that node. Pods will be replicated to new nodes, assuming theyre part of a deployment.
kubeclt drain node-name, removes pods from a nod. Recreated on different node
Node is cordoned which means no pods can be scheduled to it.
Node must be uncordoned once node comes back.

k get pods -o wide
