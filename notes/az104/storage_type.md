From <https://docs.microsoft.com/en-us/learn/modules/configure-storage-accounts/3-explore-azure-storage-services> 
From <https://docs.microsoft.com/en-us/learn/modules/configure-storage-accounts/2-implement-azure-storage> 
From <https://docs.microsoft.com/en-us/learn/modules/configure-storage-security/2-review-strategies> 

# Storage data objects
Azure Storage includes these data services, each of which is accessed through a storage account.
    1. Azure Containers (Blobs): A massively scalable object store for text and binary data.
    2. Azure Files: Managed file shares for cloud or on-premises deployments.
    3. Azure Queues: A messaging store for reliable messaging between application components.
    4. Azure Tables: A NoSQL store for schemaless storage of structured data.
    5. Container (blob) storage
Azure Blob storage is Microsoft's object storage solution for the cloud. Blob storage is optimized for storing massive amounts of unstructured data, such as text or binary data. Blob storage is ideal for:
•	Serving images or documents directly to a browser.
•	Storing files for distributed access.
•	Streaming video and audio.
•	Storing data for backup and restore, disaster recovery, and archiving.
•	Storing data for analysis by an on-premises or Azure-hosted service.
Objects in Blob storage can be accessed from anywhere in the world via HTTP or HTTPS. Users or client applications can access blobs via URLs, the Azure Storage REST API, Azure PowerShell, Azure CLI, or an Azure Storage client library. The storage client libraries are available for multiple languages, including .NET, Java, Node.js, Python, PHP, and Ruby.
Azure files
Azure Files enables you to set up highly available network file shares that can be accessed by using the standard Server Message Block (SMB) protocol. That means that multiple VMs can share the same files with both read and write access. You can also read the files using the REST interface or the storage client libraries.
One thing that distinguishes Azure Files from files on a corporate file share is that you can access the files from anywhere in the world using a URL that points to the file and includes a shared access signature (SAS) token. You can generate SAS tokens; they allow specific access to a private asset for a specific amount of time.
File shares can be used for many common scenarios:
•	Many on-premises applications use file shares. This feature makes it easier to migrate those applications that share data to Azure. If you mount the file share to the same drive letter that the on-premises application uses, the part of your application that accesses the file share should work with minimal, if any, changes.
•	Configuration files can be stored on a file share and accessed from multiple VMs. Tools and utilities used by multiple developers in a group can be stored on a file share, ensuring that everybody can find them, and that they use the same version.
•	Diagnostic logs, metrics, and crash dumps are just three examples of data that can be written to a file share and processed or analyzed later.
The storage account credentials are used to provide authentication for access to the file share. This means anybody with the share mounted will have full read/write access to the share.
Queue storage
The Azure Queue service is used to store and retrieve messages. Queue messages can be up to 64 KB in size, and a queue can contain millions of messages. Queues are used to store lists of messages to be processed asynchronously.
For example, if you want your customers to be able to upload pictures, and you want to create thumbnails for each picture. You could have your customer wait for you to create the thumbnails while uploading the pictures. An alternative would be to use a queue. When the customer finishes the upload, write a message to the queue. Then have an Azure Function retrieve the message from the queue and create the thumbnails. Each of the processing parts can be scaled separately, giving you more control when tuning it for your usage.
Table storage
Azure Table storage is now part of Azure Cosmos DB. In addition to the existing Azure Table storage service, there is a new Azure Cosmos DB Table API offering that provides throughput-optimized tables, global distribution, and automatic secondary indexes. Table storage is ideal for storing structured, non-relational data.
> Durable and highly available. Redundancy ensures that your data is safe during transient hardware failures. You replicate data across datacenters or geographical regions for protection from local catastrophe or natural disaster. Data replicated remains highly available during an unexpected outage.
> Secure. All data written to Azure Storage is encrypted by the service. Azure Storage provides you with fine-grained control over who has access to your data.
> Scalable. Azure Storage is designed to be massively scalable to meet the data storage and performance needs of today's applications.
> Managed. Microsoft Azure handles hardware maintenance, updates, and critical issues for you.
> Accessible. Data in Azure Storage is accessible from anywhere in the world over HTTP or HTTPS. Microsoft provides SDKs for Azure Storage in a various languages .NET, Java, Node.js, Python, PHP, Ruby, Go, and REST API. Azure Storage supports scripting in Azure PowerShell or Azure CLI. And the Azure portal and Azure Storage Explorer offer easy visual solutions for working with your data.
> Storage for Virtual Machines. Virtual machine storage includes disks and files. Disks are persistent block storage for Azure IaaS virtual machines. Files are fully managed file shares in the cloud.
> Unstructured Data. Unstructured data includes Blobs and Data Lake Store. Blobs are highly scalable, REST-based cloud object store. Data Lake Store is Hadoop Distributed File System (HDFS) as a service.
> Structured Data. Structured data includes Tables, Cosmos DB, and Azure SQL DB. Tables are a key/value, autoscaling NoSQL store. Cosmos DB is a globally distributed database service. Azure SQL DB is a fully managed database-as-a-service built on SQL.
# General purpose storage accounts have two tiers: Standard and Premium.
Standard storage accounts are backed by magnetic drives (HDD) and provide the lowest cost per GB. Use Standard storage for applications that require bulk storage or where data is infrequently accessed.
Premium storage accounts are backed by solid-state drives (SSD) and offer consistent low-latency performance. Use Premium storage for Azure virtual machine disks with I/O-intensive applications, like databases.
# Locally-Redundant Storage (LRS)
LRS replicates data three times within one data center located in a primary region. When LRS is enabled, Azure Storage only registers write requests as successful once data is written to three replicas. LRS provides at least 99.999999999% durability for objects during a given year.
ZRS performs replication across three Azure availability zones. Each Azure availability zone is an individual physical location with its own independent networking, power, and cooling. ZRS provides a minimum of 99.9999999999% durability for objects during a given year.
Geo-Redundant Storage (GRS)
GRS provides additional redundancy for data storage compared to LRS or ZRS. In addition to the three copies of data stored in one region, there are three copies stored in a paired Azure region. So GRS provides all the features of LRS storage in the primary zone, and additionally, provides a secondary LRS data storage in another region.
There are two disadvantages of GRS redundancy:
•	Replication between regions is asynchronous and so data is propagated with a small delay
•	The second region cannot be accessed or read until the storage account fails over
# Zone-redundant storage (ZRS)
ZRS replicates data in three separate zones within a single geo-region. Will survive a single datacenter outage. Good for resilient region based data access.
# Geo-redundant storage (GRS)
GRS replicates data 3 times in single zone in primary geo-region, and then 3 times in a zone in a secondary geo-region. Secondary region is read only should the primary fail. Does not protect from a single datacenter outage in 1 region. 
# Geo-zone-redundant storage
Geo-zone-redundant storage (GZRS) combines the high availability provided by redundancy across availability zones with protection from regional outages provided by geo-replication. Data in a GZRS storage account is copied across three Azure availability zones in the primary region and is also replicated to a secondary geographic region for protection from regional disasters. Microsoft recommends using GZRS for applications requiring maximum consistency, durability, and availability, excellent performance, and resilience for disaster recovery.
With a GZRS storage account, you can continue to read and write data if an availability zone becomes unavailable or is unrecoverable. Additionally, your data is also durable in the case of a complete regional outage or a disaster in which the primary region isn't recoverable. GZRS is designed to provide at least 99.99999999999999% (16 9's) durability of objects over a given year.
Data is spread across 3 zones in primary region, and 1 zone in secondary region. Available read only if primary goes offline. Protected from datacenter outage but not from region failure. 
# Read-Access Geo-Redundant (RA-GRS)
RA-GRS has all the same level of redundancy of standard GRS replication, with an additional benefit—the secondary copies stored in paired Azure regions are readable. This means that if your application is configured correctly, you can use multiple readable endpoints. This increases the SLA for read operations to 99.99%.
However, the SLA for write operations remains 99.9%, because a single area still controls write and update operations.
Due to their asynchronous replication, both types of GRS replication have some replication delay. You can use the LastSyncTime parameter to ensure you are reading the latest copy of the data.
Object Replication for Block Blob Storage
The preceding replication methods were relevant for all Azure storage services. This is a special replication method which is available only for Block Blob Storage. 
The object replication method is asynchronous. You can use it to automatically move data to an archive tier, to optimize data distribution and reduce costs. Or, you can use it to synchronize data to a storage resource running nearer to your users, to reduce latency.
Block blobs are replicated according to your replication policy, which specifies source/target Azure accounts and containers, and which block blobs should be replicated.
Block blob object replication copies:
•	Blob content
•	Blob version
•	Blob metadata
# Read-Access Geo-Zone-Redundant (RA-GZRS)
Similar setup as above, except with additional resilience of having data replicate in distinct zones in the primary region. 
# Security
Azure Storage provides a comprehensive set of security capabilities that together enable developers to build secure applications.
•	Encryption. All data written to Azure Storage is automatically encrypted using Storage Service Encryption (SSE).
•	Authentication. Azure Active Directory (Azure AD) and Role-Based Access Control (RBAC) are supported for Azure Storage for both resource management operations and data operations, as follows:
•	You can assign RBAC roles scoped to the storage account to security principals and use Azure AD to authorize resource management operations such as key management.
•	Azure AD integration is supported for data operations on the Blob and Queue services.
•	Data in transit. Data can be secured in transit between an application and Azure by using Client-Side Encryption, HTTPS, or SMB 3.0.
•	Disk encryption. OS and data disks used by Azure virtual machines can be encrypted using Azure Disk Encryption.
•	Shared Access Signatures. Delegated access to the data objects in Azure Storage can be granted using Shared Access Signatures.
# storage account key
Gives access to the entire account, not just read access. Two default keys, protect from view, regenerate keys, consider storing in key vault. Recommended to use AAD instead. 
# SAS (shared access signature)
A shared access signature (SAS) is a URI that grants restricted access rights to Azure Storage resources. You can provide a SAS to clients who shouldn't have access to your storage account key. By distributing a SAS URI to these clients, you grant them access to a resource for a specified period of time. SAS is a secure way to share your storage resources without compromising your account keys.
Shared access signatures. Shared access signatures (SAS) delegate access to a particular resource in your account with specified permissions and over a specified time interval.
A SAS gives you granular control over the type of access you grant to clients who have the SAS, including:
•	An account-level SAS can delegate access to multiple storage services. For example, blob, file, queue, and table.
•	An interval over which the SAS is valid, including the start time and the expiry time.
•	The permissions granted by the SAS. For example, a SAS for a blob might grant read and write permissions to that blob, but not delete permissions.
Optionally, you can also:
•	Specify an IP address or range of IP addresses from which Azure Storage will accept the SAS. For example, you might specify a range of IP addresses belonging to your organization.
•	The protocol over which Azure Storage will accept the SAS. You can use this optional parameter to restrict access to clients using HTTPS.
# AAD
Azure Active Directory (Azure AD). Azure AD is Microsoft's cloud-based identity and access management service. With Azure AD, you can assign fine-grained access to users, groups, or applications via role-based access control (RBAC).
# Anonymous access
Anonymous access to containers and blobs. You can optionally make blob resources public at the container or blob level. A public container or blob is accessible to any user for anonymous read access. Read requests to public containers and blobs do not require authorization.
# Roles
Owner: can do everything
Contributor: add things, work with things
Reader: see whats inside
Delegator: can delegate a storage account for something?
# virtual network
layered security
Limit access by ip addresses, ip ranges, subnets
# Azure files
Cloud based SMB or NFS file share. Accessible from W, Linux, Mac OS
Clients use 445. Need to know requirements: performance, sizing, redundancy. GPV2 and FileStorage are best options generally. Main difference is performance and FileStorage not having Geo-redundancy
Both go up to 100tb
GPV2: lower performance, more resilience with LRS, GRS, ZRS, and GZRS options
FileStorage (file share): higher performance, less resilient with LRS and ZRS options
Can add soft delete features to recover deleted files. 
# Tiers
Tier of a storage account can be changed after creation. 
Hot tier: Optimized for frequent file sharing and access. Cost of storage is higher but cost of access is lower.
Cold tier: Optimal for infrequent access, cost of storage is low, cost of access is high
Transcation optimized: High storage cost, lowest cost per access. Good for backend storage for applications (consistent throughput)

