# Domain 3: Cloud Technology and Services (34%)

## Task Statement 3.1: Define methods of deploying and operating in the AWS Cloud

**Key Topics:**

- Deployment methods
- Infrastructure as Code
- Access methods
- Operational tools

**Deployment Approaches:**

- Programmatic access (CLI, SDKs)
- Infrastructure as Code (CloudFormation, CDK)
- Management Console (web interface)
- AWS Elastic Beanstalk (PaaS)
- AWS App Runner (containerized apps)

---

## Task Statement 3.2: Define the AWS global infrastructure

**Key Components:**

**Regions:**

- Geographic areas with multiple Availability Zones
- Completely independent and isolated
- Choose based on: latency, compliance, service availability, pricing

**Availability Zones (AZs):**

- One or more discrete data centers
- Redundant power, networking, connectivity
- Connected with high-bandwidth, low-latency networking
- Design for multi-AZ deployment for high availability

> **BEST PRACTICE:** Production workloads should ALWAYS be deployed across multiple AZs. This is the standard architecture pattern for high availability. Single-AZ deployments are typically only used for dev/test environments to save costs.

**Edge Locations:**

- CloudFront content delivery
- Route 53 DNS
- AWS Global Accelerator
- Hundreds of edge locations worldwide

**Local Zones:**

- Place compute and storage closer to end users
- Single-digit millisecond latency
- Extension of a Region

---

## Task Statement 3.3: Identify AWS compute services

**Amazon EC2 (Elastic Compute Cloud):**

- Virtual servers in the cloud
- Full control over operating system
- Various instance types for different workloads

> **COMMON INSTANCE FAMILIES:**
> - **t3/t4g** (general purpose, burstable) - Most common for web servers, dev/test (60%+ of workloads)
> - **m6i/m7i** (general purpose, balanced) - Standard for production applications
> - **c6i/c7i** (compute optimized) - CPU-intensive workloads
> - **r6i/r7i** (memory optimized) - Database and caching workloads

**AWS Lambda:**

- Serverless compute
- Run code without managing servers
- Pay only for compute time used
- Event-driven execution

> **COMMON USAGE:** Lambda has become the default choice for event-driven workloads, API backends (with API Gateway), data processing, and automation tasks. Many organizations adopt a "serverless-first" approach for new services to reduce operational overhead.

**Amazon ECS (Elastic Container Service):**

- Container orchestration
- Run Docker containers

**Amazon EKS (Elastic Kubernetes Service):**

- Managed Kubernetes
- Run Kubernetes workloads

**AWS Elastic Beanstalk:**

- PaaS for deploying applications
- Automatically handles capacity provisioning, load balancing, auto-scaling

**Amazon Lightsail:**

- Simplified compute for simple applications
- Fixed monthly pricing

---

## Task Statement 3.4: Identify AWS database services

**Relational Databases:**

**Amazon RDS (Relational Database Service):**

- Managed relational databases
- Supports: MySQL, PostgreSQL, MariaDB, Oracle, SQL Server
- Automated backups, patching, scaling

> **COMMON USAGE:** RDS is the standard choice for relational databases, used in 80%+ of AWS database deployments. **PostgreSQL** and **MySQL** are the most popular engines. Multi-AZ deployment is standard for production (enables automatic failover).

**Amazon Aurora:**

- MySQL and PostgreSQL compatible
- 5x performance of MySQL, 3x performance of PostgreSQL
- Automatic scaling, high availability

**Non-Relational Databases:**

**Amazon DynamoDB:**

- NoSQL key-value and document database
- Serverless, automatic scaling
- Single-digit millisecond latency

> **COMMON USAGE:** DynamoDB is the most popular NoSQL option on AWS, especially for serverless applications, gaming leaderboards, IoT data, and session storage. On-demand pricing mode is commonly used for unpredictable workloads; provisioned capacity for steady, predictable traffic.

**Other Database Services:**

**Amazon Redshift:**

- Data warehousing
- Petabyte-scale analytics

**Amazon ElastiCache:**

- In-memory caching
- Supports Redis and Memcached

**Amazon Neptune:**

- Graph database

---

## Task Statement 3.5: Identify AWS network services

**Amazon VPC (Virtual Private Cloud):**

- Isolated virtual network
- Complete control over network configuration
- Public and private subnets

**Amazon Route 53:**

- DNS service
- Domain registration
- Health checking and routing policies

**Amazon CloudFront:**

- Content Delivery Network (CDN)
- Global edge locations
- Reduces latency for users

> **COMMON USAGE:** CloudFront is widely used for serving static assets (images, CSS, JS), video streaming, and API acceleration. Often paired with S3 for static website hosting. Standard for any global application to improve performance and reduce origin load.

