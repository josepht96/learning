# Domain 2: Security and Compliance (30%)

## Task Statement 2.1: Understand the AWS shared responsibility model

**Knowledge of:**

- AWS shared responsibility model

**Skills in:**

- Recognizing components of the shared responsibility model
- Describing customer responsibilities on AWS
- Describing AWS responsibilities
- Understanding how responsibilities shift based on service type

**The Model:**

**AWS Responsibility (Security OF the Cloud):**

- Physical security of data centers
- Hardware and network infrastructure
- Hypervisor and virtualization layer
- Managed service operations

**Customer Responsibility (Security IN the Cloud):**

- Customer data
- Platform, applications, identity and access management
- Operating system, network, and firewall configuration
- Client-side and server-side encryption
- Network traffic protection

**How Responsibilities Vary:**

- **IaaS (EC2):** Customer manages more (OS, patches, security groups)
- **PaaS (RDS):** AWS manages more (OS, patches), customer manages data and access
- **SaaS (Managed services):** AWS manages most, customer manages data and access control

> **EXAM FOCUS:** The shared responsibility model is one of the MOST tested concepts on the CCP exam. Remember: **AWS = security OF the cloud (hardware, infrastructure), Customer = security IN the cloud (data, access, encryption)**.

---

## Task Statement 2.2: Understand AWS Cloud security, governance, and compliance concepts

**Knowledge of:**

- AWS compliance and governance concepts
- Benefits of cloud security (encryption)
- Where to capture and locate logs associated with cloud security

**Skills in:**

- Understanding compliance programs AWS participates in
- Identifying governance and compliance tools
- Recognizing encryption benefits and capabilities

**Key Topics:**

**Compliance Programs:**

- HIPAA (healthcare)
- PCI DSS (payment card industry)
- SOC 1/2/3
- ISO certifications
- GDPR readiness
- FedRAMP (US government)

**Governance Tools:**

- AWS Organizations (multi-account management)
- AWS Control Tower (account setup automation)
- Service Control Policies (SCPs)
- AWS Config (configuration compliance)

**Encryption:**

- Data at rest encryption (S3, EBS, RDS)
- Data in transit encryption (TLS/SSL)
- AWS Key Management Service (KMS)

**Logging and Monitoring:**

- AWS CloudTrail (API activity logging)
- Amazon CloudWatch (metrics and logs)
- VPC Flow Logs (network traffic)
- S3 access logs

> **COMMON USAGE:** **CloudTrail** and **CloudWatch** are universally enabled in production environments. CloudTrail is often the first service enabled for security auditing and compliance. Most organizations centralize logs in a dedicated security account.

---

## Task Statement 2.3: Identify AWS access management capabilities

**Knowledge of:**

- Identity and access management
- Authentication and authorization concepts

**Skills in:**

- Understanding IAM components
- Implementing multi-factor authentication
- Managing users, groups, roles, and policies

**Key Components:**

**IAM Users:**

- Individual identities for people or applications
- Long-term credentials
- Should enable MFA

**IAM Groups:**

- Collection of users
- Simplifies permission management

**IAM Roles:**

- Temporary security credentials
- Used by services, applications, or federated users
- No permanent credentials

> **BEST PRACTICE:** IAM Roles are strongly preferred over IAM Users in modern AWS architectures. Use roles for EC2 instances, Lambda functions, and cross-account access. Only create IAM users when absolutely necessary (e.g., third-party integrations without role support).

**IAM Policies:**

- JSON documents defining permissions
- Attached to users, groups, or roles
- Follow principle of least privilege

**Access Methods:**

- AWS Management Console
- AWS CLI (Command Line Interface)
- AWS SDKs
- Federated access (SAML, web identity)

---

## Task Statement 2.4: Identify components and resources for security

**Knowledge of:**

- AWS security services and tools
- Security best practices

**Skills in:**

- Identifying appropriate security services for use cases
- Understanding security resources available

**Security Services:**

**Threat Detection:**

- Amazon GuardDuty (intelligent threat detection)
- AWS Security Hub (centralized security view)
- Amazon Inspector (vulnerability scanning)

**DDoS Protection:**

- AWS Shield Standard (automatic protection)
- AWS Shield Advanced (enhanced protection)

**Web Application Security:**

- AWS WAF (Web Application Firewall)
- AWS Firewall Manager (centralized firewall management)

**Data Protection:**

- AWS KMS (Key Management Service)
- AWS Secrets Manager (credentials management)
- AWS Certificate Manager (SSL/TLS certificates)

**Compliance and Auditing:**

- AWS Artifact (compliance reports)
- AWS Audit Manager (continuous auditing)

**Network Security:**

- Security Groups (instance-level firewalls)
- Network ACLs (subnet-level firewalls)
- AWS Network Firewall

> **COMMON USAGE:** **Security Groups** are used far more frequently than Network ACLs. Security Groups are stateful and easier to manage, making them the primary network security control. NACLs are typically only used for additional subnet-level deny rules or compliance requirements.
