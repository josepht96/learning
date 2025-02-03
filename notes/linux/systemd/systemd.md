systemctl list-units --type=service
systemctl list-unit-files
systemctl cat docker.service

docker run --name nginx-startup -d -p 8080:80 nginx

## location of system installed service config files
ls -al  /lib/systemd/system

## location for user created config files
/etc/systemd/system
/etc/systemd/system/nginx-startup.service

##
systemctl start nginx-startup
systemctl enable nginx-startup (to enable on startup)