java -javaagent:opentelemetry-javaagent.jar \
     -Dotel.traces.exporter=otlp \
     -Dotel.exporter.otlp.endpoint=http://localhost:4317 \
     -Dotel.service.name=otel-demo \
     -jar target/otel-demo-0.0.1-SNAPSHOT.jar

mvn clean package   # First, package the app as a JAR
java -jar target/otel-demo-0.0.1-SNAPSHOT.jar


java -javaagent:opentelemetry-javaagent.jar \
     -jar target/otel-demo-0.0.1-SNAPSHOT.jar

-Dotel.javaagent.debug=true

export OTEL_TRACES_EXPORTER=otlp
export OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:30002
export OTEL_METRICS_EXPORTER=none
export OTEL_LOGS_EXPORTER=none
export OTEL_SERVICE_NAME=otel-demo
export OTEL_JAVAAGENT_DEBUG=true


export SERVER_PORT=8081
export OTEL_TRACES_EXPORTER=otlp
export OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:30002
export OTEL_METRICS_EXPORTER=none
export OTEL_LOGS_EXPORTER=none
export OTEL_SERVICE_NAME=otel-demo-2
export OTEL_JAVAAGENT_DEBUG=true


-Dotel.traces.exporter=otlp \
     -Dotel.exporter.otlp.endpoint=http://localhost:4317 \

docker build -t josepht96/otel-demo:linux --platform linux/amd64 .