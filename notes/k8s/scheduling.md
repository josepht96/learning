# Scheduling

# Pod nodes
Pod definitions don't contain a nodeName field (generally).
Scheduler determines which node is best for the pod. nodeName is set to the determined node
Binding object is created and assigned to the pod object. nodeName can only be set at pod creation time, cannot be modifed once the pod is created.

# Labels & Selectors
Assign labels to resources
Select resources (selectors) via labels
