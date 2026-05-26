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

Direct Connect established and routed to GovCloud VPC
On-prem subnet routes to VPC CIDR via Direct Connect

S3

S3 Interface VPCE created in VPC
Security group on VPCE allows 443 from on-prem CIDR
DNS forwarding — on-prem resolves S3 to VPCE IP
Firewall rules — on-prem outbound 443 to VPCE IP

SQS

SQS Interface VPCE created in VPC
Security group on VPCE allows 443 from on-prem CIDR
DNS forwarding — on-prem resolves SQS to VPCE IP
Firewall rules — on-prem outbound 443 to VPCE IP

Databricks

Network path from on-prem to Databricks confirmed
PrivateLink or proxy in place if needed
Service principal / PAT credentials
Firewall rules — outbound 443 to Databricks endpoints
TLS CA certs trusted on on-prem hosts

AWS Credentials

IAM user created per service with least-privilege policy scoped to required S3 buckets and SQS queues
Access key and secret generated and stored in secrets manager
Credentials delivered to each on-prem service via environment variables or AWS config file
60-day rotation schedule configured and automated
CloudTrail enabled to audit all API calls per credential
Old keys disabled immediately upon rotation


curl -X POST "http://127.0.0.1:4566/" \
  -d "Action=DeleteQueue&QueueUrl=http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/my-queue"

curl -X POST "http://127.0.0.1:4566/" \
  -d "Action=CreateQueue&QueueName=my-queue"

# Verify new URL format
curl -X POST "http://127.0.0.1:4566/" \
  -d "Action=ListQueues"

  grep -i "OutOfMemoryError\|heap space" /var/log/app/*.logfree