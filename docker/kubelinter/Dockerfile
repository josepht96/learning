FROM stackrox/kube-linter:v0.6.1-alpine

RUN apk update && \
	apk add curl openssl bash

RUN curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 && \
	chmod 700 get_helm.sh && \
	./get_helm.sh

RUN cp /kube-linter /usr/local/bin

ENTRYPOINT ["/bin/sh", "-c", "tail -f /dev/null"]

