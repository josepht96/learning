# AWS Certified Solutions Architect Associate (SAA-C03) - Exam Outline

## Exam Overview

- **Exam Code:** SAA-C03
- **Duration:** 130 minutes
- **Question Format:** 65 questions
- **Question Types:** Multiple choice and multiple response
- **Passing Score:** 720 out of 1000 (scaled score)
- **Cost:** $150 USD
- **Validity:** 3 years

## Exam Purpose

This exam validates the ability to design solutions based on the AWS Well-Architected Framework. It's intended for individuals who perform in a solutions architect role and have one or more years of hands-on experience designing available, cost-efficient, fault-tolerant, and scalable distributed systems on AWS.

## Target Candidate Profile

- One or more years of hands-on experience designing cloud solutions
- Experience with compute, networking, storage, and database AWS services
- Ability to identify and define technical requirements for AWS-based applications
- Experience deploying hybrid systems with on-premises and AWS components
- Understanding of the AWS global infrastructure, security services, and the AWS shared responsibility model

---

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

## Domain 2: Design Resilient Architectures (26%)

### Task Statement 2.1: Design scalable and loosely coupled architectures

**Knowledge of:**
- API creation and management (Amazon API Gateway)
- AWS managed services with appropriate use cases
- Caching strategies
- Design principles for microservices
- Event-driven architectures
- Horizontal scaling and vertical scaling
- Stateless vs. stateful applications
- Appropriate use of edge accelerators (CloudFront)
- Load balancing concepts
- Multi-tier architectures
- Queuing and messaging

**Skills in:**
- Designing event-driven, microservices, and multi-tier architectures
- Determining scaling strategies for components
- Determining AWS services for loose coupling
- Determining when to use read replicas
- Determining when to use containers vs. serverless

**Topic Overview:**

**Loose Coupling Strategies:**

**Messaging Services:**
- **Amazon SQS (Simple Queue Service):**
  - Decouple components with message queues
  - Standard queues: Best-effort ordering, at-least-once delivery, unlimited throughput
  - FIFO queues: Exactly-once processing, order guaranteed, 3,000 msg/sec with batching
  - Visibility timeout, dead-letter queues, long polling

- **Amazon SNS (Simple Notification Service):**
  - Pub/sub messaging
  - Fan-out pattern (one message to multiple subscribers)
  - Push-based delivery to SQS, Lambda, HTTP, email, SMS
  - Message filtering

- **Amazon EventBridge:**
  - Serverless event bus
  - Schema registry
  - Integration with AWS services and SaaS applications
  - Event rules and targets

**Load Balancing:**

**Application Load Balancer (ALB):**
- Layer 7 (HTTP/HTTPS)
- Path-based and host-based routing
- Support for containers and microservices
- WebSocket support
- User authentication (Cognito, OIDC)

**Network Load Balancer (NLB):**
- Layer 4 (TCP/UDP)
- Extreme performance (millions of requests/sec)
- Static IP addresses
- Preserve source IP address
- Use for: real-time streaming, gaming, IoT

**Gateway Load Balancer:**
- Deploy third-party virtual appliances
- Traffic inspection, firewalls, intrusion detection

**Scaling Strategies:**

**Horizontal Scaling (Scale Out/In):**
- Add or remove instances
- Auto Scaling groups
- Stateless applications scale easily
- Use load balancers to distribute traffic

**Vertical Scaling (Scale Up/Down):**
- Change instance size
- Requires downtime for EC2
- Simpler but limited by instance size

**Auto Scaling:**
- **Target tracking:** Maintain metric at target value (CPU, network, custom)
- **Step scaling:** Scale based on CloudWatch alarms
- **Scheduled scaling:** Predictable load patterns
- **Predictive scaling:** Machine learning-based forecasting
- Warm-up and cooldown periods

**Caching Strategies:**

**Application Caching:**
- **Amazon ElastiCache:**
  - Redis: Advanced data structures, pub/sub, persistence, Multi-AZ
  - Memcached: Simple, multi-threaded, horizontal scaling
- Common patterns: Lazy loading, write-through, session storage

**Content Caching:**
- **Amazon CloudFront:**
  - Global CDN with edge locations
  - Origin: S3, ALB, EC2, custom origins
  - Cache based on headers, cookies, query strings
  - Signed URLs and signed cookies for private content
  - Lambda@Edge for edge computing

**Database Caching:**
- DynamoDB Accelerator (DAX): Microsecond latency
- RDS read replicas: Offload read traffic

**Microservices Architecture:**

**Containers:**
- **Amazon ECS:**
  - AWS-native container orchestration
  - Fargate (serverless) or EC2 launch types
  - Task definitions, services, clusters
  - Integration with ALB for service discovery

- **Amazon EKS:**
  - Managed Kubernetes
  - Use when you need Kubernetes features or multi-cloud portability
  - Fargate or EC2 node groups

**Serverless:**
- **AWS Lambda:**
  - Event-driven compute
  - Automatic scaling
  - Pay per invocation
  - 15-minute max execution time
  - Best for: Short-lived, stateless functions

**API Management:**
- **Amazon API Gateway:**
  - RESTful and WebSocket APIs
  - Request validation and transformation
  - Throttling and quota management
  - Caching
  - Integration with Lambda, HTTP endpoints, AWS services

**Stateless Application Design:**
- Store session data externally (ElastiCache, DynamoDB)
- Use signed cookies or JWT tokens
- Enables horizontal scaling and fault tolerance

---

### Task Statement 2.2: Design highly available and/or fault-tolerant architectures

**Knowledge of:**
- AWS global infrastructure (Availability Zones, Regions)
- AWS managed services with appropriate use cases
- Basic networking concepts (route tables, VPC)
- Disaster recovery (DR) strategies
- Distributed design patterns
- Failover strategies
- Immutable infrastructure
- Load balancing concepts
- Proxy server concepts
- Service quotas and throttling
- Storage options and characteristics

**Skills in:**
- Determining automation techniques to ensure infrastructure integrity
- Determining AWS services for high availability and fault tolerance
- Identifying metrics based on business requirements to deliver highly available solution
- Implementing designs to mitigate single points of failure
- Implementing strategies to ensure data durability and availability

**Topic Overview:**

**High Availability Design:**

**Multi-AZ Deployments:**
- Deploy resources across multiple Availability Zones
- AZs are physically separate with independent power/networking
- **RDS Multi-AZ:** Synchronous replication, automatic failover (~1-2 min)
- **Aurora:** 6 copies across 3 AZs, automatic failover
- **ELB:** Automatically distributes across AZs
- **S3:** Automatically stores across multiple AZs (except One Zone-IA)

