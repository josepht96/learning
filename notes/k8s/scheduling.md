# Scheduling
# cmds
kubectl taint nodes <node-name> key=value:taint-effect
# Pod nodes
Pod definitions don't contain a nodeName field (generally).
Scheduler determines which node is best for the pod. nodeName is set to the determined node
Binding object is created and assigned to the pod object. nodeName can only be set at pod creation time, cannot be modifed once the pod is created.

# Labels & Selectors
Assign labels to resources
Select resources (selectors) via labels
Can add labels to deployments/replica sets and to the actual resources
```yaml
spec:
    replicas: 3
    selector:
        matchLabels:
            app: App1
    template:
        metadata:
            labels:
                app: App1
                function: Backend
        spec:
            containers:
            - name: app
              image: app/v1 
```
# Taints and tolerations
Used to restrict which nodes pods can be assigned to
Taints go on nodes, tolerations go on pods
What happens if you overload a node because of taints?
Tolerations: NoSchedule, PreferNoSchedule, NoExecute (evicts existing pods on the node if they dont meet the requirement). 
Only used for restrictions, not for designation. Use node affinity if you want to designatem.
```yaml
    tolerations: 
    - key: app
      operator: "Equal"
      value: blue
      effect: NoSchedule
```
# Node selector
set selector on resource file that matches a label on a node
requiredDuringSchedulingIgnoreDuringExecution
preferredDuringSchedulingIgnoreDuringExecution
and 
requiredDuringSchedulingRequiredDuringExecution
preferredDuringSchedulingRequiredDuringExecution

```yaml
nodeSelector:
    size: Large
```
# Node affinity
```yaml
affinity:
    nodeAffinity:
        requiredDuringSchedulingIgnoreDuringExecution:
            nodeSelectorTerms:
            - matchExpression:
              - key: size
                operator: In
                values:
                - Large
                - Medium
```
```yaml
affinity:
    nodeAffinity:
        requiredDuringSchedulingIgnoreDuringExecution:
            nodeSelectorTerms:
            - matchExpression:
              - key: size
                operator: NotIn
                values:
                - Small
```
