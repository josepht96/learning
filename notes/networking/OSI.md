# OSI model

OSI model can be loglcailly separated into two parts: layers 1-4 and 5-7.
Layers 1-4 handle the transmission of data from NIC to NIC, while layers
5-7 handle actually using the data.

## Physical layer

The physical layer (layer 1) defines how data is transfered over a physical medium. 1s and 0s traveling
over a wire. Defines the format of this process - likely not concerned with what data
is actually being transfered, just the structure

## Data link layer

The data link layer (layer 2) defines how data enters and leaves layer 1. When data leaves
the physical medium, it is converted into `frames`.
This is handled by a NIC, network interface card. A NIC has a MAC address - which
is unique to it (think SSN for US people).
"The Data Link layer of the OSI model is responsible for interfacing with the Physical layer. Effectively, Layer 2 is responsible for putting 1’s and 0’s on the wire, and pulling 1’s and 0’s from the wire."

The overarching function of the Data Link layer is to deliver packets from one NIC to another. Or to put it another way, the role of Layer 2 is to deliver packets from hop to hop.

## Network layer

The network layer (layer 3 ) handles end to end communicatio between networks. It does this by using another addressing scheme that can **logically** identify every node connected to the Internet. This addressing scheme is known as the Internet Protocol address, or the IP Address.

Routers are Network Devices that operate at Layer 3 of the OSI model. A Router’s primary responsibility is to facilitate communication between Networks. As such, a Router creates a boundary between two networks. In order to communicate with any device not directly in your network, a router must be used.
Routers handle data frames that have IP addr, IP dest, and MAC addresses. Routers handle sending
the data frame to the next device (using MAC address) that will help the data reach its ultimate
destnation, IP dest. Thus, IP addresses define source and dest, MAC addresses handle individual hops
along the way.

### Why do we need MAC addresses and IP addresses

Layer 2 uses MAC addresses and is responsible for packet delivery from hop to hop.
Layer 3 uses IP addresses and is responsible for packet delivery from end to end.
When a computer has data to send, it encapsulates it in a IP header which will include information like the Source and Destination IP addresses of the two “ends” of the communication.

The IP Header and Data are then further encapsulated in a MAC address header, which will include information like the Source and Destination MAC address of the current “hop” in the path towards the final destination.

## Transport layer

The transport layer (layer 4) handles data after it has arrived at a NIC. If a NIC is handling data
many applications running on the host, it needs a way to say 'this frame is with Spotify, this next
frame is for Youtube.' Layer 4 is responsible for 'service to service' delivery. Services can be differentiated by using separate ports. 1.2.3.4:3000 is for 1 thing 1.2.3.4:3001 is for something else.
TCP header adds the necessary port to the data frame.

## Summary

Layer 4 will add a TCP header which would include a Source and Destination port
Layer 3 will add an IP header which would include a Source and Destination IP address
Layer 2 would add an Ethernet header which would include a Source and Destination MAC address

The process of preparing data to send across a wire is called encapsulation, and the process
of convert 1s and 0s back to usable data is called decapsulation.

## Session layer

The session layer (layer 5) is responsible for maintaining session data between two machines.
When two computers or other networked devices need to speak with one another, a session needs to be created, and this is done at the Session Layer. Functions at this layer involve setup, coordination (how long should a system wait for a response, for example) and termination between the applications at each end of the session.

## Presenation layer

The presenation layer, (layer 6) handles the translation from network data to application data.
It is independent of data representation at the application layer. In general, it represents the preparation or translation of application format to network format, or from network formatting to application format. In other words, the layer “presents” data for the application or the network. A good example of this is encryption and decryption of data for secure transmission; this happens at Layer 6.

## Application layer

The application layer (layer 7) handles data communication between the application layer and the display system (web browser). How does data actually get from a NIC to a web browser?
The Application Layer in the OSI model is the layer that is the “closest to the end user”. It receives information directly from users and displays incoming data to the user. Oddly enough, applications themselves do not reside at the application layer. Instead the layer facilitates communication through lower layers in order to establish connections with applications at the other end. Web browsers (Google Chrome, Firefox, Safari, etc.) TelNet, and FTP, are examples of communications  that rely  on Layer 7.
