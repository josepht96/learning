# AWS VPC Components

## VPC (Virtual Private Cloud)
The isolated network environment in AWS where your resources live. Free to create. All resources within a VPC can communicate with each other using private IPs.

---

## Subnets
Subdivisions of your VPC CIDR block. There is no native "public" or "private" subnet type in AWS — the distinction is purely a naming convention based on the associated route table.

- **Public subnet** — route table has a `0.0.0.0/0` route pointing to the IGW
- **Private subnet** — route table has no direct route to the IGW

Each subnet is associated with exactly one route table. If no explicit association exists, the subnet falls back to the VPC's main route table.

---

## Internet Gateway (IGW)
Attached at the VPC level (not inside any subnet). Enables traffic to flow between the VPC and the public internet. Performs NAT — translating private IPs to public IPs on outbound traffic, and the reverse on inbound.

- One IGW per VPC
- Free

---

## NAT Gateway
Allows resources in a private subnet to initiate outbound internet traffic without being directly reachable from the internet. Lives in a public subnet and requires an EIP. All outbound traffic from private subnet resources appears to originate from the NAT Gateway's EIP.

- Requires an EIP
- ~$0.045/hr (~$32/month) + $0.045/GB data processing
- Outbound only — internet cannot initiate connections to private resources through it

---

## Elastic IP (EIP)
A static public IPv4 address owned by your AWS account. Unlike auto-assigned public IPs (which change on EC2 stop/start), EIPs persist until explicitly released.

- Free when attached to a running resource
- ~$0.005/hr (~$3.60/month) when unattached
- Required for NAT Gateway

---

## Route Tables
Define where network traffic is directed. Each subnet is associated with one route table. AWS uses **longest prefix match** — the most specific route wins.

**Public route table example:**
| Destination | Target |
|---|---|
| 10.0.0.0/16 | local |
| 0.0.0.0/0 | IGW |

**Private route table example:**
| Destination | Target |
|---|---|
| 10.0.0.0/16 | local |
| 0.0.0.0/0 | NAT Gateway |

The **main route table** is created automatically with the VPC. Any subnet without an explicit association falls back to it. Best practice is to keep it restrictive (local only) and use explicit route tables for all subnets.

---

## Application Load Balancer (ALB)
Sits in the public subnet and accepts inbound traffic from the internet. Forwards requests to private EC2 instances via their private IPs. The ALB does not have a static public IP — it exposes a DNS name that resolves to AWS-managed IPs.

- Use a CNAME to point your domain to the ALB DNS name, never hardcode the IP
- For a static public IP in front of a load balancer, use an NLB with an EIP instead

---

## Security Groups
Stateful firewalls attached at the resource level (EC2, RDS, ALB, etc.). Return traffic for allowed connections is automatically permitted. Free.

---

## VPC Endpoints
Allow private subnet resources to reach AWS services without going through the internet or NAT Gateway.

| Type | Cost | Services |
|---|---|---|
| Gateway Endpoint | Free | S3, DynamoDB |
| Interface Endpoint | ~$0.01/hr per AZ | ECR, SSM, Secrets Manager, etc. |

Using a Gateway Endpoint for S3 eliminates S3 data transfer costs through the NAT Gateway.

---

## Typical Architecture

```
Internet
    │
    ▼
Internet Gateway (VPC boundary)
    │
    ▼
Public Subnet
├── Application Load Balancer   ← internet-facing entry point
└── NAT Gateway + EIP           ← outbound path for private subnet
    │
    ▼
Private Subnet
├── EC2 (Service A)
├── EC2 (Service B)
└── RDS / ElastiCache
```

---

## Cost Summary

| Component | Cost |
|---|---|
| VPC | Free |
| Subnets | Free |
| IGW | Free |
| Route Tables | Free |
| Security Groups | Free |
| Gateway VPC Endpoints | Free |
| NAT Gateway | ~$32/month + $0.045/GB |
| EIP (unattached) | ~$3.60/month |
| Interface VPC Endpoints | ~$7/month per AZ |