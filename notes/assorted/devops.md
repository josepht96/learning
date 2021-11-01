# DevOps Prerequisites 
Bourne Shell - sh
Bourne again Shell - bash
Echo $SHELL
cmds:
pwd - current directory
cd path; mkdir name; pwd
mkdir -p <full path>
rm -r <path> - remove directory content
cp -r dir1 dir2
touch text.txt
cat > text.txt
cat >> text.txt (append)
type (control D to leave)
cat text.txt to print
mv path2file path2file2
rm file

vi/vim:
if arrow keys insert chars, do vi $HOME/.exrc → insert set nocompatible → save
command mode: (escape)
insert mode: write content (i)
x delete char
yy
: commands
:w save
:q discard
:wq save quit
yy copy line
dd delete line
p paste

User accounts
whoami
id
ssh user@ip
su switch user
sudo ls /root
curl url -o, saves to file
wget url -O /path/file.txt
ls /etc/*release*
cat file

rpm
rpm -i package.rpm
rpm -e telnet.rpm (uninstall)
rpm -q telnet.rpm (query)
yum
installs package and necessary dependencies
yum repolist
yum -q package name

services
service httpd start
systemctl start httpd
systemctl stop httpd
systemctl status httpd
systemctl enable httpd
systemctl disable httpd
etc/systemd/system
myapp.service
[Unit]
Description=My app
[Service]
ExecStart=/usr/bin/python3 /file/path/main.go
Restart=always
[Install]
WantedBy=multi-user.target

# Networking
gateway - The gateway converts information, data or other communications from one protocol or format to another. A router may perform some of the functions of a gateway. An Internet gateway can transfer communications between an enterprise network and the Internet. 
bridge - A network bridge is a computer networking device that creates a single, aggregate network from multiple communication networks or network segments. 
router -  A router is a networking device that forwards data packets between computer networks. 
switch - A network switch (also called switching hub, bridging hub, and, by the IEEE, MAC bridge) is networking hardware that connects devices on a computer network by using packet switching to receive and forward data to the destination device.

route - display Kernel IP routing table
ip link - list and modify interfaces on host
ip addr - see ip addresses assigned to interfaces
inet6 2600:1700:831:1c00:19d1:11b8:faa5:109e/64 scope global temporary dynamic
ip addr add <ip> - set ip addresses on interfaces
sudo ip addr add <ip> dev <NIC name> (get it from ip addr)
ip route ad <ip> via <outbound ip>
valid til restart
set values in /etc/sysctl.conf to persist changes

system must have NIC or VNIC to assign ip address
subnets can communicate via routing over a shared ip
172.16.238.15/24 + 172.16.238.16/24
172.16.239.15/24 + 172.16.239.16/24 
These cannot communicate, unless there is a server that can reach both:
machine with inets: 172.16.238.10/24 and 172.16.239.10/24
route traffic from subnets over host via:
ip route add 172.16.239.0/24 via 127.16.238.10
ip route add 172.16.238.0/24 via 127.16.239.10
This allows the first ip subnet to communicate using the middle machine


# DNS
Tell system A that system B has a specific name
cat >> /etc/hosts
/etc/hosts maps ips to domains
cat /etc/resolv.conf
/etc/resolv.conf maps domains to ips
<ip> <name>
add nameserver by:
nameserver 8.8.8.8
Create DNS server so all hosts can look up DNS names in single place
add entry to /etc/resolv.conf, nameserver <ip> then it will look up names on server
suffix of domain url indicates meaning:
.com = commercial/general
.net = network
.edu = education
.org = organization
.io = tech
. .com google www (maps is sub domain, docs, etc..)
nslookup query hostname from dns server

# Web Servers
framework good for local development, web server needed for production usually
Static and dynamic websites, web server and application server
Static html → site is served, then minimal activity with web server
dynamic websites will have more interactions with the web server, which is likely serving jsonified sql data.
nginx is a good static web server, apache tomcat is a good application (dynamic server)
apache httpd -
sudo service httpd start, root is /var/www/html/<files>
Virtual box networking
IP address assigned when connecting to network, assigned to interface
Can have multiple IP addresses on a comp, due to different interfaces
Host-Only network allows VMs to communicate on host VM
Host shows up on Host-Only network
Network Address Translation (NAT) – translates each VM source IP with host IP address, so it appears as all the requests are coming from the host. When host receives response, host changes IP back to internal VM IP address. Systems from outside host cannot directly reach the VMs.  (All over LAN)
Bridge work makes the VMs appear as independent hosts, they can be accessed directly over the LAN
Host-Only network VMs can access public internet with IP forwarding
Port forwarding – forwards from port on host to port on VM. Host@8080 à VM@4000
 
# Application Basics
Python: multiple versions can be installed on same machine
Pip is package manager, pip2, pip3
Pip install -r requirements.txt – install all necessary packages
 
# Git
Git remote add <name (origin usually)> <url>
You can add multiple remotes, so git push origin branch is just saying push to origin remote, but could do git push origin2 branch
 
# IP and Ports
Ethernet cable will add an IP on the network for the comp’s adapter. If you connect Wi-Fi, another IP will be added to the comp on a different NIC. Ports are slots on the NIC. If you have two IPs, you can set an IP address to run on. Accessing port through different interface wont work. Using 0.0.0.0 will cause the app to be accessible through the specified port on every interface. When no option is specified, a server will listen on 127.0.0.1 (loopback address, “lo”). Every host has loopback interface. Localhost is standard name for loopback. Second machine trying to connect to the localhost IP is of course going to fail (as it will look for the service on itself).
 
 
# SSL & TLS basics
PKI: Guarantee trust between two parties. Ensure communication is encrypted and server is authentic (who it says it is). Symmetric encryption involves keys on both server and client. Asymmetric encryption contains a public lock and a private key. Server side can only lock files that can only be unlocked by the private key on client side.
Private key, public lock. Public lock can be locked be any private key, but only unlocked with its private key.
Server generates public and private key pair
Sniffer gets public key when client connects to server
Client encrypts symmetric key using the public key provided by the server
Sniffer has public key with client key, doesn’t have server key so can decrypt.
The public server key is sent in a certificate that is verified. 3rd parties (Certificate Authority CA) can sign the certificate. Self-signed certs are not safe. Browsers validate certificates automatically.
Public keys are .crt or .pem, private keys usually have .key or -key.pem.
Encrypt stuff with public key, only private key can decrypt. Say, you know hypothetically, you wanted to make a encrypted transaction over the web.
Vendor sends you his public key, you write a message and encrypt it. Send it back to him and only he can decrypt it with his private key.
 
/etc/httpd/certs/app01.crt
 
 

# Linux ----------------------------------------------------------------------------------------------------------------------
Linux Kernel:
Process management, memory management, device drivers, system calls and security.
Library analogy: processes are the students, kernel is librarian, books and stuff are the hardware. 
Kernel space: Kernel, device drivers - section of library where only librarian has access
User space: applications, programs, restricted access
Apps make system calls to the kernel space to access restricted resources.

Cores * threads == how many parallel processes can run at the same time
System/hardware calls:
sudo dmesg
udevadm info --query=path
udevadm monitor
lspci
lsblk
lscpu
lsmem --summary
ls -l /sbin/init - boot type
	 ls -ld /path/to/file
	df -hP
	systemctl get-default
	lshw
	su 
	useradd -u 1009 -g 1009 -d /home/robert -s /bin/bash -c “some comment” bob
	id
	chmod sets permissions
	chown changes the owner

sda  	8:0	0 447.1G  0 disk  
├─sda1   8:1	0   1.9G  0 part [SWAP]
├─sda2   8:2	0  18.6G  0 part /
└─sda3   8:3	0 426.6G  0 part /home
sdb  	8:16   0 476.9G  0 disk  
├─sdb1   8:17   0   529M  0 part  
├─sdb2   8:18   0   100M  0 part /boot/efi
├─sdb3   8:19   0	16M  0 part  
└─sdb4   8:20   0 476.3G  0 part  
Architecture:                	x86_64
CPU op-mode(s):              	32-bit, 64-bit
Byte Order:                  	Little Endian
Address sizes:               	43 bits physical, 48 bits virtual
CPU(s):                      	16
On-line CPU(s) list:         	0-15
Thread(s) per core:          	2
Core(s) per socket:          	8
Socket(s):                   	1
NUMA node(s):                	1
Vendor ID:                   	AuthenticAMD
CPU family:                  	23
Model:                       	113
Model name:                  	AMD Ryzen 7 3700X 8-Core Processor
Stepping:                    	0
Frequency boost:             	enabled
CPU MHz:                     	2200.000
CPU max MHz:                 	4426.1709
CPU min MHz:                 	2200.0000
BogoMIPS:                    	7186.52
Virtualization:              	AMD-V

Memory block size:   	128M
Total online memory:  	32G
Total offline memory:  	0B


# Boot order 
BIOS Post -> Boot loader (GRUB2) -> Kernel initialization -> INIT Process (systemd)
BIOS posts to hardware to ensure they are functioning correctly. 
Boot loader -> boot into selected OS, loads select OS into memory. GRUB2 is the primary boot loader. 
Run level determines whether the OS boots in headless or GUI mode. (3, 5 levels)

# File types
3 types of files: regular, directory, special
Regular: text files, scripts, normal stuff
Directory: stores other files within
Special:
character filles - device files
block files - files representing block devices, reads/writes to device in blocks, hard disks and RAM
inks - hard link associates two or more files that share the same data on disk, symbolic/soft	link is basically a shortcut, pointer to another file. Deleting a hard link will delete data. Deleting a sym/soft link is like deleting a shortcut
Socket files
Named pipes

Can identify files by running: ls -ld /path/to/file
first letter is the file top

# File Hierarchy
/home contains home directories for all users, user specific files
/opt contains 3rd party programs
/mnt place to mount a repository
/tmp contains temporary files, copy files from mnt
/media is where peripherals are mounted (df -hP)
/dev contains files for devices, like hard disks, mouse and keyboard…
/bin contains basic programs and binaries
/etc used to store configuration files
/lib contains shared libraries
/usr contains files that are used by all users
/var contains files/logs for system issues, cached data

# Apt vs Apt-get 
Installs package dependencies (improvement over dpkg)
apt is more user friendly, overall a better tool than apt-get (??)
Relies on a software repository (/etc/apt/sources.list)
Source can be local or remote location. 
apt update - refreshes repository (from all available sources)
run after adding new sources to ensure they're latest
apt upgrade
apt edit-sources - edit repos
apt install
apt remove
apt search 

# Shell
tar -cf <name.tar> file 1 file 2 (tar name and files to be zipped)
tar -tf test.tar - see contents
tar -xf test.tar - extract contents
tar -zcf file name, files and compress

searching for files:
	locate filename
	find 
	GREP find matches
	grep string file.txt
	grep -i string file.txt
	grep -r “string” dir

# Users 
username, UID, GID, /home/username, default shell: /bin/sh - all stored in /etc/passwd
Superuser, UID (0) unrestricted access and control
System accounts, created during install to manage utilities
Service accounts, manage services like nginx
/etc/sudoers controls user privileges with sudo command

# SSH and SCP 
ssh-keygen -t rsa
Passwordless key pair:
Create key pair on client (/home/username/.ssh/id_rsa.pub)
Private key (/home/username/.ssh/id_rsa)
Copy public key to remote server - ssh-copy-id username@server name/ip
Can view public keys view /home/username/.ssh/authorized_keys
scp myfile.txt remoteuser@remoteserver:/remote/folder/

# Firewalls
sudo apt install iptables
sudo iptables -l 
sudo iptables -A INPUT -p tcp -s <client server> --dport 22 -j ACCEPT
sudo iptables -A INPUT -p tcp --dport 22 -j REJECT
-I inserts item at the top of the table
