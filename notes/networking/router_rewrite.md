Most home routers use a special-case of NAT called PAT.

You'll also see it referred to as NAPT, or IP Masquerading. All three of the latter terms mean the same thing in general use. (The acronyms - Network Address Translation / Port Address Translation / Network Address Port Translation)

When the packet goes out from your internal machine, the source address is rewritten as you are aware. The source port is also changed, usually to a high number, and the router keeps an address translation table.

For example, let's say you have a client machine that goes to www.google.com. Your computer (e.g., 192.168.1.100) looks that address up and makes a TCP connection to 72.14.204.147 on port 80 from your internal IP address, using a random source port.

To your computer, the connection looks like this:

192.168.1.100:37641   <-->  72.14.204.147:80
Your computer sends the packet to the router, which picks a new random high port and rewrites the packet. Each outbound connection gets its own port on the router. The router then forwards the packet on to your ISP after adding it to its connection table:

PrivateIP        PrivatePort   PublicIP      PublicPort    Remote          RemotePort
-------------    ----------    -----------   -----------   ----------      -----------
192.168.1.100    37641         *10.6.23.5    59273         72.14.204.147   80
*For example purposes, I used an address starting with 10, but these aren't publicly routable. The table is also somewhat oversimplified.

To google, the connection looks like this:

10.6.23.5:59273   <-->  72.14.204.147:80
Google will send it's reponse to 10.6.23.5 on port 59273. Your router then looks up that information in the table and forwards the packet on to 192.168.1.100:37641.