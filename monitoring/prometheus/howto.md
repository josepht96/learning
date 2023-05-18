prometheus --config.file=config.yaml
node_exporter


prometheus runs in container
other apps run in container
node_exporter runs on host


## run prometheus as container
docker run  -d\
    -p 9092:9090 \
    -v /Users/joe/repos/learning/monitoring/prometheus/config-docker.yaml:/etc/prometheus/prometheus.yml \
    prom/prometheus