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

Network
  ☐ Direct Connect established and routed to GovCloud VPC
  ☐ On-prem subnet routes to VPC CIDR via Direct Connect

S3
  ☐ S3 Interface VPCE created in VPC
  ☐ Security group on VPCE allows 443 from on-prem CIDR
  ☐ DNS forwarding — on-prem resolves S3 to VPCE IP
  ☐ IAM credentials or Roles Anywhere configured
  ☐ Firewall rules — on-prem outbound 443 to VPCE IP

SQS
  ☐ SQS Interface VPCE created in VPC
  ☐ Security group on VPCE allows 443 from on-prem CIDR
  ☐ DNS forwarding — on-prem resolves SQS to VPCE IP
  ☐ Same IAM credentials cover SQS (same policy)
  ☐ Firewall rules — on-prem outbound 443 to VPCE IP

Databricks
  ☐ Network path from on-prem to Databricks confirmed
  ☐ PrivateLink or proxy in place if needed
  ☐ Service principal / PAT credentials
  ☐ Firewall rules — outbound 443 to Databricks endpoints
  ☐ TLS CA certs trusted on on-prem hosts

Credentials
  ☐ IAM Roles Anywhere or static creds strategy decided
  ☐ Secret storage and rotation plan
  ☐ Databricks credentials stored securely