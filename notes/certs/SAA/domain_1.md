## Domain 1: Design Secure Architectures (30%)

### Task Statement 1.1: Design secure access to AWS resources

**Knowledge of:**

- Access controls and management across multiple accounts
- AWS federated access and identity services (IAM, IAM Identity Center)
- AWS global infrastructure (Availability Zones, Regions)
- AWS security best practices (principle of least privilege)
- The AWS shared responsibility model

**Skills in:**

- Applying AWS security best practices to IAM users and root users (MFA)
- Designing a flexible authorization model that includes IAM users, groups, roles, and policies
- Designing a role-based access control strategy (AWS STS, role switching, cross-account access)
- Designing a security strategy for multiple AWS accounts (AWS Control Tower, SCPs)
- Determining the appropriate use of resource policies for AWS services

**Topic Overview:**

**IAM Best Practices:**

- Enable MFA for root and privileged users
- Use IAM roles instead of long-term credentials
- Implement least privilege access
- Use IAM Access Analyzer to identify overly permissive policies
- Rotate credentials regularly
- Use IAM policy conditions for fine-grained control

**Multi-Account Strategy:**

- **AWS Organizations:** Centrally manage multiple AWS accounts
- **Service Control Policies (SCPs):** Set permission guardrails across accounts
- **AWS Control Tower:** Automated multi-account setup with governance
- **AWS Resource Access Manager:** Share resources across accounts
- **Cross-account roles:** Allow access between accounts without sharing credentials

**Federated Access:**

- **SAML 2.0:** Enterprise identity federation
- **Web Identity Federation:** Mobile and web app authentication
- **AWS IAM Identity Center (SSO):** Centralized access management
- **AWS STS:** Temporary security credentials for federated users

**Resource Policies:**

- S3 bucket policies
- KMS key policies
- SQS queue policies
- Lambda resource policies
- When to use identity-based vs. resource-based policies

---

### Task Statement 1.2: Design secure workloads and applications

**Knowledge of:**

- Application configuration and credentials security
- AWS service endpoints
- Control ports, protocols, and network traffic on AWS
- Secure application access
- Security services with appropriate use cases (Amazon Cognito, GuardDuty, Macie)
- Threat vectors external to AWS (DDoS, SQL injection)

**Skills in:**

- Designing VPC architectures with security components (security groups, route tables, NACLs, NAT gateways)
- Determining network segmentation strategies (public/private subnets)
- Integrating AWS services to secure applications (AWS Shield, WAF, IAM Identity Center, Secrets Manager)
- Securing external network connections to and from AWS Cloud (VPN, Direct Connect)

**Topic Overview:**

**VPC Security Architecture:**

**Network Segmentation:**

- **Public subnets:** Resources that need internet access (load balancers, NAT gateways)
- **Private subnets:** Application servers, databases (no direct internet access)
- **Isolated subnets:** Highly sensitive workloads
- Multi-tier architecture (web, app, database layers)

**Security Groups:**

- Stateful firewall at instance level
- Allow rules only (implicit deny)
- Support for security group references
- Best practice: Specific rules, avoid 0.0.0.0/0 for inbound

**Network ACLs:**

- Stateless firewall at subnet level
- Allow and deny rules
- Numbered rules processed in order
- Use for additional layer of defense

**VPC Endpoints:**

- **Gateway endpoints:** S3 and DynamoDB (free)
- **Interface endpoints:** Other AWS services via PrivateLink
- Keep traffic within AWS network, improve security

**Application Security:**

**AWS WAF (Web Application Firewall):**

- Protect against common web exploits (SQL injection, XSS)
- Rate limiting and geo-blocking
- Managed rules and custom rules
- Integrate with CloudFront, ALB, API Gateway

**AWS Shield:**

- **Standard:** Automatic DDoS protection (free)
- **Advanced:** Enhanced DDoS protection, 24/7 DDoS Response Team, cost protection

**AWS Secrets Manager:**

- Store and rotate database credentials, API keys
- Automatic rotation for RDS, Redshift, DocumentDB
- Encryption with KMS
- Alternative: AWS Systems Manager Parameter Store (simpler, less features)

**Amazon Cognito:**

