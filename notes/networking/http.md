# http
> 131
The HyperText Transfer Protocol (HTTP), the Web’s application-layer protocol, is at the heart of the
Web. It is defined in [RFC 1945] and [RFC 2616]. HTTP is implemented in two programs: a client
program and a server program. The client program and server program, executing on different end
systems, talk to each other by exchanging HTTP messages. HTTP defines the structure of these
messages and how the client and server exchange the messages. 

HTTP uses TCP as its underlying transport protocol (rather than running on top of UDP). The HTTP
client first initiates a TCP connection with the server. Once the connection is established, the browser
and the server processes access TCP through their socket interfaces. As described in Section 2.1, on
the client side the socket interface is the door between the client process and the TCP connection; on
the server side it is the door between the server process and the TCP connection. The client sends
HTTP request messages into its socket interface and receives HTTP response messages from its
socket interface. Similarly, the HTTP server receives request messages