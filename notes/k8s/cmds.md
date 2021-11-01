# cmds:
minikube -
	start
	ip
	service <name> --url
	
kubectl -
	get pods
	get pods -n kube-system
	get pods --selector tagname=tagvalue,tag2name=tag2value,...
	get deployments
	get all
    run nginx --image nginx (name can be anything, image needs to be existing image
    hosted somewhere)
    describe pod nginx
    get pods -o wide
	port-forward service/<sn> hostport:container port
	get services <sn>
	delete service <sn>
	delete deployment <deployment name>
	create deployment nginx --image=nginx
	(apply and create work the same when creating new def, apply will throw warning though
	first time)
	create -f pod-definition.yml
	replace -f <file>
	scale --replicas=6 replicaset myapp-replicaset (wont update file)
	edit replicaset myapp-replicaset
	get all (list all objects)
    rollout status deployment.apps/myapp-deployment
	rollout history
	rollout undo
    taint nodes <node-name> key=value:taint-effect - running it again with "-" on end removes the taint
    taint nodes node01 spray=blue:NoSchedule