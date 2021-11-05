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

> ansible.cfg defines the default behavior of ansible, playbooks/cli can override

# Azure - bad
Creating a virtual machine in Azure requires several different Azure resources; a resource group, virtual network, subnet, public ip address, network security group, network interface card, and the virtual machine itself. Each of these Azure resources can be managed and modified using an Ansible module. These Ansible modules allow you to codify your infrastructure in yaml files in the form of Ansible playbooks. 

Creating a resource group via Azure CLI does not require specifying a subscription,   
even though RGs are subcomponents of a subscription. The default subscription is used,  
but can be overruled by `--subscription` flag. 