# Linux

## Tools

ls - list files in directory
   ls -lia
   joe@joes-mbp: learning $ ls -al
   total 48
   drwxr-xr-x  18 joe  staff   576 Nov 13 16:35 .
   drwxr-xr-x  20 joe  staff   640 Oct 27 17:48 ..
   -rw-r--r--@  1 joe  staff  8196 Jun 18 13:01 .DS_Store
   drwxr-xr-x  15 joe  staff   480 Nov 12 19:16 .git
   -rw-r--r--   1 joe  staff    97 Sep 21 14:49 .gitignore
   -rw-r--r--   1 joe  staff    29 Oct 18 18:04 README.md
ps - list running processes
   ps aux
   joe@joes-mbp: learning $ ps
   PID TTY           TIME CMD
   17957 ttys000    0:00.99 -bash
   34890 ttys001    0:00.01 -bash
   31369 ttys003    0:00.04 bash
   34989 ttys003    0:00.01 tmux
   34990 ttys004    0:00.01 -bash
   35046 ttys005    0:00.01 -bash
top - active view into cpu usage, memory, etc... by process
   Processes: 354 total, 2 running, 1 stuck, 351 sleeping, 2073 threads                                                                                         15:36:31
   Load Avg: 1.66, 2.52, 2.41  CPU usage: 2.13% user, 2.96% sys, 94.90% idle  SharedLibs: 492M resident, 106M data, 21M linkedit.
   MemRegions: 352063 total, 3804M resident, 273M private, 1619M shared. PhysMem: 11G used (1660M wired), 4839M unused.
   VM: 162T vsize, 3831M framework vsize, 1706882(0) swapins, 2443785(0) swapouts. Networks: packets: 149979504/163G in, 48139153/15G out.
   Disks: 291063829/4402G read, 201484938/3492G written.

   PID    COMMAND      %CPU TIME     #TH    #WQ  #PORT MEM    PURG  CMPRS  PGRP  PPID  STATE    BOOSTS            %CPU_ME %CPU_OTHRS UID  FAULTS     COW     MSGSENT
   159    WindowServer 15.9 65:54:42 24     7    4991- 572M-  27M+  144M   159   1     stuck    *0[1]             0.69216 0.63982    88   200312686+ 128644  2147483647
   78001  Google Chrom 0.9  08:07:12 19     1    358   473M   131M  44M    77990 77990 sleeping*1[590]           0.00000 0.00000    501  45925800   7458    604168618
   77990  Google Chrom 1.4  05:59:35 45     1    6421  348M+  1216K 93M    77990 1     sleeping *0[28208]         0.00000 0.00000    501  54638525+  125046  210814428+
   11799  Google Chrom 0.0  18:42.08 24     1    309   261M   0B    29M    77990 77990 sleeping*0[7]             0.00000 0.00000    501  2968758    6185    4703841
   14099  Google Chrom 0.0  00:52.43 22     1    396   247M   0B    139M   77990 77990 sleeping *0[7]
vmstat - list memory usage by process
   vmstat 
      procs -----------memory---------- ---swap-- -----io---- -system-- ------cpu-----

   r  b   swpd   free   buff  cache   si   so    bi    bo   in   cs us sy id wa st

   1  0      0 396528  38816 384036    0    0     4     2   38   79  0  0 99  0  0

   The fields are as follows:

   procs

   r - Number of processes for run time
   b - Number of processes in uninterruptible sleep
   memory

   swpd - Amount of virtual memory used
   free - Amount of free memory
   buff - Amount of memory used as buffers
   cache - Amount of memory used as cache
   swap

   si - Amount of memory swapped in from disk
   so - Amount of memory swapped out to disk
   io

   bi - Amount of blocks received in from a block device
   bo - Amount of blocks sent out to a block device
   system

   in - Number of interrupts per second
   cs - Number of context switches per second
   cpu

   us - Time spent in user time
   sy - Time spent in kernel time
   id - Time spent idle
   wa - Time spent waiting for IO

   30 08 ** * /home/pete/scripts/change_wallpaper
   The fields are as follows from left to right:

   Minute - (0-59)
   Hour - (0-23)
   Day of the month - (1-31)
   Month - (1-12)
   Day of the week - (0-7). 0 and 7 are denoted as Sunday
iostat - disk and cpu usage
   joe@joes-mbp: learning $ iostat
              disk0               disk5       cpu    load average
    KB/t  tps  MB/s     KB/t  tps  MB/s  us sy id   1m   5m   15m
   16.82   48  0.79    72.11    0  0.00   7 20 73  4.61 3.07 2.61
