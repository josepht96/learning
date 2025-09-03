DTN7HSJ12PNCaLy5dNbP03m3eVqfJjQbYkBhfHTq

grafana > tempo > ??

http://tempo:3100


helm install grafana grafana/grafana
helm install tempo grafana/tempo

kubectl apply -f https://raw.githubusercontent.com/open-telemetry/opentelemetry-collector/v0.117.0/examples/k8s/otel-config.yaml
