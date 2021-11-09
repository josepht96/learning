# Chapter 2

> overview
At the core of network application development is writing programs that run on different end systems
and communicate with each other over the network. For example, in the Web application there are two
distinct programs that communicate with each other: the browser program running in the user’s host
(desktop, laptop, tablet, smartphone, and so on); and the Web server program running in the Web
server host. As another example, in a P2P file-sharing system there is a program in each host that
participates in the file-sharing community. In this case, the programs in the various hosts may be similar
or identical.

Any message sent from one process to another must go
through the underlying network. A process sends messages into, and receives messages from, the
network through a software interface called a socket.
As shown in this figure, a socket is the interface between the application layer and the
transport layer within a host. It is also referred to as the Application Programming Interface (API)
between the application and the network, since the socket is the programming interface with which
network applications are built. The application developer has control of everything on the applicationlayer side of the socket but has little control of the transport-layer side of the socket. The only control
that the application developer has on the transport-layer side is (1) the choice of transport protocol and
(2) perhaps the ability to fix a few transport-layer parameters such as maximum buffer and maximum.

Thus, to support these applications, something has to
be done to guarantee that the data sent by one end of the application is delivered correctly and
completely to the other end of the application. If a protocol provides such a guaranteed data delivery
service, it is said to provide reliable data transfer.

# security 
Finally, a transport protocol can provide an application with one or more security services. For example,
in the sending host, a transport protocol can encrypt all data transmitted by the sending process, and in
the receiving host, the transport-layer protocol can decrypt the data before delivering the data to the
receiving process. Such a service would provide confidentiality between the two processes, even if the
data is somehow observed between sending and receiving processes. A transport protocol can also
provide other security services in addition to confidentiality, including data integrity and end-point
authentication, topics that we’ll cover in detail in Chapter 8.
> ssl
Neither TCP nor UDP provides any encryption—the data that the sending process passes into
its socket is the same data that travels over the network to the destination process. So, for
example, if the sending process sends a password in cleartext (i.e., unencrypted) into its socket,
the cleartext password will travel over all the links between sender and receiver, potentially
getting sniffed and discovered at any of the intervening links. Because privacy and other security
issues have become critical for many applications, the Internet community has developed an
enhancement for TCP, called Secure Sockets Layer (SSL). TCP-enhanced-with-SSL not only
does everything that traditional TCP does but also provides critical process-to-process security
services, including encryption, data integrity, and end-point authentication. We emphasize that
SSL is not a third Internet transport protocol, on the same level as TCP and UDP, but instead is
an enhancement of TCP, with the enhancements being implemented in the application layer. In
particular, if an application wants to use the services of SSL, it needs to include SSL code
(existing, highly optimized libraries and classes) in both the client and server sides of the
application. SSL has its own socket API that is similar to the traditional TCP socket API. When
an application uses SSL, the sending process passes cleartext data to the SSL socket; SSL in
the sending host then encrypts the data and passes the encrypted data to the TCP socket. The
encrypted data travels over the Internet to the TCP socket in the receiving process. The
receiving socket passes the encrypted data to SSL, which decrypts the data. Finally, SSL
passes the cleartext data through its SSL socket to the receiving process. We’ll cover SSL in
some detail in Chapter 8

# application layer protocols
The Web’s application-layer protocol, HTTP,
defines the format and sequence of messages exchanged between browser and Web server. Thus,
HTTP is only one piece (albeit, an important piece) of the Web application. As another example, an
Internet e-mail application also has many components, including mail servers that house user
mailboxes; mail clients (such as Microsoft Outlook) that allow users to read and create messages; a
standard for defining the structure of an e-mail message; and application-layer protocols that define how
messages are passed between servers, how messages are passed between servers and mail clients,
and how the contents of message headers are to be interpreted. The principal application-layer protocol
for electronic mail is SMTP (Simple Mail Transfer Protocol) [RFC 5321]. Thus, e-mail’s principal
application-layer protocol, SMTP, is only one piece (albeit an important piece) of the e-mail application.
