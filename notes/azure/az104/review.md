# topics to cover ##############################################################

# vnet service endpoints
Virtual Network (VNet) service endpoint provides secure and direct connectivity to Azure services over an optimized route over the Azure backbone network. Endpoints allow you to secure your critical Azure service resources to only your virtual networks. Service Endpoints enables private IP addresses in the VNet to reach the endpoint of an Azure service without needing a public IP address on the VNet.

# dns record types
# network monitoring
# VPN gateway
# WAN
Virtual WAN: The virtualWAN resource represents a virtual overlay of your Azure network and is a collection of multiple resources. It contains links to all your virtual hubs that you would like to have within the virtual WAN. Virtual WAN resources are isolated from each other and can't contain a common hub. Virtual hubs across Virtual WAN don't communicate with each other.
handle multiple s2s vpn or p2s, vnet peering. Central connection for managing most networking

# Network rules

# Azure firewall
Azure Firewall is a cloud-native and intelligent network firewall security service that provides the best of breed threat protection for your cloud workloads running in Azure. It's a fully stateful, firewall as a service with built-in high availability and unrestricted cloud scalability. It provides both east-west and north-south traffic inspection.

Azure Firewall Standard provides L3-L7 filtering and threat intelligence feeds directly from Microsoft Cyber Security. Threat intelligence-based filtering can alert and deny traffic from/to known malicious IP addresses and domains which are updated in real time to protect against new and emerging attacks.

Azure Firewall Premium provides advanced capabilities include signature-based IDPS to allow rapid detection of attacks by looking for specific patterns. These patterns can include byte sequences in network traffic, or known malicious instruction sequences used by malware. There are more than 58,000 signatures in over 50 categories which are updated in real time to protect against new and emerging exploits. The exploit categories include malware, phishing, coin mining, and Trojan attacks.

# Azure application firewall
Web Application Firewall (WAF) provides centralized protection of your web applications from common exploits and vulnerabilities. Web applications are increasingly targeted by malicious attacks that exploit commonly known vulnerabilities. SQL injection and cross-site scripting are among the most common attacks.
Azure WAF protects inbound traffic to the web workloads, and the Azure Firewall inspects inbound traffic for the other applications. The Azure Firewall will cover outbound flows from both workload types.

# Disk encryption

# kusto query language
```sql
let StartTime=ago(12h);
let StopTime=now()
T
| where Timestamp > StartTime and Timestamp <= StopTime 
| where ...
| summarize Count=count() by bin(Timestamp, 5m)
```

# Azure Bastion
Azure Bastion is a service that provides secure and seamless RDP/SSH connectivity
to your virtual machines directly from the Azure portal over TLS. When you connect
via Azure Bastion, your virtual machines do not need a public IP address, agent, or
special client software. Not requiring a public IP address protects the virtual
machine from outside port scanning.

# Storage ###################################################################### 
Lifecycle management policies apply rules to supported storage accounts to control
the transition of data to cooler storage tiers. Lifecycle management policies are
supported for block blobs and append blobs in general-purpose v2, premium block
blob, and Blob Storage accounts. FileStorage and general purpose v1 storage
accounts do not support lifecycle management.


# Virtual Networking ###########################################################
# NAT Rules
For external connections to azure resources, like ssh and rdp. Non-http traffic
# Network Rules
Similar to above, but allows traffic between resources within the virtual network. SQL database to web server or something.
# Application Rules
Handles ingress traffic at layer 7. How public internet can access resources within firewall

Deny inbound traffic from internet
allow outbound traffic to internet

Standard public IPs are zone redundant, basic are not

Application rule for Windows update

what are virtual appliances
azure firewall vs azure web application firewall

Application gateway provides a WAF for inbound connections only for HTTP/S traffic (OWASP rules and more), Azure Firewall provides both inbound and outbound filtering also for non-HTTP traffic (E.G. your VMs can only go out to FQDN X, Y on port Z, K. and block other traffic)


# Backup #######################################################################
# backup and recovery
review the differences between azure backup, site recovery, managed snapchots, recovery services vault
snapshots for database disks
determine which ones support linux and non azure vms
notes

azure recovery service vault is basically a storage resource that stores backup data of vms, sql server, or file shares
azure backup agent is for files and folders for on prem or azure vms

AC Recovery Services vault supports Azure Virtual Machines, SQL in Azure VM, Azure
Files, SAP HANA in Azure VM, Azure Backup Server, Azure Backup Agent, and DPM.
Backup vault supports Azure Database for PostgreSQL servers, Azure Blobs, and
Azure disks

The storage replication type cannot be changed after protecting items. Since Rsv3 is
the only Recovery services vault that does not contain items, it is the only one that
can be modified to use the Locally-redundant storage replication type.

# Azure Backup
For backing up Azure VMs running production workloads, use Azure Backup. Azure Backup supports application-consistent backups for both Windows and Linux VMs. Azure Backup creates recovery points that are stored in geo-redundant recovery vaults. When you restore from a recovery point, you can restore the whole VM or just specific files.

# Back up on-premises machines:
You can back up on-premises Windows machines directly to Azure by using the Azure Backup Microsoft Azure Recovery Services (MARS) agent. Linux machines aren't supported.
You can back up on-premises machines to a backup server - either System Center Data Protection Manager (DPM) or Microsoft Azure Backup Server (MABS). You can then back up the backup server to a Recovery Services vault in Azure.

# Back up Azure VMs:
You can back up Azure VMs directly. Azure Backup installs a backup extension to the Azure VM agent that's running on the VM. This extension backs up the entire VM.
You can back up specific files and folders on the Azure VM by running the MARS agent.
You can back up Azure VMs to the MABS that's running in Azure, and you can then back up the MABS to a Recovery Services vault.