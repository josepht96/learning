# Computer Networking: A Top Down Approach

# chapter 1
Different links can transmit data at different rates, with the transmission rate of a link measured in
bits/second. When one end system has data to send to another end system, the sending end system
segments the data and adds header bytes to each segment. The resulting packages of information,
known as packets in the jargon of computer networks, are then sent through the network to the
destination end system, where they are reassembled into the original data.

The sequence of communication links and packet switches traversed by a packet from the sending end system to the receiving end
system is known as a route or path through the network.

The Internet is all about connecting end
systems to each other, so the ISPs that provide access to end systems must also be interconnected.
These lower-tier ISPs are interconnected through national and international upper-tier ISPs such as
Level 3 Communications, AT&T, Sprint, and NTT.

The IP protocol specifies
the format of the packets that are sent and received among routers and end systems. The Internet’s
principal protocols are collectively known as TCP/IP.

End systems attached to the Internet provide a socket interface that specifies how a program running
on one end system asks the Internet infrastructure to deliver data to a specific destination program
running on another end system. This Internet socket interface is a set of rules that the sending program
must follow so that the Internet can deliver the data to the destination program.

> protocol 
A protocol defines the format and the order of messages exchanged between two or more
communicating entities, as well as the actions taken on the transmission and/or receipt of a message
or other event

> fiber cable
fiber optics connect the cable head end to neighborhood-level junctions, from which traditional
coaxial cable is then used to reach individual houses and apartments. Each neighborhood junction
typically supports 500 to 5,000 homes. Because both fiber and coaxial cable are employed in this
system, it is often referred to as hybrid fiber coax (HFC).

> physical mediums 
Examples of physical media include twisted-pair
copper wire, coaxial cable, multimode fiber-optic cable, terrestrial radio spectrum, and satellite radio
spectrum. Physical media fall into two categories: guided media and unguided media. With guided
media, the waves are guided along a solid medium, such as a fiber-optic cable, a twisted-pair copper
wire, or a coaxial cable. With unguided media, the waves propagate in the atmosphere and in outer
space, such as in a wireless LAN or a digital satellite channel. 
The least expensive and most commonly used guided transmission medium is twisted-pair copper wire.
For over a hundred years it has been used by telephone networks. In fact, more than 99 percent of the
wired connections from the telephone handset to the local telephone switch use twisted-pair copper
wire. 

> packet switching 
Between source and destination, each packet travels through communication links and packet switches (for
which there are two predominant types, routers and link-layer switches). Packets are transmitted over
each communication link at a rate equal to the full transmission rate of the link. So, if a source end
system or a packet switch is sending a packet of L bits over a link with transmission rate R bits/sec, then
the time to transmit the packet is L /R seconds.
Each packet switch has multiple links attached to it. For each attached link, the packet switch has an
output buffer (also called an output queue), which stores packets that the router is about to send into
that link. The output buffers play a key role in packet switching. If an arriving packet needs to be
transmitted onto a link but finds the link busy with the transmission of another packet, the arriving packet
must wait in the output buffer. Thus, in addition to the store-and-forward delays, packets suffer output
buffer queuing delays.

> ip address
When a packet
arrives at a router in the network, the router examines a portion of the packet’s destination address and
forwards the packet to an adjacent router. More specifically, each router has a forwarding table that
maps destination addresses (or portions of the destination addresses) to that router’s outbound links.
When a packet arrives at a router, the router examines the address and searches its forwarding table,
using this destination address, to find the appropriate outbound link. The router then directs the packet
to this outbound link.

> Tier 1 providers
Tier 1 Internet providers are the networks that are the backbone of the Internet. They are sometimes referred to as backbone Internet providers. These ISPs build infrastructure such as the Atlantic Internet sea cables. They provide traffic to all other ISPs, not end users. Tier 1 ISPs own and manage their operating infrastructure, including the routers and other intermediate devices (e.g., switches) that make up the Internet backbone. Key Tier 1 ISPs include AT&T, Verizon, Sprint, NTT, Singtel, PCCW, Telstra, Deutsche Telekom and British Telecom.
A Tier 1 ISP only exchanges Internet traffic with other Tier 1 providers on a non-commercial basis via private settlement-free peering interconnections. They will also interconnect at Internet Exchange Points (IXPs). Tier 1 ISPs can deliver the best network throughput over the Internet backbone through these private peering connections because they own their network infrastructure and have direct control over how traffic flows through these connections.

> Tier 2
A Tier 2 ISP is a service provider that utilizes a combination of paid transit via Tier 1 ISPs and peering with other Tier 2 ISPs to deliver Internet traffic to end customers through Tier 3 ISPs. Tier 2 ISPs are typically regional or national providers. Only a few Tier 2 ISPs can provide service to customers on more than two continents. Often, they will have slower access speeds than Tier 1 ISPs and are at least one router hop away from the backbone of the Internet.

> Tier 3
A Tier 3 ISP is a provider that strictly purchases Internet transit. A Tier 3 provider is by definition primarily engaged in delivering Internet access to end customers. Tier 3 ISPs focus on local business and consumer market conditions. They provide the "on-ramp" or local access to the Internet for end customers, through cable, DSL, fiber or wireless access networks. Their coverage is limited to specific countries or sub regions, such as a metro area. Tier 3 ISPs utilize and pay higher-tier ISPs for access to the rest of the Internet.