- User authentication and authorization for web/mobile apps
- User pools (user directory) and identity pools (AWS credentials)
- Support for social identity providers (Google, Facebook)
- SAML and OIDC integration

**Threat Detection:**

**Amazon GuardDuty:**

- Intelligent threat detection
- Analyzes CloudTrail events, VPC Flow Logs, DNS logs
- Machine learning-based anomaly detection
- Findings for compromised instances, reconnaissance

**Amazon Macie:**

- Discover and protect sensitive data in S3
- PII, financial data, credentials detection
- Data classification and security posture

**AWS Security Hub:**

- Centralized security and compliance view
- Aggregates findings from GuardDuty, Inspector, Macie, etc.
- Compliance checks (CIS, PCI-DSS)

**Hybrid Connectivity:**

**AWS VPN:**

- **Site-to-Site VPN:** Connect on-premises to VPC
- **Client VPN:** Remote user access
- IPsec encryption
- Quick to set up, uses internet

**AWS Direct Connect:**

- Dedicated network connection (1 Gbps to 100 Gbps)
- More consistent performance than VPN
- Private connectivity, doesn't use internet
- Can combine with VPN for encrypted connection
- Use Direct Connect Gateway for multi-region access

---

### Task Statement 1.3: Determine appropriate data security controls

**Knowledge of:**

- Data access and governance
- Data recovery
- Data retention and classification
- Encryption and appropriate key management

**Skills in:**

- Aligning AWS technologies to meet compliance requirements
- Encrypting data at rest (AWS KMS)
- Encrypting data in transit (AWS Certificate Manager using TLS)
- Implementing access policies for encryption keys
- Implementing data backups and replications
- Implementing policies for data access, lifecycle, and protection
- Rotating encryption keys and renewing certificates

**Topic Overview:**

**Encryption at Rest:**

**AWS KMS (Key Management Service):**

- Create and manage encryption keys
- Customer managed keys (CMKs) for full control
- AWS managed keys (automatic rotation)
- CloudHSM for dedicated hardware security module
- Key policies control access to keys
- Automatic integration with S3, EBS, RDS, etc.

**Service-Specific Encryption:**

- **S3:** SSE-S3, SSE-KMS, SSE-C (customer-provided keys), client-side encryption
- **EBS:** Encryption checkbox when creating volumes, encrypted snapshots
- **RDS:** Encryption option at creation, includes backups and read replicas
- **DynamoDB:** Encryption by default using AWS owned keys or KMS

**Encryption in Transit:**

**AWS Certificate Manager (ACM):**

- Provision, manage, and deploy SSL/TLS certificates
- Free public certificates
- Automatic renewal
- Integration with ELB, CloudFront, API Gateway
- Cannot export public certificates (private CA certificates can be exported)

**TLS/SSL Implementation:**

- Use HTTPS for data in transit
- TLS 1.2 or higher
- Secure listeners on load balancers
- API Gateway with custom domains

**Key Management Best Practices:**

- Use separate keys for different data classifications
- Enable automatic key rotation (yearly for CMKs)
- Use key policies and IAM policies together
- Enable CloudTrail to log key usage
- Use grants for temporary, granular permissions

**Data Backup and Recovery:**

**Backup Strategies:**

- **EBS snapshots:** Point-in-time backups, incremental
- **RDS automated backups:** Daily snapshots, transaction logs, 1-35 day retention
- **RDS manual snapshots:** Keep until explicitly deleted
- **DynamoDB backups:** On-demand and point-in-time recovery (PITR)
- **AWS Backup:** Centralized backup management across services

**Cross-Region Replication:**

- S3 cross-region replication (CRR) for compliance and latency
- RDS read replicas in different regions for DR
- DynamoDB global tables for multi-region replication

**Data Lifecycle and Retention:**

**S3 Lifecycle Policies:**

- Transition objects between storage classes
- Expire objects after specified time
- Delete incomplete multipart uploads
- Version-aware lifecycle policies

**S3 Object Lock:**

- WORM (Write Once Read Many) model
- Compliance mode (cannot be deleted even by root)
- Governance mode (special permissions can override)
- Legal hold

**Data Classification:**

- Use tags for sensitivity levels (public, internal, confidential, restricted)
- Amazon Macie for automated discovery and classification
- Implement access controls based on classification

---
