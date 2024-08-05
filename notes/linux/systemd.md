systemctl list-units --type=service
systemctl list-unit-files
systemctl cat docker.service

docker run --name nginx-startup -d -p 8080:80 nginx
ls -al  /lib/systemd/system