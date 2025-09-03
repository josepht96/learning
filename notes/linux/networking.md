**Linux Notes - Cleaned and Organized**

---

## Tools

### File and Directory Listing

* `ls` - list files in directory

  * `ls -lia`
  * `ls -al`

### Process Management

* `ps` - list running processes

  * `ps aux`
* `top` - live system stats (CPU, memory, processes)
* `vmstat` - system memory and process info
* `iostat` - CPU and I/O statistics
* `uptime` - system uptime and load average
* `lsof` - list open files
* `lsblk` - list block devices
* `df` - show disk usage by device
* `du` - show disk usage by file or directory
* `find` - search for files/directories
* `strace` - trace system calls
* `jobs`, `fg`, `bg` - manage background/foreground jobs

### Permissions and Ownership

* `chmod` - change permissions (e.g., `chmod 777 myfile`)
* `chown` - change file ownership (e.g., `chown user:group myfile`)

### SSH and File Transfer

* `ssh` - connect to remote machines
* `scp` - copy files between local and remote
* `rsync` - sync files between local and remote

### Networking Tools

* `ip` - show/set network interfaces, addresses, and routes
* `route` - view routing table
* `dhclient` - obtain IP via DHCP
* `ping`, `arp`, `tcpdump` - diagnostics and packet capture

### Mounting and Partitions

* `mount`, `umount` - mount and unmount devices
* `parted`, `mkfs.ext4`, `fuser`, `blkid` - partition and filesystem tools
* `mkswap`, `swapon`, `swapoff` - manage swap space

### Misc

* `dd` - copy and convert files (e.g., `dd if=... of=... bs=1M count=2`)
* `sudo` - execute commands as another user (typically root)

---

## System Concepts

### Orphan Processes

* Occur when parent dies before the child.
* Adopted by `init`, which performs `wait()`.

### Zombie Processes

* Child terminates but parent hasn’t `wait()`ed.
* Resources are freed, but a process entry remains.
* Reaped when parent or `init` calls `wait()`.

---

## Networking

### Key Terms

* **ISP** - Your internet service provider
* **Router** - Connects LAN to WAN
* **WAN** - Wide Area Network
* **LAN** - Local Area Network
* **WLAN** - Wireless LAN
* **Host** - Any device on the network

### OSI Model (Open Systems Interconnection)

* Layered model for networking:

  1. Physical
  2. Data Link
  3. Network
  4. Transport
  5. Session
  6. Presentation
  7. Application
* Helps design and troubleshoot networks.

### Firewalls

* Network security systems that control traffic.
* Types: host-based (iptables), network-based (hardware firewalls).

### NAT (Network Address Translation)

* Translates private IPs to public IPs.
* Enables multiple devices to share one IP.

### DNS (Domain Name System)

* Resolves domain names to IP addresses.

### DNS Exfiltration

* Data theft using DNS requests to bypass firewalls.

### DNS Configs

* Stored in `/etc/resolv.conf` (Linux).
* Example: `nameserver 8.8.8.8`

### ARP (Address Resolution Protocol)

* Maps IP addresses to MAC addresses.
* View: `arp -a`

### DHCP (Dynamic Host Configuration Protocol)

* Automatically assigns IPs.
* Client: `sudo dhclient`

### Multiplexing

* Combines multiple signals over a single medium.
* Example: TCP multiplexes multiple connections over one IP.

### Traceroute

* Displays path packets take to a destination.
* Usage: `traceroute example.com`

### Nmap

* Network scanner for host discovery and port scanning.
* Usage: `nmap -sS 192.168.1.1`

### Intercepts

* Techniques like MITM (Man-In-The-Middle) to intercept network traffic.

### VPN (Virtual Private Network)

* Encrypts traffic and masks IP address.

### Tor

* Anonymity network routing through volunteer relays.

### Proxy

* Forwards traffic on behalf of a client.

### BGP (Border Gateway Protocol)

* Routes traffic between autonomous systems on the internet.

### Network Traffic Tools

* Tools: `tcpdump`, `wireshark`, `iftop`, `nethogs`

### HTTP/S

* HyperText Transfer Protocol / Secure (encrypted via TLS).

### SSL/TLS

* Secure communication protocols for encryption and authentication.

### TCP/UDP

* TCP: Reliable, ordered
* UDP: Unreliable, faster

### ICMP

* Used for diagnostics (e.g., `ping`, `traceroute`)

### Mail Protocols

* **SMTP** - Send mail
* **IMAP/POP3** - Retrieve mail

### SSH (Secure Shell)

* Secure remote login protocol.
* Usage: `ssh user@host`

### Telnet

* Insecure remote login (mostly obsolete).

### IRC (Internet Relay Chat)

* Real-time chat protocol.

### FTP/SFTP

* FTP: Unencrypted file transfer
* SFTP: Encrypted, over SSH

### RPC (Remote Procedure Call)

* Allows a program to execute on another machine.

### Service Ports

* Common ports: 22 (SSH), 80 (HTTP), 443 (HTTPS), 25 (SMTP), 53 (DNS)

### HTTP Header

* Metadata in HTTP request (e.g., `User-Agent`, `Host`)

### HTTP Response Header

* Metadata in HTTP response (e.g., `Content-Type`, `Server`)

### UDP Header

* Contains source/destination port, length, and checksum

### Broadcast and Collision Domains

* **Broadcast domain**: devices that receive all broadcast frames
* **Collision domain**: segment where data packets can collide

### Root Stores

* OS/Browser-trusted Certificate Authorities (CAs)

### CAM Table Overflow

* Switch attack flooding MAC addresses to force broadcast behavior

---

## TCP/IP Model

* **Application Layer**: HTTP, SMTP
* **Transport Layer**: TCP, UDP
* **Network Layer**: IP, ICMP
* **Link Layer**: Ethernet, Wi-Fi

### CIDR (Classless Inter-Domain Routing)

* `10.0.0.0/24` means 24 bits are for network.
* `2^(32 - subnet bits) - 2 = usable hosts`

### Network Interface Management

```bash
ip link show                     # show interfaces
ip -s link show eth0             # show interface stats
ip address show                  # show IPs
ip address add ... dev eth0      # add IP
ip link set eth0 up/down         # enable/disable interface
```

### View Routes

```bash
route get google.com
```

### View LAN Devices

```bash
ifconfig | grep broadcast
ping <broadcast-ip>
arp -a
```

---

## File Systems

### File System Types

* View with: `df -T`

### Create Partition

```bash
sudo parted
> select /dev/sdc
> mkpart name type start end
```

### Format and Mount

```bash
mkfs.ext4 /dev/sdc2
mkdir /var/lib/mount-target
mount /dev/sdc2 /var/lib/mount-target
```

### Unmount and Manage

```bash
umount /var/lib/mount-target
fuser -v /var/lib/mount-target
kill <pid>
```

---

## Inodes

* Describes metadata for each file:

  * File type, owner, group
  * Permissions, timestamps (atime, mtime, ctime)
  * Size, blocks, pointers to data

---

## Performance Tools

* `uptime`, `top`, `vmstat`, `iostat`, `ps`, `lsof`, `strace`, `tcpdump`

---

## Misc Commands

```bash
scp myfile.txt user@host:/path
rsync /src user@host:/dest
python -m SimpleHTTPServer     # quick HTTP server
```

---

## Samba (Share Between OSes)

```bash
sudo apt install samba
sudo smbpasswd -a username
sudo service smbd restart
smbclient //HOST/dir -U user
sudo mount -t cifs ...
```

---

Let me know if you'd like this exported as Markdown, PDF, or something else!
