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

## Cluster upgrade process

kubeadm will install control plane components
it will not install kubelet or kubectl
kubelet is what actually creates containers on nodes
kubectl is CLI

### cmds

```text
k get nodes
kubeadm upgrade plan
kubectl -n kube-system get cm kubeadm-config -oyaml
k drain *node*
k cordon *node*
k uncordon *node*
apt update
apt install kubeadm=1.20.0-00
kubeadm upgrade apply v1.20.0
apt install kubelet=1.20.0-00
systemctl restart kubelet

# kubeadm must be run on every node
# for non controlplane nodes use:

kubeadm upgrade node
```

### etcd upgrade

etcd is a distributed database of key value stores. It can run as a single static
pod, but can also be deployed in multiple pods for HA

review more

```text
# take backup

etcdctl --endpoints=https://127.0.0.1:2379 \
--cacert=<trusted-ca-file> --cert=<cert-file> --key=<key-file> \
snapshot save <backup-file-location>

ETCDCTL_API=3 etcdctl --endpoints=https://127.0.0.1:2379 --cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/server.crt --key=/etc/kubernetes/pki/etcd/server.key snapshot save /opt/snapshot-pre-boot.db

# restore

ETCDCTL_API=3 etcdctl --endpoints 10.2.0.9:2379 snapshot restore snapshotdb

1. Get etcdctl utility if it's not already present.
go get github.com/coreos/etcd/etcdctl

2. Backup
ETCDCTL_API=3 etcdctl --endpoints=https://[127.0.0.1]:2379 --cacert=/etc/kubernetes/pki/etcd/ca.crt \
     --cert=/etc/kubernetes/pki/etcd/server.crt --key=/etc/kubernetes/pki/etcd/server.key \
          snapshot save /opt/snapshot-pre-boot.db

          -----------------------------

          Disaster Happens

          -----------------------------
3. Restore ETCD Snapshot to a new folder
ETCDCTL_API=3 etcdctl --endpoints=https://[127.0.0.1]:2379 --cacert=/etc/kubernetes/pki/etcd/ca.crt \
     --name=master \
     --cert=/etc/kubernetes/pki/etcd/server.crt --key=/etc/kubernetes/pki/etcd/server.key \
     --data-dir /var/lib/etcd-from-backup \
     --initial-cluster=master=https://127.0.0.1:2380 \
     --initial-cluster-token etcd-cluster-1 \
     --initial-advertise-peer-urls=https://127.0.0.1:2380 \
     snapshot restore /opt/snapshot-pre-boot.db

 4. Modify /etc/kubernetes/manifests/etcd.yaml
 Update --data-dir to use new target location
 --data-dir=/var/lib/etcd-from-backup

 Update new initial-cluster-token to specify new cluster
 --initial-cluster-token=etcd-cluster-1

 Update volumes and volume mounts to point to new path
      volumeMounts:
          - mountPath: /var/lib/etcd-from-backup
            name: etcd-data
          - mountPath: /etc/kubernetes/pki/etcd
            name: etcd-certs
   hostNetwork: true
   priorityClassName: system-cluster-critical
   volumes:
   - hostPath:
       path: /var/lib/etcd-from-backup
       type: DirectoryOrCreate
     name: etcd-data
   - hostPath:
       path: /etc/kubernetes/pki/etcd
       type: DirectoryOrCreate
     name: etcd-certs

```

### kubeadm install components

```text
apt-get update && apt-get install -y apt-transport-https curl
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
cat <<EOF >/etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
EOF
apt-get update
```

### kubeadm install latest

```text
apt-get update
apt-get install -y kubelet kubeadm kubectl
apt-mark hold kubelet kubeadm kubectl
```

### kubeadm install specific versions

```text
apt-get install \
  kubelet=1.14.2-00 \
  kubeadm=1.14.2-00 \
  kubectl=1.14.2-00 \
  kubernetes-cni=0.7.5-00
apt-mark hold kubelet kubeadm kubectl
```

### kubeadm get versions and errors

```text
kubelet --version
kubeadm version
kubectl version
dpkg --list |grep kubernetes-cni
```

## Netowkring

Kubernetes IP addresses exist at the Pod scope - containers within a Pod share their network namespaces - including their IP address and MAC address. This means that containers within a Pod can all reach each other's ports on localhost. This also means that containers within a Pod must coordinate port usage, but this is no different from processes in a VM. This is called the "IP-per-pod" model.

Services and pods get DNS records. When a pod needs to reach service, look up the DNS record, return IP address and route network traffic to that IP address (using TCP I assume?)

check ip/network interfaces
ip a | grep -B2 1.2.3.4
ip link show eth0
arp node01
ip route show default
netstat -nplt
netstat -anp | grep etcd
ps -aux | grep <service>
Look up CNI's - very confused
