openssl req -newkey rsa:4096 \
            -x509 \
            -sha256 \
            -days 3650 \
            -nodes \
            -out /etc/nginx/nginx.cert \
            -keyout /etc/nginx/nginx.key

k create secret generic client-credential -n istio-system --from-file=cacert=nginx.cert


## curl to port 80 using https from istio to istio pod, where dest pod is expecting http
curl: (35) OpenSSL SSL_connect: SSL_ERROR_SYSCALL in connection to nginx.off.svc.cluster.local:80 

## curl to port 80 using https from istio to non istio pod
curl: (35) error:1408F10B:SSL routines:ssl3_get_record:wrong version number