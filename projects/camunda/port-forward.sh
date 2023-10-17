kubectl port-forward svc/camunda-keycloak 18080:80 --address=0.0.0.0 & \
kubectl port-forward svc/camunda-operate 8081:80 --address=0.0.0.0 & \
kubectl port-forward svc/camunda-tasklist 8082:80 --address=0.0.0.0