uptime
   1 5 15min intervals
   joe@joes-mbp: ~ $ uptime
   uptime: /dev/ttys002: No such file or directory
   13:33  up 118 days,  7:03, 3 users, load averages: 1.63 2.23 2.71
lsof
   joe@joes-mbp: learning $ lsof .
   COMMAND     PID USER   FD   TYPE DEVICE SIZE/OFF   NODE NAME
   Code\x20H 14607  joe  cwd    DIR   1,15      576 541547 .
   terraform 14633  joe  cwd    DIR   1,15      576 541547 .
   gopls     14661  joe  cwd    DIR   1,15      576 541547 .
   Code\x20H 14740  joe  cwd    DIR   1,15      576 541547 .
   bash      17957  joe  cwd    DIR   1,15      576 541547 .
   screen    31367  joe  cwd    DIR   1,15      576 541547 .
   bash      31369  joe  cwd    DIR   1,15      576 541547 .
   tmux      34889  joe  cwd    DIR   1,15      576 541547 .
   bash      34890  joe  cwd    DIR   1,15      576 541547 .
   tmux      34989  joe  cwd    DIR   1,15      576 541547 .
   bash      34990  joe  cwd    DIR   1,15      576 541547 .
   bash      35046  joe  cwd    DIR   1,15      576 541547 .
   lsof      78625  joe  cwd    DIR   1,15      576 541547 .
   lsof      78626  joe  cwd    DIR   1,15      576 541547 .
lsblk
df - list disk devices
   joe@joes-mbp: learning $ df
   Filesystem     512-blocks      Used Available Capacity iused      ifree %iused  Mounted on
   /dev/disk3s1s1  965595304  30106944 582216456     5%  502144 2911082280    0%   /
   devfs                 697       697         0   100%    1206          0  100%   /dev
   /dev/disk3s6    965595304   4194344 582216456     1%       2 2911082280    0%   /System/Volumes/VM
   /dev/disk3s2    965595304   1155584 582216456     1%    1216 2911082280    0%   /System/Volumes/Preboot
   /dev/disk3s4    965595304     17768 582216456     1%      48 2911082280    0%   /System/Volumes/Update
   /dev/disk1s2      1024000     12328    985264     2%       1    4926320    0%   /System/Volumes/xarts
   /dev/disk1s1      1024000     12512    985264     2%      24    4926320    0%   /System/Volumes/iSCPreboot
   /dev/disk1s3      1024000      4176    985264     1%      52    4926320    0%   /System/Volumes/Hardware
   /dev/disk3s5    965595304 345738008 582216456    38% 1824345 2911082280    0%   /System/Volumes/Data
   map auto_home           0         0         0   100%       0          0  100%   /System/Volumes/Data/home
   /dev/disk5s2        44408     44408         0   100%      31 4294967248    0%   /private/tmp/KSInstallAction.sY
du
   joe@joes-mbp: learning $ du -ah | sort -rh
   1.1G    .
   475M    ./projects
   410M    ./projects/nextjs/myapp
   410M    ./projects/nextjs
   392M    ./projects/nextjs/myapp/.next
   389M    ./terraform/aws/.terraform/providers/registry.terraform.io/hashicorp/aws/4.67.0/darwin_arm64/terraform-provider-aws_v4.67.0_x5
   389M    ./terraform/aws/.terraform/providers/registry.terraform.io/hashicorp/aws/4.67.0/darwin_arm64
   389M    ./terraform/aws/.terraform/providers/registry.terraform.io/hashicorp/aws/4.67.0
   389M    ./terraform/aws/.terraform/providers/registry.terraform.io/hashicorp/aws
   389M    ./terraform/aws/.terraform/providers/registry.terraform.io/hashicorp
   389M    ./terraform/aws/.terraform/providers/registry.terraform.io
   389M    ./terraform/aws/.terraform/providers
   389M    ./terraform/aws/.terraform
find
   find . -name 'node_modules' -type d
strace - trace the syscalls of a specific command
fg/bg/jobs
   kubectl port-forward...
   send to background: ctrl z -> bg -> query with jobs
   or run with &
   fg %1

chmod - set permissions on a file
   sudo chmod 777 myfile - (4 - read, 2 - write, 1 execute)
chown - set ownership of file
   sudo chown myuser:mygroup myfile
ssh
   ssh-keygen
   ssh-copy-id user@host
   ssh -i ~/.ssh/id_rsa user@server
   ssh user@server
scp
   scp myfile.txt username@remotehost.com:/remote/directory
rsync
   copy from local to remote
   rsync /local/directory username@remotehost.com:/remote/directory
   copy to remote from local
   rsync username@remotehost.com:/remote/directory /local/directory
simple http server
   python -m SimpleHTTPServer