**AWS Direct Connect:**

- Dedicated network connection from on-premises to AWS
- More consistent network performance

**Elastic Load Balancing:**

- Distribute traffic across multiple targets
- Types: Application Load Balancer, Network Load Balancer, Gateway Load Balancer

> **COMMON USAGE:**
> - **Application Load Balancer (ALB)** - Most commonly used (90%+ of cases), works at Layer 7 (HTTP/HTTPS), supports path-based and host-based routing
> - **Network Load Balancer (NLB)** - Used for extreme performance needs, TCP/UDP traffic, or static IP requirements
> - **Classic Load Balancer** - Legacy, avoid for new deployments

**AWS VPN:**

- Secure connection between on-premises and AWS
- Site-to-Site VPN and Client VPN

**Amazon API Gateway:**

- Create, publish, and manage APIs
- Serverless API management

---

## Task Statement 3.6: Identify AWS storage services

**Amazon S3 (Simple Storage Service):**

- Object storage
- 11 9's of durability (99.999999999%)
- Storage classes for different access patterns:
  - S3 Standard: Frequent access
  - S3 Intelligent-Tiering: Automatic cost optimization
  - S3 Standard-IA: Infrequent access
  - S3 One Zone-IA: Lower cost, single AZ
  - S3 Glacier: Archive storage
  - S3 Glacier Deep Archive: Lowest cost archive

> **COMMON STORAGE CLASSES:**
> - **S3 Standard** - Default and most widely used (70%+ of objects) for active data
> - **S3 Intelligent-Tiering** - Growing in popularity for automatic cost optimization without performance impact
> - **S3 Glacier Flexible Retrieval** - Standard for long-term archives (backups, compliance data)
> - **S3 Standard-IA** - Used for infrequent but immediate access needs (disaster recovery, older logs)

**Amazon EBS (Elastic Block Store):**

- Block storage for EC2 instances
- Persistent storage
- Types: SSD (gp3, io2) and HDD (st1, sc1)

> **COMMON VOLUME TYPES:**
> - **gp3 (General Purpose SSD)** - Most popular choice (80%+ of volumes), best price-performance for most workloads
> - **gp2** - Legacy, being replaced by gp3
> - **io2 Block Express** - Only for highest performance databases requiring 100K+ IOPS
> - **st1 (Throughput Optimized HDD)** - Big data, data warehouses (cost-effective for sequential access)

**Amazon EFS (Elastic File System):**

- Managed file storage
- NFS protocol
- Shared across multiple EC2 instances

**AWS Storage Gateway:**

- Hybrid cloud storage
- Connects on-premises to cloud storage

**AWS Snowball Family:**

- Physical data transport
- Snowball, Snowball Edge, Snowmobile
- For large data migrations (TBs to PBs)

---

## Task Statement 3.7: Identify AWS AI/ML and analytics services

**AI/ML Services:**

**Amazon SageMaker:**

- Build, train, and deploy machine learning models

**Amazon Rekognition:**

- Image and video analysis

**Amazon Comprehend:**

- Natural language processing

**Amazon Lex:**

- Build chatbots

**Amazon Polly:**

- Text-to-speech

**Amazon Transcribe:**

- Speech-to-text

**Analytics Services:**

**Amazon Athena:**

- Query S3 data using SQL
- Serverless

**Amazon QuickSight:**

- Business intelligence and visualization

**Amazon Kinesis:**

- Real-time streaming data
- Kinesis Data Streams, Data Firehose, Data Analytics

**Amazon EMR (Elastic MapReduce):**

- Big data processing
- Hadoop, Spark frameworks

**AWS Glue:**

- ETL (Extract, Transform, Load) service
- Data catalog and preparation

---

## Task Statement 3.8: Identify services from other in-scope AWS service categories

**Application Integration:**

- Amazon SQS (Simple Queue Service): Message queuing
- Amazon SNS (Simple Notification Service): Pub/sub messaging
- Amazon EventBridge: Event bus for application integration

**Developer Tools:**

- AWS CodeCommit: Source control (Git)
- AWS CodeBuild: Build and test code
- AWS CodeDeploy: Automated deployments
- AWS CodePipeline: CI/CD pipeline

**Management and Governance:**

- AWS CloudFormation: Infrastructure as Code
- AWS CloudWatch: Monitoring and observability
- AWS CloudTrail: API logging and governance
- AWS Systems Manager: Operational insights and automation
- AWS Config: Configuration management and compliance
- AWS Trusted Advisor: Best practice recommendations

**Customer Engagement:**

- Amazon Connect: Cloud contact center
- Amazon SES (Simple Email Service): Email sending
