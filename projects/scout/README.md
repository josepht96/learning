# Scout

## Purpose

The purpose of this project is to create a service that continuously probes kubernetes nodes to provide insight into network reliability, restrictions, latency, and reliability.

I have observed situations where applications running on k8s behave in unexpected ways, with my suspicion being there are network anomolies occuring. These sorts of issues are very hard to detect, especially working in unfamiliar environments. Scout should provide insight into network performance regardless of the kubernetes or CNI implementation.

## Components

Scout will be deployed via a daemonset. Each pod should listen for requests from remote containers, while continously probing the same remote containers. Scout must determine the IPs/pod names of its associated containers.

kubectl proxy --port=8081

## Example output

```
2023/10/27 20:12:58 probing: scout-66dd96cbb4-2vxrk -> scout-66dd96cbb4-qcnlc.default.10.244.1.75 @ node: kind-worker
2023/10/27 20:12:58     connection start: 2023-10-27 20:12:58.828585551 +0000 UTC m=+255.410049993
2023/10/27 20:12:58     connection establish: 2023-10-27 20:12:58.828809176 +0000 UTC m=+255.410273618
2023/10/27 20:12:58     latency connection: 223.625µs
2023/10/27 20:12:58     latency firstByte: 3.066584ms
2023/10/27 20:12:58     latency total: 3.17425ms
2023/10/27 20:12:58     response: {"status":"OK","statusCode":200,"body":{"message":"connected to scout","time":"2023-10-27 20:12:58.830873134 +0000 UTC m=+260.955084662"}}
```