**Multi-Region Deployments:**
- For disaster recovery and global applications
- **Route 53 routing policies:**
  - Failover: Active-passive failover
  - Geolocation: Route based on user location
  - Geoproximity: Route based on geography with bias
  - Latency: Route to lowest latency endpoint
  - Weighted: Distribute traffic across multiple resources
  - Multi-value answer: Return multiple healthy endpoints
- Cross-region read replicas for databases
- S3 cross-region replication
- CloudFront for global content delivery

**Disaster Recovery Strategies:**

**RTO and RPO:**
- **RTO (Recovery Time Objective):** How long to recover
- **RPO (Recovery Point Objective):** How much data loss is acceptable

**DR Strategies (from fastest to slowest):**

1. **Multi-Site Active-Active:**
   - Run full production in multiple regions
   - Route 53 routes traffic to both
   - Lowest RTO/RPO, highest cost

2. **Warm Standby:**
   - Scaled-down version running in another region
   - Scale up when needed
   - Medium RTO/RPO, medium cost

3. **Pilot Light:**
   - Core components running (databases), other resources off
   - Start resources when needed
   - Higher RTO/RPO, lower cost

4. **Backup and Restore:**
   - Regular backups to S3
   - Restore when needed
   - Highest RTO/RPO, lowest cost

**Data Durability and Availability:**

