Network rules
Azure firewall blocks traffic by default

# NAT Rules
For external connections to azure resources, like ssh and rdp. Non-http traffic
# Network Rules
Similar to above, but allows traffic between resources within the virtual network. SQL database to web server or something.
# Application Rules
Handles ingress traffic at layer 7. How public internet can access resources within firewall

# virtual networking
Deny inbound traffic from internet
allow outbound traffic to internet

Standard public IPs are zone redundant, basic are not

Application rule for Windows update

what are virtual appliances
azure firewall vs azure web application firewall

Application gateway provides a WAF for inbound connections only for HTTP/S traffic (OWASP rules and more), Azure Firewall provides both inbound and outbound filtering also for non-HTTP traffic (E.G. your VMs can only go out to FQDN X, Y on port Z, K. and block other traffic)

# backup and recovery
review the differences between azure backup, site recovery, managed snapchots, recovery services vault
snapshots for database disks
determine which ones support linux and non azure vms
notes

azure recovery service vault is basically a storage resource that stores backup data of vms, sql server, or file shares
azure backup agent is for files and folders for on prem or azure vms

# Azure Backup
For backing up Azure VMs running production workloads, use Azure Backup. Azure Backup supports application-consistent backups for both Windows and Linux VMs. Azure Backup creates recovery points that are stored in geo-redundant recovery vaults. When you restore from a recovery point, you can restore the whole VM or just specific files.

# Back up on-premises machines:
You can back up on-premises Windows machines directly to Azure by using the Azure Backup Microsoft Azure Recovery Services (MARS) agent. Linux machines aren't supported.
You can back up on-premises machines to a backup server - either System Center Data Protection Manager (DPM) or Microsoft Azure Backup Server (MABS). You can then back up the backup server to a Recovery Services vault in Azure.

# Back up Azure VMs:
You can back up Azure VMs directly. Azure Backup installs a backup extension to the Azure VM agent that's running on the VM. This extension backs up the entire VM.
You can back up specific files and folders on the Azure VM by running the MARS agent.
You can back up Azure VMs to the MABS that's running in Azure, and you can then back up the MABS to a Recovery Services vault.