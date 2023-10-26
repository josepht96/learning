# Scout

## Purpose

The purpose of this project is to create a service that continuously probes kubernetes nodes to provide insight into network reliability, restrictions, latency, and reliability.

I have observed situations where applications running on k8s behave in unexpected ways, with my suspicion being there are network anomolies occuring. These sorts of issues are very hard to detect, especially working in unfamiliar environments. Scout should provide insight into network performance regardless of the kubernetes or CNI implementation.

## Components

Scout will be deployed via a daemonset. Each pod should listen for requests from remote containers, while continously probing the same remote containers. Scout must determine the IPs/pod names of its associated containers.

kubectl proxy --port=8081