# Azure file sync
Allows users to sync azure file share with on-prem servers. Replication occurs between onprem datacenters and azure. Administrators can use with existing fileshare infrastructure.
>Cloud endpoint: Azure file share thats being synced 
>Server endpoint: windows server and its local file system path that syncs with azure file share
>Sync group: defines relationship between cloud endpoint and server endpoint. Can only have 1 cloud endpoint per sync group, but multiple server endpoints.
Steps:
    1. Deploy the Storage Sync Service
    2. Create a sync group and cloud endpoint
    3. Install the Azure File Sync agent on Windows Servers
    4. Register Windows Serv with the Storage Sync Service
    5. Create server endpoint and wait for sync

Azure File Sync agent. The Azure File Sync agent is a downloadable package that enables Windows Server to be synced with an Azure file share. The Azure File Sync agent has three main components:
FileSyncSvc.exe: The background Windows service that is responsible for monitoring changes on server endpoints, and for initiating sync sessions to Azure.
StorageSync.sys: The Azure File Sync file system filter, which is responsible for tiering files to Azure Files (when cloud tiering is enabled).
PowerShell management cmdlets: PowerShell cmdlets that you use to interact with the Microsoft.StorageSync Azure resource provider. You can find these at the following (default) locations:
C:\Program Files\Azure\StorageSyncAgent\StorageSync.Management.PowerShell.Cmdlets.dll
C:\Program Files\Azure\StorageSyncAgent\StorageSync.Management.ServerCmdlets.dll
Server endpoint. A server endpoint represents a specific location on a registered server, such as a folder on a server volume. Multiple server endpoints can exist on the same volume if their namespaces do not overlap (for example, F:\sync1 and F:\sync2). You can configure cloud tiering policies individually for each server endpoint. You can create a server endpoint via a mountpoint. Note, mountpoints within the server endpoint are skipped. You can create a server endpoint on the system volume but, there are two limitations if you do so:
Cloud tiering cannot be enabled.
Rapid namespace restore (where the system quickly brings down the entire namespace and then starts to recall content) is not performed.
Cloud endpoint. A cloud endpoint is an Azure file share that is part of a sync group. The entire Azure file share syncs, and an Azure file share can be a member of only one cloud endpoint. Therefore, an Azure file share can be a member of only one sync group. If you add an Azure file share that has an existing set of files as a cloud endpoint to a sync group, the existing files are merged with any other files that are already on other endpoints in the sync group.

