Network rules
Azure firewall blocks traffic by default

# NAT Rules
For external connections to azure resources, like ssh and rdp. Non-http traffic
# Network Rules
Similar to above, but allows traffic between resources within the virtual network. SQL database to web server or something.
# Application Rules
Handles ingress traffic at layer 7. How public internet can access resources within firewall

Deny inbound traffic from internet
allow outbound traffic to internet

Standard public IPs are zone redundant, basic are not

Application rule for Windows update

what are virtual appliances
azure firewall vs azure web application firewall

Application gateway provides a WAF for inbound connections only for HTTP/S traffic (OWASP rules and more), Azure Firewall provides both inbound and outbound filtering also for non-HTTP traffic (E.G. your VMs can only go out to FQDN X, Y on port Z, K. and block other traffic)


review the differences between azure backup, site recovery, managed snapchots, recovery services vault
snapshots for database disks
determine which ones support linux and non azure vms
notes