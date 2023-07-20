Binding a clusterrole using a rolebinding will limit the permissions to just that namespace
You can then bind another namespace's entities to a clusterrole scope to your namespace
```
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: access-for-a
rules:
- apiGroups: [""]
  resources: ["pods", "deployments"] # etc
  verbs: ["get", "list", "create"]
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: my-app
  namespace: a
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: access-for-a
  namespace: b
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: access-for-a
subjects:
- kind: ServiceAccount
  name: my-app
  namespace: a
```