start nfs
   sudo service nfsclient start
   sudo mount server:/directory /mount_directory
samba -  share files between linux, mac, windows
   sudo apt update
   sudo apt install samba
   sudo smbpasswd -a [username]
   mkdir /my/directory/to/share
   sudo service smbd restart
   smbclient //HOST/directory -U user
   sudo mount -t cifs servername:directory mountpount -o user=username,pass=password

Orphan Processes

When a parent process dies before a child process, the kernel knows that it's not going to get a wait call, so instead it makes these processes "orphans" and puts them under the care of init (remember mother of all processes). Init will eventually perform the wait system call for these orphans so they can die.

Zombie Processes

What happens when a child terminates and the parent process hasn't called wait yet? We still want to be able to see how a child process terminated, so even though the child process finished, the kernel turns the child process into a zombie process. The resources the child process used are still freed up for other processes, however there is still an entry in the process table for this zombie. Zombie processes also cannot be killed, since they are technically "dead", so you can't use signals to kill them. Eventually if the parent process calls the wait system call, the zombie will disappear, this is known as "reaping". If the parent doesn't perform a wait call, init will adopt the zombie and automatically perform wait and remove the zombie. It can be a bad thing to have too many zombie processes, since they take up space on the process table, if it fills up it will prevent other processes from running.

ps a
kubectl port-forward...
send to background: ctrl z -> bg -> query with jobs
or run with &
fg %1

dd if=/home/pete/backup.img of=/dev/sdb bs=1M count=2
if=file - Input file, read from a file instead of standard input
of=file - Output file, write to a file instead of standard output
bs=bytes - Block size, it reads and writes this many bytes of data at a time. You can use different size metrics by denoting the size with a k for kilobyte, m for megabyte, etc, so 1024 bytes is 1k
count=number - Number of blocks to copy

differenet types of filesystems
df -T

## Create partition

parted > select /dev/sdc/ > mkpart
name, type, start, finish (type is just an indicator of what type of filesystem can be created on
the partition, it does not actually create the filesystem, use mkfs.ext4 for that)
rm to delete. Make sure correct disk is select

## Format partition

mkfs.ext4 (or whatever fs type) /dev/sdc2

## Mount

mkdir /var/lib/mount-target
mount /dev/sdc2 /var/lib/mount-target

## Umount

umount /var/lib/mount-target

## Get processes using system

fuser -v /var/lib/mount-target

## Kill process

kill <process id>

sudo blkid

## swap space

First make sure we don't have anything on the partition
Run: mkswap /dev/sdb2 to initialize swap areas
Run: swapon /dev/sdb2 this will enable the swap device
If you want the swap partition to persist on bootup, you need to add an entry to the /etc/fstab file. sw is the filesystem type that you'll use.
To remove swap: swapoff /dev/sdb2

use df for utilization
use du to see storage by file

What is an inode?

An inode (index node) is an entry in this table and there is one for every file. It describes everything about the file, such as:

File type - regular file, directory, character device, etc
Owner
Group
Access permissions
Timestamps - mtime (time of last file modification), ctime (time of last attribute change), atime (time of last access)
Number of hardlinks to the file
Size of the file
Number of blocks allocated to the file
Pointers to the data blocks of the file - most important!

bios bootloader kernel init

strace ls

# performance tools

uptime - CPU usage over time, 1m 5m 15m
13:29  up 75 days,  5:59, 2 users, load averages: 2.53 2.70 2.91

top
Processes: 410 total, 3 running, 407 sleeping, 2515 threads                                                                                                        13:31:13
Load Avg: 2.70, 2.76, 2.90  CPU usage: 7.48% user, 10.10% sys, 82.41% idle  SharedLibs: 534M resident, 102M data, 24M linkedit.
MemRegions: 377318 total, 2165M resident, 87M private, 5322M shared. PhysMem: 15G used (1532M wired), 86M unused.
VM: 185T vsize, 3831M framework vsize, 1022048(0) swapins, 1520209(0) swapouts. Networks: packets: 88293278/96G in, 28435665/10G out.
Disks: 180791313/2719G read, 119287278/2115G written.

ps

iostat
              disk0       cpu    load average
    KB/t  tps  MB/s  us sy id   1m   5m   15m
   16.89   46  0.76   2  7 91  3.00 3.06 3.00

strace
tcpdump

for i in $(docker container ls --format "{{.ID}}"); do docker inspect -f '{{.State.Pid}} {{.Name}}' $i; done

vmstat
iostat
top -o %MEM
ps
lsof

ps aux | grep 14573
ps aux

find . -name "test*"
lsof ~/work/test.swp
sudo lsof returns more results
sudo lsof +D

