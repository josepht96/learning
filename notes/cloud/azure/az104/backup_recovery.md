Back up Azure IaaS VMs. Azure Backup provides independent and isolated backups to guard against accidental destruction of original data. Backups are stored in a Recovery Services vault with built-in management of recovery points. Configuration and scalability is simple, backups are optimized, and you can easily restore as needed.
Get unlimited data transfer. Azure Backup does not limit the amount of inbound or outbound data you transfer, or charge for the data that is transferred.
**Recovery Services vaults store backup data for various Azure services such as IaaS VMs (Linux or Windows) and Azure SQL databases. Recovery Services vaults support System Center DPM, Windows Server, Azure Backup Server, and other services. Recovery Services vaults make it easy to organize your backup data, while minimizing management overhead. The Recovery Services vault can be used to back up Azure file shares.**

For on prem:
*Create the recovery services vault*. Within your Azure subscription, you will need to create a recovery services vault for the backups.
*Download the agent and credential file.* The recovery services vault provides a link to download the Azure Backup Agent. The Backup Agent will be installed on the local machine. There is also a credentials file that is required during the installation of the agent. You must have the latest version of the agent. Versions of the agent below 2.0.9083.0 must be upgraded by uninstalling and reinstalling the agent.
*Install and register agent*. The installer provides a wizard to configure the installation location, proxy server, and passphrase information. The downloaded credential file will be used to register the agent.
*Configure the backup*. Use the agent to create a backup policy including when to backup, what to backup, how long to retain items, and settings like network throttling.


Implement backup and recovery

Create a Recovery Services vault.
Create and configure backup policy.
Perform backup and restore operations by using Azure Backup.
Perform site-to-site recovery by using Azure Site Recovery.
Configure and review backup reports.

*For backing up Azure VMs running production workloads, use Azure Backup.* Azure Backup supports application-consistent backups for both Windows and Linux VMs. Azure Backup creates recovery points that are stored in geo-redundant recovery vaults. When you restore from a recovery point, you can restore the whole VM or just specific files.
Backup the virtual machine. The Azure VM Agent must be installed on the Azure virtual machine for the Backup extension to work. However, if your VM was created from the Azure gallery, then the VM Agent is already present on the virtual machine. VMs that are migrated from on-premises data centers would not have the VM Agent installed. In such a case, the VM Agent needs to be installed. Doesnt work for linux

Azure Site Recovery protects your VMs from a major disaster scenario when a whole region experiences an outage due to major natural disaster or widespread service interruption. You can configure Azure Site Recovery for your VMs so that you can recover your application with a single click in matter of minutes. You can replicate to an Azure region of your choice.


# Snapshot
In development and test environments, snapshots provide a quick and simple option for backing up VMs that use Managed Disks. A managed disk snapshot is a read-only full copy of a managed disk that is stored as a standard managed disk by default. With snapshots, you can back up your managed disks at any point in time. 
An Azure backup job consists of two phases. First, a virtual machine snapshot is taken. Second, the virtual machine snapshot is transferred to the Azure Recovery Services vault.

# Azure monitor
enables you to gather monitoring and diagnostic information about the health of your services. You can use this information to visualize and analyze the causes of problems that might occur in your app.
Monitor and visualize metrics. Metrics are numerical values available from Azure resources helping you understand the health, operation and performance of your system.
Query and analyze logs. Logs are activity logs, diagnostic logs, and telemetry from monitoring solutions; analytics queries help with troubleshooting and visualizations.
Setup alerts and actions. Alerts notify you of critical conditions and potentially take automated corrective actions based on triggers from metrics or logs