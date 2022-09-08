
# Debugging stuff

nslookup

```text
joe@Joes-MacBook-Pro books % nslookup google.com
# info about DNS server
Server:  150.199.101.1 
Address: 150.199.101.1#53

# info about google.com
Non-authoritative answer:
Name: google.com
Address: 142.250.191.110
```

kubectl get endpoints *service-name*

// check if kube-proxy is running:
ps auxw | grep kube-proxy

// check logs of a specific container
kubectl logs pod-name -c init-container-2

// get cluster info
kubectl cluster-info dump