# Azure blob storage
Objects in Blob storage can be accessed from anywhere in the world via HTTP or HTTPS. Users or client applications can access blobs via URLs, the Azure Storage REST API, Azure PowerShell, Azure CLI, or an Azure Storage client library. The storage client libraries are available for multiple languages, including .NET, Java, Node.js, Python, PHP, and Ruby.
Azure files
Azure Files enables you to set up highly available network file shares that can be accessed by using the standard Server Message Block (SMB) protocol. That means that multiple VMs can share the same files with both read and write access. You can also read the files using the REST interface or the storage client libraries.
One thing that distinguishes Azure Files from files on a corporate file share is that you can access the files from anywhere in the world using a URL that points to the file and includes a shared access signature (SAS) token. You can generate SAS tokens; they allow specific access to a private asset for a specific amount of time.
File shares can be used for many common scenarios:
•	Many on-premises applications use file shares. This feature makes it easier to migrate those applications that share data to Azure. If you mount the file share to the same drive letter that the on-premises application uses, the part of your application that accesses the file share should work with minimal, if any, changes.
•	Configuration files can be stored on a file share and accessed from multiple VMs. Tools and utilities used by multiple developers in a group can be stored on a file share, ensuring that everybody can find them, and that they use the same version.
•	Diagnostic logs, metrics, and crash dumps are just three examples of data that can be written to a file share and processed or analyzed later.
The storage account credentials are used to provide authentication for access to the file share. This means anybody with the share mounted will have full read/write access to the share.
Queue storage
The Azure Queue service is used to store and retrieve messages. Queue messages can be up to 64 KB in size, and a queue can contain millions of messages. Queues are used to store lists of messages to be processed asynchronously.
For example, if you want your customers to be able to upload pictures, and you want to create thumbnails for each picture. You could have your customer wait for you to create the thumbnails while uploading the pictures. An alternative would be to use a queue. When the customer finishes the upload, write a message to the queue. Then have an Azure Function retrieve the message from the queue and create the thumbnails. Each of the processing parts can be scaled separately, giving you more control when tuning it for your usage.
Table storage
Azure Table storage is now part of Azure Cosmos DB. In addition to the existing Azure Table storage service, there is a new Azure Cosmos DB Table API offering that provides throughput-optimized tables, global distribution, and automatic secondary indexes. Table storage is ideal for storing structured, non-relational data.

# key things
Create and configure Azure file shares
Understand how file sync is deployed
implement and configure azure blobs
tiering options for Azure blobs, migrating
Lifecycle managment
why and how to implement blob replication
get some hands on experience

# Import & Export service
Moving large amounts of data in and out of Azure blob storage.
1. send data over the wire - good for small amounts of data. Large amounts of data can take a long time and be costly. 
2. Ship data to datacenter (tbs of data)
Azure import/export service, WAImpoirtExportTool, drives
Prep drives with WAImportExportTool and Bitlocker
Create import job in Azure portal
Ship drives to azure datacenter
# Managing data with Azure storage

# AzCopy
azcopy copy 'C:\data' 'https://someazureendpoint.net/2321' --recursive
Blob storage access: SAS or AAD
FileStorage: SAS only
azcopy list uri/sastoken
