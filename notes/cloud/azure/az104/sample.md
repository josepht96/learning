AB
D
A
A
AB
E - Geo-redundant storage (GRS) copies your data synchronously three times within a
single physical location in the primary region using LRS. It then copies your data
asynchronously to a single physical location in a secondary region that is hundreds
of miles away from the primary region. Storage accounts configured with the
Premium performance setting only support LRS. Any storage account already
configured with ZRS cannot be changed or directly switched to another replication
setting. In this scenario, the only storage account that is not set to LRS or Premium
performance is storage2, which can be switched to use Geo-redundant storage.

B
A
B, maybe D  - Availability and scaling is used to set up and manage VM high availability, not for
creating additional VMs based upon a set configuration setting.
B
C
AC Recovery Services vault supports Azure Virtual Machines, SQL in Azure VM, Azure
Files, SAP HANA in Azure VM, Azure Backup Server, Azure Backup Agent, and DPM.
Backup vault supports Azure Database for PostgreSQL servers, Azure Blobs, and
Azure disks



D
B
B
A
B
D
E
D

Lifecycle management policies apply rules to supported storage accounts to control
the transition of data to cooler storage tiers. Lifecycle management policies are
supported for block blobs and append blobs in general-purpose v2, premium block
blob, and Blob Storage accounts. FileStorage and general purpose v1 storage
accounts do not support lifecycle management.

Azure Bastion is a service that provides secure and seamless RDP/SSH connectivity
to your virtual machines directly from the Azure portal over TLS. When you connect
via Azure Bastion, your virtual machines do not need a public IP address, agent, or
special client software. Not requiring a public IP address protects the virtual
machine from outside port scanning.

The storage replication type cannot be changed after protecting items. Since Rsv3 is
the only Recovery services vault that does not contain items, it is the only one that
can be modified to use the Locally-redundant storage replication type.