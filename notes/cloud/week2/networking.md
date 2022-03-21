# mac address
media access control address, physical address of a piece of hardware. Set when the unit is manufactured. Unique. 
48 bits, half for the vendor code, half for the device

# ipv4
assigned to the device, logical address. Divided into network and host address (think street and house address).
32 bit address, divides in the middle. First 16 bits are the network, second 16 bits are the host address. 
Can use NAT to create more IP addresses

# ipv6
More bits than ipv4, allowing many more addresses. Has a prefix mask that determines how many bits are part of the network, how many
are part of the host address. Usually written in hexadecimal. 
128 bits. 

# ethernet switch
determines routing for devices on its network
At first they send requests to all devices looking for device with proper mac address
on a local network, all devices have the same ip address
make forwarding decisions based on MAC addresses

# router
connects switches to each other. Selects route with the closest match based on subnet mask. 
Make forwarding decisions based on ip addresses

# wireless access points
devices that produce wireless signal
roaming is looking for a closer (stronger signal) device

# copper cabling
twisted pairs of copper wires. Twisting shields the wires from electromagnetic interference. Have limited range ~100m

# fiber optics
use light to represent data, no electromagnetic interference. Can travel much longer distances
single and multi mode, single better because no chance for timing delays.

# OSI model

## physical layer 1
bits, wires and stuff. Data is bits

## data link layer 2
physical addressing. Switches operate at this layer. Data is composed of frames

## network layer 3
IP addresses, routers. Data is packets

## transport layer 4
reliability, confirmation of success of transmission. TCP/UDP. Data is segments

## session layer 5
establishing and monitoring communication sessions
APIs allow communication between applications

## presentation layer 6
how data is represented, encryption protocols. JPG formating, how data is composed into accessible memory

## application layer
How users use the web. HTTP, HTTPS, DNS servers - domain name system

# tcp/ip model
combines layers 5-7. Different names for the layers, achieves the same result

# common protocols
http: tcp port 80
https: tcp port 443
dns server: tcp or udp port 53
ntp: udp port 123
dhcp: udp port 67

# dhcp 
sends out IP address information to all network devices
sends out discover broadcast to everywhere in the subnet, doesnt cross a router
To cross a router you need to configure router as a dhcp relay agent

# dns
converts names into ip addresses. Computer might learn DNS server address via DHCP discovery. 

# nat conversion
translates private network addresses to public addresses. Never make it to internet, are not accessible via internet.
NAT has a Port Address Translation setting, allows multiple internal devices use the same public ip address. 
Dynamically picks a source port number, like 41025 or something uncommon. Router will map a port to a specific device inside the network. 

# ntp
allows network devices to keep their time in sync. Internet based time sources, hyper accurate time. Communicates with UDP over port 123
Servers throughout internet use accurate clocks to serve time values to everyone else. Not quite as accurate but not critical. 

# qos quality of service
managed unfairness. Treat some traffic better than others based on business needs.
Splits data into queues, priority queues and best effort queues. Treat critical traffic better. 

# sdn
software defined networking
commands are sent to SDN controller which sends commands to all desired devices on network.

