sudo docker build -t sender -f ./sender/Dockerfile .


use local docker repo

eval $(minikube docker-env)    
docker build -t sender -f ./sender/Dockerfile .
kubectl run sender --image sender:latest --image-pull-policy=Never


expose node port
kubectl expose pod sender --port=80 --type=NodePort
then access the pod via the node (minikube ip) and the 30000 port

docker exec -it ubuntu_bash bash
docker run -d --name somecontainername --mount somevolumename,target=/sometargetinsidethecontainer nginx:latest