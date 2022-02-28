# DAS
Each computer has its own storage for its own use, like your PC's hard disk. Local storage, not accessible over network.
# SAN
Storage Area Network 
# NAS
Network Attached Storage

# General
Storage system has similar components to normal computer. CPU, RAM, Disk Subsystem, Motherboard, NIC. Use a special type of OS. 
Generally good idea to provision more space than you need. Benefits of centralized storage is the size can be changed on the fly without disrupting services. Client machine does not even know space is being added. Dont have to pay for unused storage. Deduplicaiton allows storage to share common resources. Redundant data is reduced. Just in case versus just in time. 
Latency might be worse, but data can be written to multiple disks allowing quicker access. New technology usually goes to enterprise storage systems first. Resiliency is important, redundant components can help prevent single points of failure. Easier to manage all storage in central location as well. 
Clients can boot up with no logical disks. Different types of storage is easier to configure. 
Backing up data is easier because it only has to be done in a single location. 
Snapshots are easier because it can be done on at a single location, rather than on every machine. 

# Types of storage
HDD - hard disk drive. Rotating platter (slower)
SSD - solid state drive (faster)
Servers usually use SAS (serial attached SCSI) which are more performant than SATA. 

# RAID 
Redundant array of inexpensive Disks. Combine multiple physical disks into a single logical unit. 
RAID 1 involves mirroring data between two disks. Read time is improved because the data can be retrieved from either disk. 