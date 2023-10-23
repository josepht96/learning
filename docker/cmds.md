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


docker images -a
docker history --no-trunc 

docker run --rm -it \
      -v /var/run/docker.sock:/var/run/docker.sock \
      -v  "$(pwd)":"$(pwd)" \
      -w "$(pwd)" \
      -v "$HOME/.dive.yaml":"$HOME/.dive.yaml" \
      wagoodman/dive:latest build -t l1 .