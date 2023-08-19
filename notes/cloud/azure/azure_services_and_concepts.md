1/27/2020
Microsoft Azure Services and Concepts
Understanding Azure Architecture and Management
Azure availability zones: Microsoft has datacenters all over the world. Azure services are created in locations which determines where the data for that service is stored
Resource groups: Contain collections of resources – VMs, Network security groups, virtual networks, storage space, etc.… Resources can be moved to a separate resource group or subscription. Tags can be added to specify ownership or anything else. Export template creates a json representation of the resource group, I imagine these can be saved and deployed in other subscriptions. A resource group is assigned to a datacenter location. Cost analysis feature allows you to track the cost of a resource group. 
Azure Resource Manager (ARM):  Creation, deployment, and modification of Azure resources. 
Resource manager templates are useful for sharing resource groups with other teams. They can also be saved so that if you remove the group, you can set it up easily at a later date (assuming you saved the json files). Downloading the template gives you template and parameters json files. Easy to redeploy the json files.
Azure Service Health: Overview of Azure servers in all regions. This may be useful if you see an outage in a certain region that you rely on. Azure Service Health is for Microsoft managed resources.
Azure Monitor is useful for monitoring specific resources that you manage. Provides a number of tools and approaches to monitoring various things (from networks down to microservices). 
Azure Mobile App: Resources can be managed from the mobile app, like they can be managed through the browser. Even execute Azure CLI or PowerShell commands. 
Azure Advisor: Recommendations on cost saving opportunities, security vulnerabilities, performance improvement, reliability, and ‘operational excellence’. These are personalized for your subscription, not generic recommendations. A big cost is VMs, so it might alert you if a VM is running (costing money) but not actively being used. 
Exploring Azure Core Products
Azure Compute: Logical grouping of services that are useful for running applications. Includes VMs, Containers, Azure App Service (PaaS), serverless computing. 
VMs: IaaS. Full control over the operating system but must manage it as well. Three factors to consider are: Type of image, size of VM, and availability. Can setup Ubuntu (!) or Redhat, even windows VMs…  Some images have already installed services, like SQL Server. Similar to using Container images with predefined applications installed. Can upload your own images with the configuration you want. Multiple types of VMs, I think some even have dedicated GPUs. Can setup VMs in different datacenter locations, helping protect against faults and outages.
Containers: VMs virtualize underlying hardware, contain the OS. Containers virtualize the OS, only container things necessary for the application to run. Can run on local workstation, on-prem servers, VMs in Azure, Azure Container Instances, Azure Kubernetes Service (AKS), or Azure App Services.
Containers can be configured to run on public internet, as well as on virtual networks. 
Azure App Service:  Manages underlying infrastructure needed for your applications to run. App Service plan defines required underlying infrastructure (pricing tier). Domains can be bought from Azure. As well, authentication options for authentication through 3rd party platforms (Facebook, google, Microsoft) or through Active Directory. 
Serverless Compute: Azure Functions allows you to run small functions to execute on trigger. Azure Logic Apps are like Functions but have access to more library connectors. Azure Event Grid helps you build apps with event based architecture. Functions can be triggered from other Azure services. When a container fails maybe you want to kick off a custom function to send an email. 
Core networking:
Virtual network is the basic communication pathway for Azure apps. Subnets can be created for VMs, resources on a subnet can communicate with resources in other subnets. By default, resources over one Virtual Network cannot communicate with resources on another, but this can be enabled via VNET Peering. Public IP addresses can be assigned to a VNET resource. Load balancers can be added to virtual networks to help with inbound traffic from outside sources. 4414274virtual networks to help with inbound traffic from outside sources. 
On-premise resources can be connected to Azure resources via a Hybrid network. Need to create VPN Gateway. Another type is ExpressRoute. Site-to-site VPN can be provided through Cisco VPN. Outbound and inbound security is control from NSG, network security group, basically firewalls. But there’s also Azure firewall, which provides more robust services.  
Windows Virtual Desktop: Windows Server 2019, 2016, 2012R2, W10 ent, W7 ent.
Content Delivery Network: Store cached data to reduce latency of traffic. Dynamic Site Acceleration is used for speeding up delivery of dynamic content. 
Data Storage in Azure
Azure database with multiple languages, structured data. For unstructured data, Azure Blob Storage and Azure File Storage, Azure Disk Storage. Semi-Structured data can be stored with Cosmos Db. 
Managed Database Products: SQL Server on VMs, full control over product (how we have it now). Can also provision a VM with SQL Server already installed. Azure SQL Database (PaaS). You don’t have to setup a VM with this service. Options for Elastic pool or single database. Elastic pool lets you scale up and down automatically. Azure SQL Managed Instance is lets you setup a VM with SQL Server already installed.
Connection string gives DB string. Geo-Replication is a disaster recovery feature, allowing synchronization between multiple dbs in different regions. 
Cosmos DB: Good for globally distributed data that may need to expand rapidly. Flexible schema. Used for many top-tier production systems, including Microsoft’s own services. Options to replicate data in selected datacenters. 
Blob data – can be stored in multiple tiers – Hot Tier (High cost storage cost, easy access), Cool Tier (Low storage cost, slower access), Archive Tier (Lowest storage cost, slowest data access). 
Exploring Azure Platform Solutions
Azure IoT (Internet of Things) Central:  Platform solution, managed app platform. 
Azure IoT Hub: Platform service that provides building blocks for building services in the cloud. Intended to operate at a large scale. Central message hub (distributed system) that devices send messages to. Can be scheduled or in real-time. 
Azure Sphere: Handles device security… not relevant to our work. 

