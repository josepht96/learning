# Merge me


# creating inventory
Use ssh keys to access remote servers
ssh-keygen -t rsa -f ~/.ssh/ansible
Put public key on vms, keep private key on the control plane
~/.ssh/authorized_keys - pass in private key when you access the server ssh -i id_ras user@server
can use ssh-copy-id -i id_ras will copy public key to remote server
Ansible assumes default root user, ansible will autoamtically pickup public keys if theyre in the default location


# cmds
ansible -m ping all
ansible --version
ansible -m setup <server name>
ansibe -m ping localhost
ansible -m ping -i ~/inventoryfile all/<server names>
ansible -m command -a data -i ~/inventoryfile <server name>