# Ansible notes

Creating a virtual machine in Azure requires several different Azure resources; a resource group, virtual network, subnet, public ip address, network security group, network interface card, and the virtual machine itself. Each of these Azure resources can be managed and modified using an Ansible module. These Ansible modules allow you to codify your infrastructure in yaml files in the form of Ansible playbooks. 

Creating a resource group via Azure CLI does not require specifying a subscription,   
even though RGs are subcomponents of a subscription. The default subscription is used,  
but can be overruled by `--subscription` flag. 