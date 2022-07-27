# CKAD

## Architecture

Installing Kubernetes you install:
API Server - frontend for k8s, cli use it to communicate with cluster
etcd - key value store to store data about the cluster
kubelet - agent that runs on each node, responsible for maintaining running containers on nodes
Scheduler - schedule workloads to nodes
Controller - notice and respond when things are off desired state
Container Runtime - underlying software used to run containers

## Overview

Core concepts: 13%
Configuration: 18%
Multi-container pods: 10%
Observability: 18%
Pod design: 20%
Service and networking: 13%
State persistence: 8%

## Core concepts

Deployment - manages deployment of multiple pods of the same spec
Stateful set - manages deployment of multiple pods where state is important. Pods are not interchangable. Makes it easier to match PVs to specific pods.
Daemon set - deploys pod to each node

The Pod remains on that node until the Pod finishes execution, the Pod object is deleted, the Pod is evicted for lack of resources, or the node fails. 
Pod updates may not change fields other than spec.containers[*].image, spec.initContainers[*].image, spec.activeDeadlineSeconds or spec.tolerations. For spec.tolerations, you can only add new entries.

Each Pod is assigned a unique IP address for each address family. Every container in a Pod shares the network namespace, including the IP address and network ports. Inside a Pod (and only then), the containers that belong to the Pod can communicate with one another using localhost. When containers in a Pod communicate with entities outside the Pod, they must coordinate how they use the shared network resources (such as ports). Within a Pod, containers share an IP address and port space, and can find each other via localhost. The containers in a Pod can also communicate with each other using standard inter-process communications like SystemV semaphores or POSIX shared memory. Containers in different Pods have distinct IP addresses and can not communicate by OS-level IPC without special configuration. Containers that want to interact with a container running in a different Pod can use IP networking to communicate.

Containers within the Pod see the system hostname as being the same as the configured name for the Pod.

## Probing

### probe mechanisms

**exec**
Executes a specified command inside the container. The diagnostic is considered successful if the command exits with a status code of 0.

**httpGet**
Performs an HTTP GET request against the Pod's IP address on a specified port and path. The diagnostic is considered successful if the response has a status code greater than or equal to 200 and less than 400.

**tcpSocket**
Performs a TCP check against the Pod's IP address on a specified port. The diagnostic is considered successful if the port is open. If the remote system (the container) closes the connection immediately after it opens, this counts as healthy.

### probe types

**livenessProbe**
Indicates whether the container is running. If the liveness probe fails, the kubelet kills the container, and the container is subjected to its restart policy. If a container does not provide a liveness probe, the default state is Success.

**readinessProbe**
Indicates whether the container is ready to respond to requests. If the readiness probe fails, the endpoints controller removes the Pod's IP address from the endpoints of all Services that match the Pod. The default state of readiness before the initial delay is Failure. If a container does not provide a readiness probe, the default state is Success.

**startupProbe**
Indicates whether the application within the container is started. All other probes are disabled if a startup probe is provided, until it succeeds. If the startup probe fails, the kubelet kills the container, and the container is subjected to its restart policy. If a container does not provide a startup probe, the default state is Success.

When a liveness probe fails, it signals to OpenShift that the probed container is dead and should be restarted. When a readiness probe fails, it indicates to OpenShift that the container being probed is not ready to receive incoming network traffic. The application might become ready in the future, but it should not receive traffic now.

### memory limits

requests cannot exceed limits. Controller will terminate pod if it exceeds limits.

The CPU resource is measured in CPU units. One CPU, in Kubernetes, is equivalent to:

1 AWS vCPU
1 GCP Core
1 Azure vCore
1 Hyperthread on a bare-metal Intel processor with Hyperthreading

PV to PVC is always one to one.
Projected volume maps several volumes into the same directory

### security context

Discretionary Access Control: Permission to access an object, like a file, is based on user ID (UID) and group ID (GID).

Security Enhanced Linux (SELinux): Objects are assigned security labels.

Running as privileged or unprivileged.

Linux Capabilities: Give a process some privileges, but not all the privileges of the root user.

AppArmor: Use program profiles to restrict the capabilities of individual programs.

Seccomp: Filter a process's system calls.

allowPrivilegeEscalation: Controls whether a process can gain more privileges than its parent process. This bool directly controls whether the no_new_privs flag gets set on the container process. allowPrivilegeEscalation is always true when the container:

is run as privileged, or
has CAP_SYS_ADMIN
readOnlyRootFilesystem: Mounts the container's root filesystem as read-only.

### service accounts

The service account has to exist at the time the pod is created, or it will be rejected.

You cannot update the service account of an already created pod.

OIDC: [link](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#service-account-token-volume-projection)

### private docker registry

Pods can pull images from private repos by the following:
`kubectl create secret docker-registry regcred --docker-server=<your-registry-server> --docker-username=<your-name> --docker-password=<your-pword> --docker-email=<your-email>`

which yields:

```yaml
apiVersion: v1
kind: Secret
metadata:
  ...
  name: regcred
  ...
data:
  .dockerconfigjson: eyJodHRwczovL2luZGV4L ... J0QUl6RTIifX0=
type: kubernetes.io/dockerconfigjson
```

Referenced via:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: private-reg
spec:
  containers:
  - name: private-reg-container
    image: <your-private-image>
  imagePullSecrets:
  - name: regcred

```

### Lifecycle events

Kubernetes supports the postStart and preStop events. Kubernetes sends the postStart event immediately after a Container is started, and it sends the preStop event immediately before the Container is terminated. A Container may specify one handler per event. postStart runs after container initialization, after init containers. No guarantee it runs before container image entrypoint cmd.

### Annotations

You can use either labels or annotations to attach metadata to Kubernetes objects. Labels can be used to select objects and to find collections of objects that satisfy certain conditions. In contrast, annotations are not used to identify and select objects. The metadata in an annotation can be small or large, structured or unstructured, and can include characters not permitted by labels.

### Deployments

`kubectl rollout history deployment/nginx-deployment --revision=2`
`kubectl rollout undo deployment/nginx-deployment --to-revision=2`
`kubectl scale deployment/nginx-deployment --replicas=10`
s
A Deployment's revision history is stored in the ReplicaSets it controls.

.spec.revisionHistoryLimit is an optional field that specifies the number of old ReplicaSets to retain to allow rollback. These old ReplicaSets consume resources in etcd and crowd the output of kubectl get rs. The configuration of each Deployment revision is stored in its ReplicaSets; therefore, once an old ReplicaSet is deleted, you lose the ability to rollback to that revision of Deployment. By default, 10 old ReplicaSets will be kept, however its ideal value depends on the frequency and stability of new Deployments.

More specifically, setting this field to zero means that all old ReplicaSets with 0 replicas will be cleaned up. In this case, a new Deployment rollout cannot be undone, since its revision history is cleaned up.
