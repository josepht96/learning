**Linux Notes - Cleaned and Organized**

---

## Tools

### File and Directory Listing

* `ls` - list files in directory

  * Example: `ls -al`
  * Lists all files, including hidden ones, in long format.

### Process Management

* `ps` - list running processes

  * Example: `ps aux`
  * Shows detailed information about all running processes.
* `top` - live system stats (CPU, memory, processes)

  * Example: `top`
  * Real-time view of running processes.
* `vmstat` - system memory and process info

  * Example: `vmstat 1`
  * Displays system performance every second.
* `iostat` - CPU and I/O statistics

  * Example: `iostat -xz 1`
  * Extended disk stats with CPU usage updated every second.
* `uptime` - system uptime and load average

  * Example: `uptime`
  * Shows how long the system has been up.
* `lsof` - list open files

  * Example: `lsof +D /home/user`
  * Lists open files in the specified directory.
* `lsblk` - list block devices

  * Example: `lsblk`
  * Shows block device hierarchy.
* `df` - show disk usage by device

  * Example: `df -h`
  * Human-readable output of disk usage.
* `du` - show disk usage by file or directory

  * Example: `du -ah | sort -rh | head`
  * Lists top disk usage files/directories.
* `find` - search for files/directories

  * Example: `find . -name '*.log'`
  * Finds all `.log` files recursively.
* `strace` - trace system calls

  * Example: `strace ls`
  * Traces syscalls made by `ls`.
* `jobs`, `fg`, `bg` - manage background/foreground jobs

  * Example: `./longtask &` then `jobs`, `fg %1`
  * Background a task and manage it.

### Permissions and Ownership

* `chmod` - change permissions

  * Example: `chmod 755 script.sh`
  * Gives owner rwx, group/others rx.
* `chown` - change file ownership

  * Example: `chown user:group file`
  * Changes file ownership.

### SSH and File Transfer

* `ssh` - connect to remote machines

  * Example: `ssh user@host`
  * Opens an SSH session.
* `scp` - copy files between local and remote

  * Example: `scp file.txt user@host:/remote/path`
* `rsync` - sync files between local and remote

  * Example: `rsync -avz ./dir user@host:/remote/dir`

### Networking Tools

* `ip` - show/set network interfaces, addresses, and routes

  * Example: `ip a`
  * Shows IP addresses.
* `route` - view routing table

  * Example: `route -n`
  * Displays network routes.
* `dhclient` - obtain IP via DHCP

  * Example: `sudo dhclient`
  * Requests IP configuration.
* `ping`, `arp`, `tcpdump` - diagnostics and packet capture

  * Example: `ping 8.8.8.8`, `arp -a`, `tcpdump -i eth0`

### Mounting and Partitions

* `mount`, `umount` - mount and unmount devices

  * Example: `mount /dev/sdb1 /mnt`
  * Example: `umount /mnt`
* `parted`, `mkfs.ext4`, `fuser`, `blkid`

  * Example: `sudo parted /dev/sdb`
  * Example: `mkfs.ext4 /dev/sdb1`
  * Example: `fuser -v /mnt`
  * Example: `blkid`
* `mkswap`, `swapon`, `swapoff`

  * Example: `mkswap /dev/sdb2`
  * Example: `swapon /dev/sdb2`
  * Example: `swapoff /dev/sdb2`

### Misc

* `dd` - copy and convert files

  * Example: `dd if=/dev/sda of=/backup.img bs=1M count=2`
* `sudo` - execute commands as another user

  * Example: `sudo apt update`

---

## System Concepts

### Orphan Processes

* When parent dies, child becomes orphan, adopted by `init`.

### Zombie Processes

* Child terminates but parent hasn’t waited. Eventually reaped.

---

## Networking

### Key Terms

* **ISP** - Internet provider
* **Router** - Connects LAN to WAN
* **WAN/LAN/WLAN** - Network types
* **Host** - Device on network

### TCP/IP Model

* **Application Layer**: HTTP, SMTP
* **Transport Layer**: TCP, UDP
* **Network Layer**: IP, ICMP
* **Link Layer**: Ethernet, Wi-Fi

### CIDR

* `10.0.0.0/24` => 256 addresses, 254 usable
* Formula: `2^(32 - mask bits) - 2`

### Interface Commands

```bash
ip link show
ip address show
ip link set eth0 up
ip address add 192.168.1.1/24 dev eth0
```

### Routing

```bash
route get google.com
```

### View LAN Devices

```bash
ifconfig | grep broadcast
ping 192.168.1.255
arp -a
```

---

## File Systems

### File System Types

* `df -T` - Show FS type

### Partitioning

```bash
sudo parted /dev/sdc
(parted) mkpart primary ext4 0% 100%
```

### Format, Mount, Unmount

```bash
mkfs.ext4 /dev/sdc1
mkdir /mnt/mydisk
mount /dev/sdc1 /mnt/mydisk
umount /mnt/mydisk
```

### Tools

* `fuser -v /mnt` - show processes using path
* `kill <pid>` - terminate process

---

## Inodes

* Inode contains metadata:

  * Type, permissions, owner, timestamps
  * Block pointers for data location

---

## Performance Tools

* `uptime` - system load
* `top` - interactive process monitor
* `vmstat` - memory, processes
* `iostat` - I/O statistics
* `ps`, `lsof`, `strace`, `tcpdump` - system analysis

---

## Misc Commands

```bash
scp file.txt user@host:/path
rsync -avz /src user@host:/dest
python3 -m http.server 8000  # HTTP server
```

---

## Samba (File Sharing)

```bash
sudo apt install samba
sudo smbpasswd -a username
sudo service smbd restart
smbclient //HOST/dir -U user
sudo mount -t cifs //HOST/dir /mnt -o user=user,pass=pass
```

---

Let me know if you'd like additional examples or export options!
