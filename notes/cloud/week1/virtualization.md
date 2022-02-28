# types of hypervisors 
type 1 - bare metal hypervisor. Go buy server, install hypervisor directly on server. Essentially replaces standard OS installation. 
VMWare, MS, Citrix, etc... HyperV is an example of the software you'd install. 
type 2 - More typical virtualization method, running hypervisor ontop of standard OS. OS > Hypervisor > VM

# vm files and state
CPU, live memory, processes, network bandwidth, etc... happen on the host. Utilize storage adapters to send data to certain location. VM writing to C:/ drive is redirected to a different location,
since VM doesn't actually have a hard disk attached to it. 

# virtual CPU
vCPU means VM has access to physical CPU core. 2vCPU = 2 physical CPU cores. Access is not restricted to that VM. 

# virtual memory
Windows OS on VM cannot tell its memory is virtual, bound by set RAM limits. HyperVisor maps guest memory to physical memory. VM with 4gbs wont take 4gbs until it actually needs it. Reserved memory cannot be used by other VMs -- try to avoid when possible. 

# virtual networking
vNIC attached to VM. OS expects same hardware that physical server would have. Windows sees drivers for a NIC, except its a virtual NIC (fake hardware, presented as real). Hypervisor has a virtual switch that allows vNIC to communicate with each other, as long as they are on the same vLAN. vSwitch can have uplink to physical switch so VMs can communicate with host. Physical NIC is attached to vSwitch. pNIC has connection to physical Switch (pSwitch).

# virtual storage
virtual SCSI controller interfaces with storage adapter to physical disks. Windows uses SCSI controller normally, so Windows VMs need vSCSI controller. Hypervisor controls writing data to physical disk. VM can be moved to new host if datastore is shared by hosts. 
