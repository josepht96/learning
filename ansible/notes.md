# Ansible notes

# run options
--check - dry run
--start-at-task - task to start at
--tags "install" - run only specific tasks with the matching tags
--skip-tags "install" - skips specified tag matching tasks
ansible -m setup localhost | grep distribution_version
ansible -i inventory -m setup web1 | grep distribution
ansible -m setup localhost

```yaml
gather_facts: no # disables fact gathering
```
# Ansible.cfg
ansible.cfg defines the default behavior of ansible, playbooks/cli can override
Can also override it entirely with a new config file (copy the default to a new folder somewhere, or rename it). Set env var to path ($ANSIBLE_CONFIG)
Can override specific properties with export $ANSIBLE_SOMEATTRIBUTE=somevalue
ansible-config view

# creating inventory
Use ssh keys to access remote servers
ssh-keygen -t rsa -f ~/.ssh/ansible
Put public key on vms, keep private key on the control plane
~/.ssh/authorized_keys - pass in private key when you access the server ssh -i id_ras user@server
can use ssh-copy-id -i id_ras will copy public key to remote server
Ansible assumes default root user, ansible will autoamtically pickup public keys if theyre in the default location


>generate keys: ssh-keygen -f /file/path
>copy public key to remote server: ssh-copy-id -i /path/output/from/above user@server

# cmds
ansible -m ping all
ansible --version
ansible -m setup <server name>
ansibe -m ping localhost
ansible -m ping -i ~/inventoryfile all/<server names>
ansible -m command -a data -i ~/inventoryfile <server name>

# Creating shell scripts
```bash
export ANSIBLE_GATHERING=explicit
echo hello there
ansible -m ping all
ansible -a 'cat /etc/hosts' all
ansible-playbook playbook.yml
chmod +x <filename.sh>
-vv verbose mode
--ask-become-pass - #prompts for user input password
sudo chown <user:grpoup> /etc/file/path
```

# elevating users
```yml
- name: install nginx
  become: yes
  become_method: doas
  become_user: nginx
  hosts: all
  - yum:
    name: nginx
    state: latest

```
# Azure - bad
Creating a virtual machine in Azure requires several different Azure resources; a resource group, virtual network, subnet, public ip address, network security group, network interface card, and the virtual machine itself. Each of these Azure resources can be managed and modified using an Ansible module. These Ansible modules allow you to codify your infrastructure in yaml files in the form of Ansible playbooks. 

Creating a resource group via Azure CLI does not require specifying a subscription,   
even though RGs are subcomponents of a subscription. The default subscription is used,  
but can be overruled by `--subscription` flag. 