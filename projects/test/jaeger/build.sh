docker build -t josepht96/jaeger-test:v0.0.1 .
docker push josepht96/jaeger-test:v0.0.1
kubectl delete -f deploy.yaml
kubectl apply -f deploy.yaml