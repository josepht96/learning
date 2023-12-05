docker build -t josepht96/jaeger-test:latest .
docker push josepht96/jaeger-test:latest
kubectl delete -f deployment.yaml
kubectl apply -f deployment.yaml