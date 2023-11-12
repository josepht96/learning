# Scout

## Purpose

The purpose of this project is to create a service that continuously probes kubernetes nodes to provide insight into network reliability, restrictions, latency, and reliability.

I have observed situations where applications running on k8s behave in unexpected ways, with my suspicion being there are network anomolies occuring. These sorts of issues are very hard to detect, especially working in unfamiliar environments. Scout should provide insight into network performance regardless of the kubernetes or CNI implementation.

## Components

Scout will be deployed via a daemonset. Each pod should listen for requests from remote containers, while continously probing the same remote containers. Scout must determine the IPs/pod names of its associated containers.

kubectl proxy --port=8081

## Example output

```
2023/10/30 17:09:49 server is listening at http://localhost:8080
2023/10/30 17:09:49 scout pods:
2023/10/30 17:09:49     node: kind-worker
2023/10/30 17:09:49             scout-66dd96cbb4-rf6z9.default.10.244.1.30
2023/10/30 17:09:49     node: kind-worker
2023/10/30 17:09:49             scout-66dd96cbb4-rg8d8.default.10.244.1.31
2023/10/30 17:09:49     node: kind-worker
2023/10/30 17:09:49             scout-66dd96cbb4-wwc5z.default.10.244.1.32
2023/10/30 17:09:49 probing: localhost -> scout-66dd96cbb4-rf6z9.default.10.244.1.30 @ node: kind-worker
2023/10/30 17:09:49     connection start: 2023-10-30 17:09:49.89942 -0500 CDT m=+0.046433751
2023/10/30 17:09:49     latency dns: 866.5µs
2023/10/30 17:09:49     latency connection: 289.542µs
2023/10/30 17:09:49     latency write request: 43.959µs
2023/10/30 17:09:49     latency server processing: 264.25µs
2023/10/30 17:09:49     latency content transfer: 21.75µs
2023/10/30 17:09:49     latency total: 1.548375ms
2023/10/30 17:09:49     response: {"status":"OK","statusCode":200,"body":{"message":"connected to scout","time":"2023-10-30 17:09:49.900801 -0500 CDT m=+0.047814418"}}
2023/10/30 17:09:49 probing: localhost -> scout-66dd96cbb4-rg8d8.default.10.244.1.31 @ node: kind-worker
2023/10/30 17:09:49     connection start: 2023-10-30 17:09:49.901013 -0500 CDT m=+0.048026251
2023/10/30 17:09:49     latency dns: 212.25µs
2023/10/30 17:09:49     latency connection: 148.625µs
2023/10/30 17:09:49     latency write request: 12.75µs
2023/10/30 17:09:49     latency server processing: 127.792µs
2023/10/30 17:09:49     latency content transfer: 13.291µs
2023/10/30 17:09:49     latency total: 554.25µs
2023/10/30 17:09:49     response: {"status":"OK","statusCode":200,"body":{"message":"connected to scout","time":"2023-10-30 17:09:49.901505 -0500 CDT m=+0.048519168"}}
2023/10/30 17:09:49 probing: localhost -> scout-66dd96cbb4-wwc5z.default.10.244.1.32 @ node: kind-worker
2023/10/30 17:09:49     connection start: 2023-10-30 17:09:49.901598 -0500 CDT m=+0.048611918
2023/10/30 17:09:49     latency dns: 210.916µs
2023/10/30 17:09:49     latency connection: 154.542µs
2023/10/30 17:09:49     latency write request: 15.709µs
2023/10/30 17:09:49     latency server processing: 156µs
2023/10/30 17:09:49     latency content transfer: 15µs
2023/10/30 17:09:49     latency total: 593.208µs
2023/10/30 17:09:49     response: {"status":"OK","statusCode":200,"body":{"message":"connected to scout","time":"2023-10-30 17:09:49.902103 -0500 CDT m=+0.049116751"}}
```