1 5 15min intervals
joe@joes-mbp: ~ $ uptime
uptime: /dev/ttys002: No such file or directory
13:33  up 118 days,  7:03, 3 users, load averages: 1.63 2.23 2.71

ISP - Your internet service provider, the company you pay to get Internet at your house.
Router - The router allows each machine on your network to connect to the Internet. In most modern routers, you can connect via wireless or an Ethernet cable.
WAN - Wide Area Network, this is what we call the network that encompasses everything between your router and a wider network such the Internet.
WLAN - Wireless Local Area Network, this is the network between your router and any wireless devices you may have such as laptops.
LAN - Local Area Network, this is the network between your router and any wired devices such as Desktop PCs.
Hosts - Each machine on a network is known as a host.

Application Layer

The top layer of the TCP/IP model. It determines how your computer's programs (such as your web browser) interface with the transport layer services to view the data that gets sent or received.

This layer uses:

HTTP (Hypertext Transfer Protocol) - used for the webpages on the Internet.
SMTP (Simple Mail Transfer Protocol) - electronic mail (email) transmission
Transport Layer

How data will be transmitted, includes checking the correct ports, the integrity of the data, and basically delivering our packets.

This layer uses:

TCP (Transmission Control Protocol) - reliable data delivery
UDP (User Datagram Protocol) - unreliable data delivery
Network Layer

This layers specifies how to move packets between hosts and across networks.

This layer uses:

IP (Internet Protocol) - Helps route packets from one machine to another.
ICMP (Internet Control Message Protocol) - Helps tell us what is going on, such as error messages and debugging information.
Link Layer

This layer specifies how to send data across a physical piece of hardware. Such as data travelling through Ethernet, fiber, etc.

The lists above of protocols each layer uses is not extensive and you'll encounter many other protocols that come into play.

In the following lessons, we will dive through each of these layers and discuss how our packet traverses through the network in the eyes of the TCP/IP model (there are many perspectives on how a packet travels across networks, we won't look at them all, but be aware that they exist).

Pete sends Patty an email: this data gets sent to the transport layer.
The transport layer encapsulates the data into a TCP or UDP header to form a segment, the segment attaches the destination and source TCP or UDP port, then the segment is sent to the network layer.
The network layer encapsulates the TCP segment inside an IP packet, it attaches the source and destination IP address. Then routes the packet to the link layer.
The packet then reaches Pete's physical hardware and gets encapsulated in a frame. The source and destination MAC address get added to the frame.
Patty's receives this data frame through her physical layer and checks each frame for data integrity, then de-encapsulates the frame contents and sends the IP packet to the network layer.
The network layer reads the packet to find the source and destination IP that was previously attached. It checks if its IP is the same as the destination IP, which it is! It de-encapsulates the packet and sends the segment to the transport layer.
The transport layer de-encapsulates the segments, checks the TCP or UDP port numbers and makes a connection to the application layer based on those port numbers.
The application layer receives the data from the transport layer on the port that was specified and presents it to Patty in the form of the final email message

CIDR
   CIDR (classless inter-domain routing) is used to represent a subnet mask in a more compact way. You may see subnets notated in CIDR notation, where a subnet such as the 10.42.3.0/255.255.255.0 is written as 10.42.3.0/24 which just means it includes both the subnet prefix and the subnet mask.

   Remember an IP address consists of 4 bytes or 32 bits, CIDR indicates the amount of bits used as the network prefix. So 123.12.24.0/23 means that the first 23 bits are used. Well what does that mean? How many hosts is that?

   A simple trick is to subtract the total of bits an IP address can have (32) from the CIDR address (23), so that leaves 9 bits, 2^9 = 512, but we have to remove 2 addresses (subnet address and broadcast address) so we have 510 usable hosts.

To show interface information for all interfaces

$ ip link show
To show the statistics of an interface
$ ip -s link show eth0
To show ip addresses allocated to interfaces
$ ip address show
To bring interfaces up and down
$ ip link set eth0 up
$ ip link set eth0 down
To add an IP address to an interface
$ ip address add 192.168.1.1/24 dev eth0

## Review route tables

joe@joes-mbp: learning $ route get google.com
   route to: yyz08s10-in-f174.1e100.net
destination: default
       mask: default
    gateway: 192.168.5.1
  interface: en0
      flags: <UP,GATEWAY,DONE,STATIC,PRCLONING,GLOBAL>
 recvpipe  sendpipe  ssthresh  rtt,msec    rttvar  hopcount      mtu     expire
       0         0         0         0         0         0      1500         0 

dhcp - dynamically provisions ip address on a network. 
   sudo dhclient

## partitions

sudo parted -l