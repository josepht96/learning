kubectl port-forward svc/camunda-keycloak 18080:80 --address=0.0.0.0 & \
kubectl port-forward svc/camunda-identity 8080:80 --address=0.0.0.0 & \
kubectl port-forward svc/camunda-operate 8081:80 --address=0.0.0.0 & \
kubectl port-forward svc/camunda-tasklist 8082:80 --address=0.0.0.0 & \
kubectl port-forward svc/camunda-optimize 8083:80 --address=0.0.0.0 & \
kubectl port-forward svc/camunda-zeebe-gateway 26500:26500 --address=0.0.0.0 & 


# k get secret camunda-keycloak -ojsonpath='{.data.admin-password}' | base64 -d