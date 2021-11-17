service nginx reload
nginx -c /etc/nginx/nginx.conf -t (use sudo if need be)

comment out #include /etc/nginx/sites-enabled/*;
server {
    listen 80 default_server;
    location / {
        proxy_pass http://localhost:4041;
    }
}