**S3 Durability and Availability:**
- Durability: 99.999999999% (11 9's) for most storage classes
- Availability: 99.9% to 99.99% depending on storage class
- Versioning: Protect against accidental deletion
- MFA Delete: Require MFA to delete versions
- Cross-region replication for additional protection

**EBS Durability:**
- Annual failure rate (AFR) 0.1% - 0.2%
- Take regular snapshots to S3 (11 9's durability)
- Copy snapshots across regions for DR
- Can create volumes from snapshots in any AZ within region

**Database Durability:**
- RDS automated backups to S3
- Read replicas for data redundancy
- DynamoDB point-in-time recovery
- Aurora backtrack (rewind without restore)

**Eliminating Single Points of Failure:**
- Use multiple AZs
- Implement load balancing
- Auto Scaling for automatic replacement
- Use managed services (built-in redundancy)
- Regular health checks
- Automated failover

**Immutable Infrastructure:**
- Don't update servers, replace them
- Use Auto Scaling and launch templates
- Version all infrastructure changes
- Golden AMIs or container images
- Reduces configuration drift

**Service Quotas and Throttling:**
- **Service Quotas:** Limits per account per region
- Request quota increases when needed
- Design for throttling with retries and exponential backoff
- Use Service Quotas service to monitor and manage

---

## Domain 3: Design High-Performing Architectures (24%)

### Task Statement 3.1: Determine high-performing and/or scalable storage solutions

**Knowledge of:**
- Hybrid storage solutions
- Storage services with appropriate use cases
- Storage types with associated characteristics

**Skills in:**
- Determining storage services and configurations for performance
- Determining storage services for data migration

**Topic Overview:**

**S3 Performance Optimization:**

**Performance Characteristics:**
- 3,500 PUT/COPY/POST/DELETE requests per second per prefix
- 5,500 GET/HEAD requests per second per prefix
- Use multiple prefixes for higher throughput
- Multipart upload for files > 100 MB (required for > 5 GB)
- S3 Transfer Acceleration for long-distance uploads (uses CloudFront edge locations)
- Byte-range fetches for faster downloads and partial retrieval

**S3 Storage Classes (Performance Perspective):**
- **S3 Standard:** Millisecond latency, frequent access
- **S3 Intelligent-Tiering:** Automatic cost optimization, no retrieval fees
- **S3 Standard-IA:** Millisecond latency, infrequent access, retrieval fee
- **S3 One Zone-IA:** Lower cost, single AZ, infrequent access
- **S3 Glacier Instant Retrieval:** Millisecond retrieval, quarterly access
- **S3 Glacier Flexible Retrieval:** Minutes to hours (expedited: 1-5 min, standard: 3-5 hours, bulk: 5-12 hours)
- **S3 Glacier Deep Archive:** Lowest cost, 12-48 hours retrieval

**EBS Performance:**

**EBS Volume Types:**

**SSD (Solid State Drives):**
- **gp3 (General Purpose SSD):**
  - 3,000-16,000 IOPS, 125-1,000 MB/s throughput
  - IOPS and throughput configurable independently
  - Most cost-effective SSD

- **gp2 (General Purpose SSD - Previous Gen):**
  - IOPS scale with volume size (3 IOPS/GB)
  - Burstable to 3,000 IOPS
  - Good for: boot volumes, dev/test

- **io2 Block Express (Provisioned IOPS SSD):**
  - 256,000 IOPS, 4,000 MB/s throughput
  - 99.999% durability
  - Sub-millisecond latency
  - Use for: mission-critical, large databases

- **io2/io1 (Provisioned IOPS SSD):**
  - Up to 64,000 IOPS (io2), 32,000 MB/s
  - Multi-attach (share volume across instances in same AZ)

**HDD (Hard Disk Drives):**
- **st1 (Throughput Optimized HDD):**
  - 500 MB/s max throughput
  - Good for: big data, data warehouses, log processing

- **sc1 (Cold HDD):**
  - 250 MB/s max throughput
  - Lowest cost
  - Good for: infrequent access

**EBS Optimization:**
- EBS-optimized instances (dedicated bandwidth)
- Use RAID 0 for increased performance (striping)
- EBS snapshots (don't impact performance of in-use volumes)

**EFS Performance:**

**Performance Modes:**
- **General Purpose:** Low latency, default choice
- **Max I/O:** Higher latency, higher throughput, massive parallelism

**Throughput Modes:**
- **Bursting:** Throughput scales with file system size
- **Provisioned:** Set throughput independent of size
- **Elastic:** Automatically scale throughput (newest, recommended)

**Storage Classes:**
- **Standard:** Frequent access
- **Infrequent Access (IA):** Lower cost, retrieval fee
- Lifecycle management to move files to IA

**Hybrid and Migration Storage:**

**AWS Storage Gateway:**
- **File Gateway:** NFS/SMB interface to S3
- **Volume Gateway:**
  - Stored volumes: Full copy on-premises, async backup to S3
  - Cached volumes: Frequently accessed data on-premises, full copy in S3
- **Tape Gateway:** Virtual tape library backup to S3/Glacier

**AWS DataSync:**
- Automated data transfer between on-premises and AWS
- 10x faster than open-source tools
- Bandwidth throttling
- Transfer to: S3, EFS, FSx

**AWS Snow Family:**
- **Snowcone:** 8 TB usable storage, portable
- **Snowball Edge:**
  - Storage Optimized: 80 TB
  - Compute Optimized: 42 TB + compute
- **Snowmobile:** Exabyte-scale (up to 100 PB per truck)

**Database Storage Performance:**
- Aurora: 6 copies, auto-scaling up to 128 TB
- RDS: Choose appropriate storage type (gp3, io1)
- DynamoDB: Auto-scaling, on-demand or provisioned capacity

---

### Task Statement 3.2: Design high-performing and elastic compute solutions

**Knowledge of:**
- AWS compute services with appropriate use cases
- Distributed computing concepts
- Queuing and messaging concepts
- Scalability capabilities
- Serverless technologies
- Virtualization technologies

**Skills in:**
- Decoupling workloads
- Identifying metrics and conditions for scaling
- Selecting appropriate compute options and features
- Selecting appropriate resource type and size

**Topic Overview:**

**EC2 Instance Selection:**

**Instance Families:**
- **General Purpose (T, M):**
  - T3/T4g: Burstable, baseline CPU with credits
  - M6i/M7g: Balanced compute, memory, networking

- **Compute Optimized (C):**
  - C6i/C7g: High-performance processors
  - Use for: batch processing, gaming, HPC, scientific modeling

- **Memory Optimized (R, X, High Memory):**
  - R6i/R7g: Large memory, high memory-to-CPU ratio
  - X2: Lowest cost per GiB RAM
  - Use for: in-memory databases, real-time big data analytics

- **Storage Optimized (I, D, H):**
  - I4i: NVMe SSD, high IOPS
  - D3: Dense HDD storage
  - Use for: NoSQL databases, data warehousing, Elasticsearch

- **Accelerated Computing (P, G, F):**
  - P4: GPU instances for ML training
  - G5: GPU for graphics/ML inference
  - F1: FPGA for hardware acceleration

**Graviton Instances:**
- AWS-designed ARM processors
- 40% better price-performance vs. x86
- T4g, M7g, C7g, R7g
- Great for: web servers, containerized microservices

**Placement Groups:**
- **Cluster:** Low-latency, high throughput, same AZ (HPC)
- **Spread:** Separate hardware, max 7 instances per AZ per group (critical instances)
- **Partition:** Large distributed workloads (Hadoop, Cassandra)

**Enhanced Networking:**
- Elastic Network Adapter (ENA): Up to 100 Gbps
- Elastic Fabric Adapter (EFA): HPC and ML, bypasses OS kernel

**Serverless Compute:**

**AWS Lambda:**
- Memory: 128 MB to 10,240 MB (CPU scales with memory)
- Timeout: 900 seconds (15 minutes) max
- **Provisioned Concurrency:** Pre-initialized execution environments (reduce cold starts)
- **Reserved Concurrency:** Limit concurrency for function
- Use Lambda Power Tuning tool to optimize memory/cost

**AWS Fargate:**
- Serverless containers
- No EC2 instances to manage
- Right-sized compute and memory
- Use with ECS or EKS

**Container Compute:**

**Amazon ECS:**
- **EC2 Launch Type:**
  - Full control over instances
  - Can use Reserved/Spot instances
  - Better for: large workloads, need specific instance types

- **Fargate Launch Type:**
  - Serverless
  - Pay per vCPU and memory
  - Better for: variable workloads, operational simplicity

**Amazon EKS:**
- Managed Kubernetes control plane
- **Node options:**
  - Managed node groups (EC2)
  - Self-managed nodes
  - Fargate
- Use for: Kubernetes expertise, multi-cloud, complex orchestration

**Batch Processing:**

**AWS Batch:**
- Fully managed batch processing
- Dynamically provisions optimal compute resources
- Integrates with EC2 and Spot instances
- Job queues, job definitions, compute environments

**Performance Optimization:**

**Compute Scaling Strategies:**
- Auto Scaling based on CPU, memory, custom metrics
- Predictive scaling for cyclic patterns
- Step scaling for gradual increases
- Target tracking for maintaining metric

**Compute Cost Optimization:**
- Use Spot instances for fault-tolerant workloads (90% discount)
- Reserved instances for baseline capacity
- Savings Plans for flexible commitment
- Right-size instances using Compute Optimizer

---

### Task Statement 3.3: Determine high-performing database solutions

**Knowledge of:**
- AWS global infrastructure (Availability Zones, Regions)
- Caching strategies and services (ElastiCache)
- Data access patterns (read-intensive vs. write-intensive)
- Database capacity planning (capacity units, instance types)
- Database connections and proxies
- Database engines with appropriate use cases
- Database replication (read replicas)
- Database types and services (serverless, relational vs. non-relational, in-memory)

**Skills in:**
- Configuring read replicas to meet business requirements
- Designing database architectures
- Determining appropriate database engines
- Determining appropriate database types for specific workloads
- Integrating caching to meet business requirements

**Topic Overview:**

**Relational Database Performance:**

**Amazon RDS:**

**Read Replicas:**
- **Asynchronous replication** from primary
- Up to 15 read replicas per instance (Aurora)
- Can be in different AZs or regions
- Use for: Read-heavy workloads, reporting, analytics
- Can promote to standalone database

**Multi-AZ:**
- **Synchronous replication** to standby
- Automatic failover (no data loss)
- Not for scaling reads (standby not accessible)
- Use for: High availability and disaster recovery

**RDS Proxy:**
- Connection pooling
- Reduce database load from opening/closing connections
- Improve failover time (66% faster)
- Use with Lambda (many concurrent connections)

**Performance Insights:**
- Database performance monitoring
- Identify performance bottlenecks
- Wait event analysis
- Top SQL queries

**Amazon Aurora:**

**Performance Features:**
- **Aurora Global Database:**
  - Span multiple regions
  - <1 second cross-region replication
  - Fast regional failover

- **Aurora Serverless v2:**
  - Instant scaling from 0.5 to 128 ACUs
  - Fine-grained scaling
  - Use for: Variable, unpredictable workloads

- **Aurora Parallel Query:**
  - Push down query processing to storage layer
  - Faster analytical queries on current data

- **Aurora Auto Scaling:**
  - Add read replicas based on CPU or connections

**Database Engine Selection:**
- **MySQL:** Open source, wide ecosystem
- **PostgreSQL:** Advanced features, geospatial (PostGIS)
- **MariaDB:** MySQL fork, additional features
- **Oracle:** Enterprise features, licensing considerations
- **SQL Server:** Windows applications, .NET integration
- **Aurora (MySQL/PostgreSQL compatible):** Best performance, AWS-native features

**Non-Relational Database Performance:**

**Amazon DynamoDB:**

**Capacity Modes:**
- **On-Demand:**
  - Pay per request
  - Automatic scaling
  - Use for: Unpredictable workloads, spiky traffic

- **Provisioned:**
  - Specify RCUs and WCUs
  - Auto-scaling available
  - Reserved capacity for cost savings
  - Use for: Predictable workloads

**Performance Features:**
- **DynamoDB Accelerator (DAX):**
  - In-memory cache
  - Microsecond latency
  - Write-through cache
  - No application changes needed

- **Global Tables:**
  - Multi-region, multi-active replication
  - <1 second replication
  - Automatic conflict resolution

- **DynamoDB Streams:**
  - Capture item-level changes
  - Integrate with Lambda for real-time processing

**Design for Performance:**
- Choose right partition key (uniform distribution)
- Use composite sort keys for query flexibility
- Avoid hot partitions
- Use sparse indexes
- Query vs. Scan (query is more efficient)

**Caching Strategies:**

**Amazon ElastiCache:**

**Redis:**
- **Features:**
  - Advanced data structures (lists, sets, sorted sets, hashes)
  - Pub/Sub messaging
  - Persistence (RDB snapshots, AOF)
  - Replication and Multi-AZ
  - Cluster mode for horizontal scaling

- **Use cases:**
  - Session store
  - Leaderboards (sorted sets)
  - Geospatial data
  - Rate limiting

**Memcached:**
- **Features:**
  - Simple key-value store
  - Multi-threaded
  - Horizontal scaling (sharding)
  - No persistence

- **Use cases:**
  - Simple caching
  - Multi-threaded performance needs

**Caching Patterns:**
- **Lazy Loading (Cache-Aside):**
  - Load data into cache on cache miss
  - Only cache requested data
  - Stale data possible

- **Write-Through:**
  - Add/update cache when database is written
  - Data always current
  - Cache churn if data not accessed

- **TTL (Time-To-Live):**
  - Expire old data
  - Balance freshness and load

**Specialized Databases:**

**Amazon Redshift:**
- Petabyte-scale data warehouse
- Columnar storage
- Massively parallel processing (MPP)
- Redshift Spectrum: Query S3 data directly
- Use for: OLAP, analytics, business intelligence

**Amazon Neptune:**
- Graph database
- Use for: Social networks, fraud detection, knowledge graphs

**Amazon DocumentDB:**
- MongoDB-compatible
- Use for: Content management, catalogs, user profiles

**Amazon Timestream:**
- Time-series database
- Use for: IoT, operational metrics, application monitoring

**Amazon QLDB (Quantum Ledger Database):**
- Immutable, cryptographically verifiable transaction log
- Use for: Financial transactions, audit trails

**Connection Management:**
- Use connection pooling
- RDS Proxy for serverless/Lambda
- Monitor max_connections parameter
- Use read replicas to distribute connection load

---

### Task Statement 3.4: Determine high-performing network architectures

**Knowledge of:**
- Edge networking services
- Network architecture (VPC, subnets, route tables)
- Network connection options (Direct Connect, VPN)
- Network performance metrics and monitoring

**Skills in:**
- Selecting appropriate network services for performance
- Determining network configuration for high throughput
- Implementing edge services for performance

**Topic Overview:**

**VPC Performance:**

**Subnet Design:**
- Plan CIDR blocks for growth
- Use /16 for VPC, /24 for subnets (recommended)
- Reserve space for future expansion

**Route Tables:**
- Most specific route wins
- Local route for VPC CIDR (cannot be modified)
- Avoid route proliferation (impacts performance)

**Elastic Network Interfaces (ENIs):**
- Attach multiple ENIs for separate management/data networks
- ENI warm pools (pre-create for faster container startup)

**Enhanced Networking:**
- **ENA (Elastic Network Adapter):**
  - Up to 100 Gbps bandwidth
  - Supported on most modern instance types

- **EFA (Elastic Fabric Adapter):**
  - For HPC and ML workloads
  - Bypasses OS kernel (OS-bypass)
  - Supports MPI (Message Passing Interface)

**VPC Endpoints for Performance:**
- **Gateway Endpoints (S3, DynamoDB):**
  - No hourly charge, no data processing charge
  - Private connectivity without NAT or Internet Gateway

- **Interface Endpoints (PrivateLink):**
  - Private connectivity to AWS services and SaaS
  - Charged per hour and per GB processed
  - Better than going through internet

**Content Delivery and Edge Services:**

**Amazon CloudFront:**

**Performance Features:**
- 450+ global edge locations
- Edge caching reduces origin load
- Regional Edge Caches (larger, fewer locations)
- HTTP/2 and HTTP/3 support
- WebSocket support

**Optimization Techniques:**
- **Origin Shield:** Additional caching layer
- **Field-Level Encryption:** Encrypt sensitive fields at edge
- **Signed URLs/Cookies:** Control access to content
- **Lambda@Edge:** Run code at edge locations
- **CloudFront Functions:** Even lower latency (sub-millisecond)

**Cache Behavior:**
- TTL configuration
- Cache based on headers, cookies, query strings
- Cache hit ratio monitoring

**AWS Global Accelerator:**

**Use Cases:**
- Improve global application availability and performance
- Static Anycast IP addresses
- Automatic routing to healthy endpoints
- DDoS protection with AWS Shield

**How It Works:**
- Traffic enters AWS network at nearest edge location
- Routes over AWS global network (not internet)
- 60% performance improvement for some applications

**Comparison with CloudFront:**
- **CloudFront:** Caching, HTTP/HTTPS content delivery
- **Global Accelerator:** Non-HTTP use cases (TCP/UDP), static IPs, fast regional failover

**Direct Connect and VPN:**

**AWS Direct Connect:**

**Performance Benefits:**
- **Bandwidth:** 50 Mbps to 100 Gbps
- Consistent network performance
- Reduced bandwidth costs for high volumes
- Private connectivity to AWS

**Direct Connect Gateway:**
- Connect to multiple VPCs across regions
- Single Direct Connect connection to many VPCs

**LAG (Link Aggregation Groups):**
- Combine multiple connections for higher throughput
- Up to 4 connections in a LAG

**AWS VPN:**
- IPsec VPN over internet
- Up to 1.25 Gbps per tunnel (2 tunnels per VPN)
- VPN CloudHub for multiple site-to-site VPNs
- Accelerated VPN: VPN + Global Accelerator for better performance

**Route 53 Performance:**

**Routing Policies for Performance:**
- **Latency-based routing:** Route to lowest latency endpoint
- **Geoproximity routing:** Route based on location with bias
- **Multi-value answer:** Return multiple IPs, client chooses

**DNS Optimization:**
- Alias records (no charge, no TTL)
- Health checks with automatic failover
- Traffic flow for visual policy management

**Network Performance Monitoring:**

**VPC Flow Logs:**
- Capture IP traffic information
- Analyze traffic patterns
- Troubleshoot connectivity issues
- Security and compliance

**CloudWatch Metrics:**
- NetworkIn/NetworkOut
- NetworkPacketsIn/NetworkPacketsOut
- EBS throughput metrics

**VPC Reachability Analyzer:**
- Network path analysis
- Troubleshoot connectivity issues
- No packet sending (config analysis only)

---

### Task Statement 3.5: Determine high-performing data ingestion and transformation solutions

**Knowledge of:**
- Data analytics and visualization services
- Data ingestion patterns
- Data transfer services and migration
- ETL (Extract, Transform, Load) services
- Streaming data services

**Skills in:**
- Building and securing data lakes
- Designing data streaming architectures
- Designing solutions for data transfer and migration
- Implementing visualization services
- Selecting appropriate compute options for data processing

**Topic Overview:**

**Real-Time Data Ingestion:**

**Amazon Kinesis:**

**Kinesis Data Streams:**
- Real-time data streaming
- Retain data 24 hours to 365 days
- Shards for parallelism (1 MB/s write, 2 MB/s read per shard)
- Consumers: Lambda, Kinesis Data Analytics, custom apps
- Use for: Real-time analytics, log aggregation, IoT telemetry

**Kinesis Data Firehose:**
- Load streaming data to destinations
- Near real-time (60 seconds minimum latency)
- Automatic scaling
- Destinations: S3, Redshift, Elasticsearch, Splunk, custom HTTP
- Data transformation with Lambda
- No shard management
- Use for: Simple streaming to data lake/warehouse

**Kinesis Data Analytics:**
- Real-time analytics on streaming data
- SQL or Apache Flink
- Auto-scaling
- Use for: Real-time dashboards, metrics, anomaly detection

**Amazon MSK (Managed Streaming for Apache Kafka):**
- Fully managed Apache Kafka
- Use when you need Kafka specifically
- Compatible with Kafka tools and APIs

**Batch Data Ingestion:**

**AWS Database Migration Service (DMS):**
- Migrate databases to AWS
- Homogeneous (Oracle to Oracle) or heterogeneous (Oracle to Aurora)
- Continuous replication
- Schema Conversion Tool (SCT) for heterogeneous migrations
- Zero downtime migration

**AWS DataSync:**
- Transfer data between on-premises and AWS
- 10x faster than open-source tools
- Schedule transfers
- Data validation

**AWS Snow Family:**
- Physical data transport for large volumes
- Use when network transfer would take too long

**Data Transformation (ETL):**

**AWS Glue:**

**Glue Data Catalog:**
- Central metadata repository
- Integrates with Athena, Redshift Spectrum, EMR
- Automatic schema discovery with crawlers

**Glue ETL Jobs:**
- Serverless Spark or Python Shell
- Visual ETL designer
- Generated code (editable)
- Job bookmarks (track processed data)
- Use for: Preparing data for analytics

**Glue DataBrew:**
- Visual data preparation tool
- No code required
- 250+ pre-built transformations

**Amazon EMR (Elastic MapReduce):**
- Managed Hadoop, Spark, HBase, Presto, Flink
- EC2 or EKS deployment
- Use for: Large-scale data processing, ML, custom applications

**Data Lakes:**

**Amazon S3 as Data Lake:**
- Scalable, durable storage
- Store structured, semi-structured, unstructured data
- Lifecycle policies for cost optimization
- S3 Object Lock for compliance

**AWS Lake Formation:**
- Build and manage data lakes
- Centralized permissions
- Data ingestion and cataloging
- Fine-grained access control (column/row level)

**Data Analytics:**

**Amazon Athena:**
- Serverless SQL queries on S3
- Pay per query (amount of data scanned)
- ANSI SQL, JDBC/ODBC
- Federated queries (query non-S3 sources)
- Use partitioning and compression to reduce costs

**Amazon Redshift:**
- Petabyte-scale data warehouse
- Columnar storage, compression
- Massively parallel processing
- **Redshift Spectrum:** Query S3 directly without loading
- Concurrency Scaling for burst workloads

**Amazon QuickSight:**
- Business intelligence and visualization
- Serverless
- ML Insights (anomaly detection, forecasting)
- Embedded analytics
- SPICE engine (in-memory)

**Amazon OpenSearch Service (Elasticsearch):**
- Search and analytics
- Log analytics
- Real-time application monitoring
- Kibana for visualization

**Data Pipeline and Orchestration:**

**AWS Step Functions:**
- Orchestrate workflows
- Visual workflow designer
- Integrates with Lambda, Batch, ECS, SNS, SQS
- Error handling and retries

**Amazon Managed Workflows for Apache Airflow (MWAA):**
- Managed Apache Airflow
- Orchestrate complex data pipelines
- Python-based DAGs (Directed Acyclic Graphs)

---

## Domain 4: Design Cost-Optimized Architectures (20%)

### Task Statement 4.1: Design cost-optimized storage solutions

**Knowledge of:**
- Access options (S3 presigned URLs, VPC endpoints)
- AWS cost management tools
- AWS storage services with appropriate use cases
- Backup strategies
- Block storage options (EBS volume types)
- Data lifecycles
- Hybrid storage options
- Storage access patterns
- Storage tiering
- Storage types with associated characteristics

**Skills in:**
- Designing appropriate backup and retention policies
- Determining correct storage class based on use case
- Determining most cost-effective data transfer method
- Determining when storage auto-scaling is required
- Managing S3 object lifecycles
- Selecting appropriate backup/archival solution
- Selecting appropriate storage tier
- Selecting correct data lifecycle for storage

**Topic Overview:**

**S3 Cost Optimization:**

**Storage Class Selection:**

**Decision Criteria:**
- Access frequency (frequent, infrequent, rare)
- Access time requirements (immediate, minutes, hours)
- Data criticality (redundancy requirements)

**Cost Comparison (per GB per month, approximate):**
- S3 Standard: $0.023
- S3 Intelligent-Tiering: $0.023 + monitoring fee
- S3 Standard-IA: $0.0125 + retrieval fee
- S3 One Zone-IA: $0.01 + retrieval fee
- S3 Glacier Instant Retrieval: $0.004 + retrieval fee
- S3 Glacier Flexible Retrieval: $0.0036 + retrieval fee
- S3 Glacier Deep Archive: $0.00099 + retrieval fee

**S3 Lifecycle Policies:**
- Transition to IA after 30 days of no access
- Archive to Glacier after 90 days
- Delete after retention period
- Separate rules for current and previous versions

**S3 Intelligent-Tiering:**
- Automatic cost optimization
- Moves objects between access tiers
- No retrieval fees
- Small monthly monitoring fee
- Good for: Unknown or changing access patterns

**S3 Cost Optimization Techniques:**
- Use S3 Storage Lens for usage insights
- Enable S3 Analytics for access pattern recommendations
- Compress data before storing
- Use multipart upload with lifecycle rules to clean incomplete uploads
- Requester Pays for data shared with others
- S3 Select/Glacier Select (retrieve only needed data)

**EBS Cost Optimization:**

**Volume Type Selection:**
- gp3: Best cost/performance for most workloads
- gp2: Consider migrating to gp3 for cost savings
- io2: Only when you need very high IOPS
- st1: Throughput-intensive, lower cost than SSD
- sc1: Lowest cost, infrequent access

**EBS Cost Techniques:**
- Delete unused volumes (orphaned volumes)
- Snapshot and delete old development/test volumes
- Use EBS snapshots instead of keeping full volumes
- Snapshots are incremental (only changed blocks)
- Delete old snapshots
- Copy snapshots to S3 Glacier via Data Lifecycle Manager

**EFS Cost Optimization:**
- Use Infrequent Access (IA) storage class
- Lifecycle management to move files to IA
- Elastic throughput (instead of provisioned)
- Delete unused file systems

**Data Transfer Cost Optimization:**

**Understanding Data Transfer Costs:**
- **Free:**
  - Inbound data transfer
  - Data transfer between services in same region
  - Data transfer between S3 and CloudFront

- **Charged:**
  - Outbound to internet
  - Cross-region data transfer
  - Data transfer through NAT Gateway

**Optimization Techniques:**
- Use VPC endpoints for S3/DynamoDB (avoid NAT Gateway costs)
- CloudFront for content delivery (lower data transfer costs)
- Direct Connect for large, ongoing transfers (cheaper than internet)
- Compress data before transfer
- S3 Transfer Acceleration (faster, but additional cost)
- Keep data and compute in same region when possible

**Backup and Archival:**

**AWS Backup:**
- Centralized backup management
- Automated backup scheduling
- Cross-region and cross-account backup
- Lifecycle policies (move to cold storage)
- Cost: Storage + restore costs

**Backup Strategy Cost Optimization:**
- Incremental backups (only changes)
- Define retention policies (don't keep forever)
- Archive old backups to Glacier
- Use S3 Lifecycle for backup data in S3
- Test restores to avoid paying for useless backups

---

### Task Statement 4.2: Design cost-optimized compute solutions

**Knowledge of:**
- AWS cost management service features
- AWS cost management tools
- AWS global infrastructure
- AWS purchasing options (Spot, Reserved, Savings Plans)
- Distributed computing strategies (edge processing)
- Hybrid compute options (Outposts)
- Instance types, families, and sizes
- Optimization of compute utilization
- Scaling strategies
- Serverless compute options

**Skills in:**
- Determining appropriate load balancing strategy
- Determining appropriate scaling methods and strategies
- Determining cost-effective AWS compute services
- Determining required instance size based on business requirements
- Selecting appropriate instance family for workload
- Selecting appropriate instance size for workload

**Topic Overview:**

**EC2 Purchasing Options:**

**On-Demand Instances:**
- Pay by second (Linux/Windows)
- No commitment
- Use for: Short-term, unpredictable workloads
- Most expensive option

**Reserved Instances:**
- **Commitment:** 1-year or 3-year
- **Discount:** Up to 75% vs. On-Demand
- **Types:**
  - Standard RI: Highest discount, locked to instance type/family
  - Convertible RI: Can change instance type, lower discount (~54%)
  - Scheduled RI: Reserve for specific time windows (deprecated, use Savings Plans)
- **Payment Options:**
  - All Upfront: Highest discount
  - Partial Upfront: Medium discount
  - No Upfront: Lowest discount, monthly payments
- Use for: Steady-state workloads (databases, always-on apps)

**Savings Plans:**
- Commit to consistent usage ($/hour) for 1 or 3 years
- Up to 72% discount
- **Types:**
  - **Compute Savings Plans:**
    - Most flexible
    - Apply to EC2, Lambda, Fargate
    - Any instance family, region, OS, tenancy
  - **EC2 Instance Savings Plans:**
    - Locked to instance family and region
    - Any size, OS, tenancy within family
    - Higher discount than Compute Savings Plans
- Recommendations in Cost Explorer
- Use for: Flexible commitment, evolving workloads

**Spot Instances:**
- **Discount:** Up to 90% vs. On-Demand
- Can be interrupted by AWS with 2-minute warning
- **Strategies:**
  - Diversify across multiple instance types
  - Use Spot Fleet (mix of Spot and On-Demand)
  - Spot Block (deprecated): Reserve for 1-6 hours
  - Capacity-optimized allocation strategy
- **Use cases:**
  - Batch processing
  - Data analysis
  - Stateless web servers
  - CI/CD
  - High-performance computing
- **Not suitable for:**
  - Databases
  - Critical workloads that can't handle interruption

**Dedicated Hosts/Instances:**
- Physical server dedicated to you
- Most expensive
- Use only when needed for: Compliance, BYOL (Oracle, Windows Server)

**Rightsizing:**

**AWS Compute Optimizer:**
- Machine learning-based recommendations
- Analyzes CloudWatch metrics
- Recommends optimal instance types
- Considers: CPU, memory, network, disk
- Savings estimation

**Rightsizing Best Practices:**
- Start with reasonable estimates, then adjust
- Monitor utilization (CloudWatch)
- Scale down over-provisioned instances
- Use burstable instances (T3/T4g) for variable workloads
- Review monthly using Compute Optimizer

**Serverless Cost Optimization:**

**AWS Lambda:**
- **Pricing:**
  - Per request: $0.20 per 1M requests
  - Per GB-second: $0.0000166667
- **Optimization:**
  - Right-size memory (CPU scales with memory)
  - Use Lambda Power Tuning tool
  - Reduce package size (faster cold starts, lower cost)
  - Increase memory for faster execution (can be cheaper overall)
  - Use ARM (Graviton2) for 20% cost savings

**AWS Fargate:**
- Pay per vCPU and memory per second
- No over-provisioning like EC2
- Use Fargate Spot for 70% discount (interruptible)

**Scaling Strategies for Cost:**

**Auto Scaling:**
- Scale in aggressively, scale out conservatively
- Use scheduled scaling for predictable patterns
- Use target tracking (maintain CPU at X%)
- Set appropriate cooldown periods
- Use predictive scaling to pre-scale

**Right Sizing Auto Scaling Groups:**
- Don't over-provision min/max/desired capacity
- Review and adjust based on actual usage

**Load Balancing Cost Optimization:**
- Application Load Balancer (ALB): Charged per hour + LCU (Load Balancer Capacity Units)
- Network Load Balancer (NLB): Charged per hour + LCU, also per GB processed
- ALB is better for HTTP/HTTPS (more features, similar cost)
- Use NLB only when you need Layer 4 or static IP

**Edge Computing for Cost:**

**CloudFront:**
- Reduce origin compute load
- Cache content at edge
- Cheaper data transfer than EC2/S3 direct
- CloudFront Functions (lower cost than Lambda@Edge)

**Lambda@Edge:**
- Run code at edge locations
- Reduce main compute load
- Consider CloudFront Functions for simpler use cases (cheaper)

**AWS Outposts:**
- Hybrid cloud (AWS infrastructure on-premises)
- Use when: Data residency, low latency requirements
- More expensive than pure cloud, but can meet specific needs

---

### Task Statement 4.3: Design cost-optimized database solutions

**Knowledge of:**
- AWS cost management service features
- AWS cost management tools
- Caching strategies
- Data retention policies
- Database capacity planning (capacity units, instance types)
- Database connections and proxies
- Database engines with appropriate use cases
- Database replication (read replicas)
- Database types and services (serverless, relational vs. non-relational)

**Skills in:**
- Designing appropriate backup and retention policies
- Determining an appropriate database engine
- Determining cost-effective AWS database services
- Determining cost-effective AWS database types (serverless, relational vs. non-relational)

**Topic Overview:**

**RDS Cost Optimization:**

**Instance Selection:**
- Right-size database instances (use CloudWatch metrics)
- Use Graviton instances (db.t4g, db.m6g, etc.) for 40% better price-performance
- Reserve database instances (1 or 3 years) for steady workloads
- Use Aurora Serverless v2 for variable workloads

**Storage Optimization:**
- Choose appropriate storage type:
  - gp3: General purpose, most cost-effective (provision IOPS independently)
  - gp2: Consider migrating to gp3
  - io1: Only for high-performance needs
- Enable storage autoscaling with appropriate max limit
- Monitor allocated vs. used storage

**Backup Cost:**
- Automated backups included (up to 100% of DB size)
- Snapshots beyond that are charged
- Delete old manual snapshots
- Use Backup Vault for lifecycle management

**Multi-AZ:**
- Doubles cost (second instance)
- Use only for production/critical databases
- Not needed for dev/test

**Read Replicas:**
- Can be more cost-effective than scaling up primary
- Use for read-heavy workloads
- Can be cross-region (additional data transfer cost)

**Aurora Cost Optimization:**

**Aurora vs. RDS:**
- Aurora more expensive per hour
- But: Better performance, auto-scaling storage, faster replication
- Calculate TCO (Total Cost of Ownership), not just hourly cost

**Aurora Serverless v2:**
- Scale from 0.5 to 128 ACUs
- Pay per ACU per second
- Use for: Intermittent, unpredictable workloads
- Set min/max ACUs to control costs

**Aurora Global Database:**
- Expensive (replication across regions)
- Use only when truly needed (DR, global reads)

**Aurora I/O-Optimized:**
- No I/O charges (included in instance price)
- Use when I/O costs > 25% of Aurora database spend

**DynamoDB Cost Optimization:**

**Capacity Modes:**

**On-Demand:**
- Pay per request
- 2.5x more expensive than provisioned (per RCU/WCU)
- Use for: Unpredictable workloads, new applications
- Good when you don't want to manage capacity

**Provisioned:**
- Specify RCUs and WCUs
- Auto-scaling available
- Reserved Capacity (1-3 years) for additional savings
- Use for: Predictable workloads
- Most cost-effective for steady traffic

**DynamoDB Cost Techniques:**
- Use auto-scaling to avoid over-provisioning
- Enable TTL to auto-delete expired items (free)
- Use sparse indexes (only items with attribute are indexed)
- Optimize item size (smaller items = lower costs)
- Use batch operations when possible
- Consider DynamoDB Standard-IA for infrequently accessed tables

**DynamoDB Accelerator (DAX):**
- Adds cost (DAX cluster)
- But reduces DynamoDB read capacity needed
- Calculate if caching cost < reduced read cost
- Use for: Read-heavy, latency-sensitive workloads

**Caching for Database Cost Reduction:**

**ElastiCache:**
- Redis or Memcached
- Reduces database load
- Can use smaller/fewer database instances
- Reserved nodes for additional savings

**When Caching Makes Financial Sense:**
- High read-to-write ratio
- Repeated queries for same data
- Database is bottleneck
- Compare: cache cost < database scaling cost

**Database Engine Selection:**

**Cost Considerations:**
- **Open source (MySQL, PostgreSQL, MariaDB):** No license fees
- **Commercial (Oracle, SQL Server):** License included or BYOL
  - License Included: Higher per-hour cost
  - BYOL: Need to manage licenses, but can be cheaper
- **Aurora:** More expensive per hour, but better performance (may need fewer/smaller instances)

**Database Type Selection:**

**Relational vs. Non-Relational:**
- **DynamoDB:**
  - Better for: Simple queries, key-value access, massive scale
  - Can be cheaper at scale (no instance costs, just throughput)
- **RDS/Aurora:**
  - Better for: Complex queries, joins, transactions
  - Fixed instance costs

**Serverless vs. Provisioned:**
- **Serverless (Aurora Serverless, DynamoDB On-Demand):**
  - Good for: Intermittent usage, unpredictable workloads
  - No idle costs (Aurora Serverless can scale to 0)
- **Provisioned:**
  - Good for: Consistent usage
  - More cost-effective for steady workloads

**Specialized Databases for Cost Efficiency:**
- Don't use RDS for time-series data (use Timestream)
- Don't use DynamoDB for analytics (use Redshift or Athena on S3)
- Use right tool for the job (avoid over-provisioning general-purpose DB)

**Data Retention and Cleanup:**
- Implement data retention policies
- Archive old data to S3 (much cheaper than database storage)
- Delete or archive based on business requirements
- Use DynamoDB TTL for automatic deletion

**Connection Management:**
- Use RDS Proxy to reduce connections (especially with Lambda)
- Reduces database load
- Can use smaller instance size

---

### Task Statement 4.4: Design cost-optimized network architectures

**Knowledge of:**
- AWS cost management service features
- AWS cost management tools
- Data transfer costs
- Load balancing concepts
- NAT gateways
- Network connectivity (VPN, Direct Connect)
- Network routing, topology, and peering
- Network services with appropriate use cases

**Skills in:**
- Configuring appropriate NAT gateway types
- Configuring appropriate network connections (Direct Connect, VPN)
- Configuring appropriate network routes to minimize costs
- Determining strategic needs for CDN and edge caching
- Reviewing existing workloads for network optimizations
- Selecting appropriate throttling strategy
- Selecting bandwidth allocation for network devices

**Topic Overview:**

**Data Transfer Cost Principles:**

**AWS Data Transfer Costs:**
- **Free:**
  - Inbound from internet
  - Between services in same AZ (using private IP)
  - Between S3 and CloudFront
  - Between EC2 and CloudFront

- **Low Cost:**
  - Between services in same region, different AZs: $0.01/GB

- **Higher Cost:**
  - Outbound to internet: $0.09/GB (decreases with volume)
  - Cross-region: $0.02/GB
  - Through NAT Gateway: $0.045/GB processed + data transfer costs

**Cost Optimization Strategies:**

**VPC Endpoint Optimization:**

**Gateway Endpoints (S3, DynamoDB):**
- **Free** (no hourly charge, no data processing charge)
- Avoids NAT Gateway costs
- Always use for S3/DynamoDB access from private subnets

**Interface Endpoints (PrivateLink):**
- Charged: ~$0.01/hour + $0.01/GB processed
- Still cheaper than NAT Gateway for most AWS services
- Use for: High-volume AWS service access from private subnets

**Cost Comparison Example (S3 access from private subnet):**
- NAT Gateway: $0.045/GB + $0.09/GB internet transfer = $0.135/GB
- Gateway Endpoint: $0 (free)
- **Savings: 100%**

**NAT Gateway Cost Optimization:**

**NAT Gateway Costs:**
- $0.045 per hour (~$32/month)
- $0.045 per GB processed
- Each AZ needs separate NAT Gateway for HA

**Cost Reduction Strategies:**
- Use VPC endpoints instead (for AWS services)
- Single NAT Gateway for dev/test (not HA, but saves cost)
- NAT Instance (cheaper, but you manage it)
  - t3.nano can handle low traffic
  - Configure instance with iptables
  - Not as highly available as NAT Gateway
- S3 Gateway Endpoint for S3 access (free)
- Reduce unnecessary internet-bound traffic

**Direct Connect vs. VPN:**

**AWS Direct Connect:**
- **Costs:**
  - Port hours: $0.30/hour for 1 Gbps
  - Data transfer out: $0.02/GB (cheaper than internet)
- **Use when:**
  - High data transfer volumes (>1 TB/month typically)
  - Need consistent network performance
  - Ongoing connectivity needed
- **ROI calculation:**
  - Compare: (internet transfer cost) vs. (port cost + DX transfer cost)
  - Typically breaks even at 5-15 TB/month

**AWS VPN:**
- **Costs:**
  - $0.05/hour per VPN connection
  - Standard data transfer rates
- **Use when:**
  - Low to moderate data transfer
  - Backup for Direct Connect
  - Quick setup needed

**CloudFront for Cost Optimization:**

**When CloudFront Saves Money:**
- Serving content to geographically distributed users
- High volume of requests to same content
- Outbound data transfer from CloudFront cheaper than EC2/S3

**Cost Comparison:**
- S3 to internet: $0.09/GB (first 10 TB)
- CloudFront to internet: $0.085/GB (varies by region)
- Plus CloudFront reduces origin load

**CloudFront Cost Optimization:**
- Use Origin Shield for high request rates to origin
- Set appropriate TTLs (cache longer = fewer origin requests)
- Use compression (reduces data transfer)
- Price Class: Choose distribution coverage (All edge locations vs. reduced set)

**Inter-Region Data Transfer:**

**Costs:**
- $0.02/GB between most regions
- Free within same region (different AZs: $0.01/GB)

**Optimization:**
- Design applications to minimize cross-region traffic
- Use Regional services when possible
- Aggregate data before cross-region transfer
- Compress data before transfer
- Consider S3 same-region replication vs. cross-region based on requirements

**Load Balancer Costs:**

**Application Load Balancer:**
- $0.0225/hour (~$16/month)
- LCU (Load Balancer Capacity Unit) charges
- LCU based on: new connections, active connections, bandwidth, rule evaluations

**Network Load Balancer:**
- $0.0225/hour (~$16/month)
- NLCU charges
- Also charges for data processed
- Generally more expensive than ALB

**Gateway Load Balancer:**
- Similar pricing to NLB

**Cost Optimization:**
- Consolidate multiple applications on single ALB (host-based/path-based routing)
- Don't create separate load balancers unnecessarily
- Use ALB for HTTP/HTTPS (cheaper per GB than NLB typically)

**PrivateLink (VPC Endpoint Services):**
- $0.01/hour per AZ
- $0.01/GB data processed
- Use for: Service provider/consumer model
- More cost-effective than VPC peering for specific service access

**VPC Peering:**
- Free (just pay data transfer)
- $0.01/GB between AZs in same region
- $0.02/GB between regions
- Use for: Full network connectivity between VPCs

**Transit Gateway:**
- $0.05/hour per attachment
- $0.02/GB data processed
- Use when: Multiple VPCs (>3-5) need to communicate
- Cost comparison: Calculate vs. individual VPC peering

**Global Accelerator:**
- $0.025/hour per accelerator
- $0.015/GB data transfer premium
- Use only when performance improvement justifies cost
- For global applications with low-latency requirements

---

## Study Tips for SAA-C03

### Key Focus Areas (by weight):
1. **Design Secure Architectures (30%)** - Master IAM, VPC security, encryption
2. **Design Resilient Architectures (26%)** - Multi-AZ, disaster recovery, high availability
3. **Design High-Performing Architectures (24%)** - Storage performance, compute selection, caching
4. **Design Cost-Optimized Architectures (20%)** - Purchasing options, rightsizing, data transfer costs

### Recommended Study Approach:
1. Hands-on practice is essential (use AWS Free Tier)
2. Build multi-tier architectures with Auto Scaling and load balancing
3. Practice disaster recovery scenarios
4. Understand the "why" behind architectural decisions
5. Study the AWS Well-Architected Framework deeply
6. Take practice exams and review incorrect answers
7. Use AWS whitepapers and re:Invent videos

### Common Exam Patterns:
- Compare services (when to use Aurora vs. RDS vs. DynamoDB)
- Choose appropriate solution based on requirements
- Identify cost-optimization opportunities
- Design for high availability and fault tolerance
- Implement security best practices

### Comparison with CCP:
- CCP: "What is this service?"
- SAA: "When should I use this service, and how should I architect it?"
- SAA requires deeper technical knowledge
- More focus on trade-offs and architectural decisions
- Real-world scenario-based questions

### Hands-On Labs to Practice:
1. Build VPC with public/private subnets, NAT Gateway, security groups
2. Deploy multi-AZ RDS with read replicas
3. Create Auto Scaling group with ELB
4. Set up S3 bucket with lifecycle policies and encryption
5. Implement CloudFront distribution
6. Configure cross-region disaster recovery
7. Build serverless application with Lambda, API Gateway, DynamoDB
8. Implement monitoring with CloudWatch
9. Set up multi-account structure with AWS Organizations

### Resources:
- AWS Solutions Architect Associate official exam guide
- AWS Well-Architected Framework
- AWS whitepapers (especially "Architecting for the Cloud")
- AWS Architecture Center
- Re:Invent session recordings
- AWS Skill Builder
- Practice exams (official and third-party)

### Time Management:
- 130 minutes for 65 questions = 2 minutes per question
- Flag difficult questions and return later
- Read questions carefully (they often contain specific requirements)
- Eliminate obviously wrong answers first
- Look for keywords in questions (cost-effective, high performance, secure, etc.)

### Next Steps After SAA:
- AWS Certified Solutions Architect Professional
- AWS Certified DevOps Engineer Professional
- Specialty certifications (Security, Database, Data Analytics, etc.)
