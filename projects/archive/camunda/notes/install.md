# camunda install for kind

docker run -d --name zeebe -p 26500-26502:26500-26502 camunda/zeebe:latest
docker run -d --name camunda -p 8080:8080 camunda/camunda-bpm-platform:latest

docker run -e ZEEBE_CLIENT_ID="test-auth" \
-e ZEEBE_CLIENT_SECRET="d6apt9pQkI6666bC17YiH3Jjp8mBEy6w" \
-e ZEEBE_AUTHORIZATION_SERVER_URL="http://localhost:30000/realms/master/protocol/openid-connect/token" \
-e ZEEBE_ADDRESS="0.0.0.0:30001" josepht96/zeebe-worker:latest



helm install camunda camunda/camunda-platform -f helm/camunda-platform/camunda-platform-core-kind-values.yaml --version 8.1.5