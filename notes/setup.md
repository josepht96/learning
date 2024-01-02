# setup

sudo apt install openssh-client openssh-server -y
ssh-keygen

## git
sudo apt install git
git config --global alias.ac '!git add -A && git commit'
git config --global user.email "joebthomas4@gmail.com"
git config --global user.name "josepht96"

## helm
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
chmod 700 get_helm.sh
./get_helm.sh

## docker

bash docker.sh
sudo groupadd docker && \
sudo usermod -aG docker $USER && \
newgrp docker

## go

scp go1.21.5.linux-amd64.tar.gz joe@linux:/home/joe
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

## kind

bash kind.sh
kind create cluster --config='./config.yaml'

## kubectl 
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"

## istio
curl -L https://istio.io/downloadIstio | sh -
export PATH=$PWD/bin:$PATH
istioctl install -y