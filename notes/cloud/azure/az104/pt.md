Scales sets and availability sets
Backup vms
managed disks
types of network records
gateways

Availaiblity sets are data center bound, just dont share a power source based on fault domain. 
Same with scale sets

# Scale sets
An availability set consists of a set of discrete VMs which have their own names and individual properties, but are spread across fault domains, which means when you have more than one VM in a set it reduces the chances of losing all your VMs in event of a hardware failure in the host or rack.
A scale set consists of a set of identically configured VMs, also spread across fault domains (in fact a scale set is an implicit availability set with 5 fault domains). The main difference is scale sets, being identical, make it very easy to add or remove VMs from the set while preserving high availability, which in turn makes it easy to implement autoscale, and to perform operations on the whole set or a subset of VMs. There are also API calls that support re-imaging and upgrading VMs, allowing you to roll out an update while keeping the service running. They are useful for cloud architectures which require deploying large numbers of similar VMs, or need to be elastic. A typical architecture might use a scale set for agent or worker nodes, and an availability set for master or control nodes. See https://azure.microsoft.com/en-us/services/virtual-machine-scale-sets/ for more detail.

# Availability sets
VMs are also aligned with disk fault domains. This alignment ensures that all the managed disks attached to a VM are within the same fault domains.

Only VMs with managed disks can be created in a managed availability set. The number of managed disk fault domains varies by region - either two or three managed disk fault domains per region.

what is azure cost center
--
Select all true statements that apply to the use of Azure Disk Encryption (ADE) for Azure VM disk protection.
For more information about the Azure Disk Encryption, go to:

https://docs.microsoft.com/en-us/azure/security/fundamentals/azure-disk-encryption-vms-vmss

https://docs.microsoft.com/en-us/learn/modules/secure-your-azure-virtual-machine-disks/2-encryption-options-for-protecting-windows-and-linux-vms
--
You have defined an autoscale condition with four autoscale rules. The first rule scales out when the CPU utilization reaches 70 percent. The second rule scales back in when the CPU utilization drops below 50 percent. The third rule scales out if memory occupancy exceeds 75 percent. The fourth rule scales back in when memory occupancy falls below 50 percent. When will the system scale out?

For more information on the Import-Export service, go to:

https://docs.microsoft.com/en-us/azure/storage/common/storage-import-export-service
--
You need to allow traffic onto certain FQDN’s via the Azure Firewall. Which of the following rules would you create for this requirement?
--
Which of the following network watcher feature would you use for the following requirement?
--
Find out if there is outbound connectivity between an Azure virtual machine and an external host.
--
You have set up a computer named getcloudskillsclient1 that has a point-to-site VPN connection to an Azure virtual network named getcloudskillsnetwork. The point-to-site connection makes use of a self-signed certificate. You now have to establish a point-to-site VPN connection to the same virtual network from another computer named getcloudskillsclient2. The VPN client configuration package is downloaded and installed on the getcloudskillsclient2 computer.

You decide to join the getcloudskillsclient2 computer to Azure AD.

Would this fulfill the requirement?
--
Azure Kubernetes Service provides two ways of scaling your application: scaling nodes and scaling pods. Both types of scaling can be done manually or automatically. You can scale clusters by scaling the nodes. Autoscaler for the clusters monitors the pods scaling. When there are not enough resources for the pods, the cluster autoscaler adds new nodes.

For more information about AKS pods scaling, go to:

https://docs.microsoft.com/en-us/azure/aks/tutorial-kubernetes-scale#manually-scale-pods

https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#scale

https://docs.microsoft.com/en-us/cli/azure/aks?view=azure-cli-latest#az_aks_scale

https://docs.microsoft.com/en-us/azure/aks/concepts-scale

--
For more information on availability zones, go to:

https://docs.microsoft.com/en-us/azure/availability-zones/az-overview

--
Explanation
Availability sets can’t protect virtual machines from a data center-level failure. Availability zones protect VMs from data center failure.

You need to distribute your virtual machines across three Availability Zones.

For more information on Availability sets, go to:

- https://docs.microsoft.com/en-us/azure/virtual-machines/windows/manage-availability

- https://social.technet.microsoft.com/wiki/contents/articles/51828.azure-vms-availability-sets-and-availability-zones.aspx

--
Explanation
Neither Scale Sets nor Managed Disks can't help to meet the goal of 99.95%.

For more information about Scale Set and Managed Disks SLA, visit:

- https://azure.microsoft.com/en-ca/support/legal/sla/virtual-machine-scale-sets/v1_1/

- https://azure.microsoft.com/en-in/support/legal/sla/managed-disks/v1_0/

--
Users are reporting that when they attempt to access myapps.microsoft.com, they are prompted multiple times to sign in and are forced to use an account name that ends with onmicrosoft.com. You discover that there is a UPN mismatch between Azure AD and the on-premises Active Directory. You need to ensure that the users can use single-sign-on (SSO) to access Azure resources. What should you do first?


E - Geo-redundant storage (GRS) copies your data synchronously three times within a
single physical location in the primary region using LRS. It then copies your data
asynchronously to a single physical location in a secondary region that is hundreds
of miles away from the primary region. Storage accounts configured with the
Premium performance setting only support LRS. Any storage account already
configured with ZRS cannot be changed or directly switched to another replication
setting. In this scenario, the only storage account that is not set to LRS or Premium
performance is storage2, which can be switched to use Geo-redundant storage.

B, maybe D  - Availability and scaling is used to set up and manage VM high availability, not for
creating additional VMs based upon a set configuration setting