# CKS

# Attack surface
zsh port-scan
Cloud > cluster > container > code

kubectl api-resources --namespaced=true/false
# CLuster setup and hardening
CIS benchmark - series of standard benchmarks for security. Best practices for security
Physical access, user access, network access, service access, filesystem, audting, logging.
CIS benchmarking tool describes these security measures. Center for internet security, has guides for cloud, OS, networks, etc...
https://www.cisecurity.org/

# Security primitives
Authentication - 
k8s doesnt have id management, uses external provider.


# Service accounts
Used by services...
Each namespace has its own SA
SA is automatically mounted to pod in that namespace
cannot edit SA of an active pod, have to delete and recreate

# TLS
Guarantee trust between two parties. 
public: *.crt, *.pem
private: *.key *-key.pem
services use their own key pairs:
api server 
etcd
kubelet

clients use kube-api server with their own key pairs
kubectl, rest api
kube-scheduler
kube-controller manager
kube-proxy

# CSR Certificate Signing Request
A CertificateSigningRequest (CSR) resource is used to request that a certificate be signed by a denoted signer, after which the request may be approved or denied before finally being signed.

Kube proxy is used for pod networking, kubectl proxy is to proxy kube-api server

# Authorization
Once authenticated, authorization determines what a user can do
node, ABAC, RBAC, webhook
Authorization is priority based, checks based on order set. Deny > deny > approved, etc...

# RBAC

```yaml
# Create role
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: default
  name: pod-reader
rules:
- apiGroups: [""] # "" indicates the core API group
  resources: ["pods"]
  verbs: ["get", "watch", "list"]

---
# bind user
apiVersion: rbac.authorization.k8s.io/v1
# This role binding allows "jane" to read pods in the "default" namespace.
# You need to already have a Role named "pod-reader" in that namespace.
kind: RoleBinding
metadata:
  name: read-pods
  namespace: default
subjects:
# You can specify more than one "subject"
- kind: User
  name: jane # "name" is case sensitive
  apiGroup: rbac.authorization.k8s.io
roleRef:
  # "roleRef" specifies the binding to a Role / ClusterRole
  kind: Role #this must be Role or ClusterRole
  name: pod-reader # this must match the name of the Role or ClusterRole you wish to bind to
  apiGroup: rbac.authorization.k8s.io
```

# Cluster roles
For nodes and cluster wide resources
Can grant users/groups permissions to specific resources cluster wide

# ssh
Client use private key, server holds public client keys at ~/.ssh/authorized_keys
Known hosts contains the IP addresses of server names:
 A file associated with a specific account that contains one or more host keys. Each host key is associated with an SSH server address (IP or hostname) so that the server can be authenticated when a connection is initiated.
The known_hosts file lets the client authenticate the server, to check that it isn't connecting to an impersonator. The authorized_keys file lets the server authenticate the user.  
# System hardening
User accounts - access to cluster, admins, devs
Super user - root, UID=0, contorl over other users
Systema accounts - created when nodes created, ssh, mail
Remove users from node system, regular linux user management

To add SSH key to node, add

# cmds
id <user>
cat /etc/passwd
userdel <user>
usermod -s /usr/sbin/nologin himanshi
useradd -d /opt/sam -s /bin/bash -G admin -u 2328 sam
ssh-copy-id -i ~/.ssh/id_rsa.pub jim@node01
# remove unwanted services