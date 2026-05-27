# AWS Certified Cloud Practitioner (CLF-C02) - Exam Outline

## Exam Overview

- **Exam Code:** CLF-C02
- **Duration:** 90 minutes
- **Question Format:** 65 questions (50 scored, 15 unscored)
- **Question Types:** Multiple choice and multiple response
- **Passing Score:** 700 out of 1000 (scaled score)
- **Cost:** $100 USD
- **Validity:** 3 years

## Exam Purpose

This exam validates foundational AWS Cloud knowledge including cloud concepts, security, billing, pricing, support, compliance, technology, and services. It's designed for individuals with basic understanding of IT services and their uses in the AWS Cloud platform.

---

## Domain 1: Cloud Concepts (24%)

### Task Statement 1.1: Define the benefits of the AWS Cloud

**Knowledge of:**
- Value proposition of the AWS Cloud

**Skills in:**
- Understanding economies of scale (cost savings)
- Understanding benefits of global infrastructure (speed of deployment, global reach)
- Understanding advantages of high availability, elasticity, and agility

**Overview:**
The AWS Cloud provides significant advantages over traditional on-premises infrastructure. Key benefits include:
- **Economies of scale:** AWS purchases hardware at massive scale, passing savings to customers
- **Global infrastructure:** Deploy applications in multiple regions worldwide in minutes
- **High availability:** Design fault-tolerant systems across multiple data centers
- **Elasticity:** Scale resources up or down based on demand
- **Agility:** Experiment and innovate faster with quick resource provisioning

---

### Task Statement 1.2: Identify design principles of the AWS Cloud

**Knowledge of:**
- AWS Well-Architected Framework

**Skills in:**
- Understanding the pillars of the Well-Architected Framework
- Identifying differences between the pillars

**The Six Pillars:**

1. **Operational Excellence**
   - Running and monitoring systems to deliver business value
   - Continually improving processes and procedures

2. **Security**
   - Protecting information and systems
   - Implementing identity and access management

3. **Reliability**
   - Ensuring workloads perform their intended functions correctly and consistently
   - Recovering from failures quickly

4. **Performance Efficiency**
   - Using computing resources efficiently
   - Maintaining efficiency as demand changes

5. **Cost Optimization**
   - Avoiding unnecessary costs
   - Understanding where money is spent

6. **Sustainability**
   - Minimizing environmental impact
   - Maximizing energy efficiency

---

### Task Statement 1.3: Understand the benefits of and strategies for migration to the AWS Cloud

**Knowledge of:**
- Cloud adoption strategies
- Resources to support the cloud migration journey

**Skills in:**
- Understanding benefits of AWS Cloud Adoption Framework (AWS CAF)
- Identifying appropriate migration strategies

**AWS Cloud Adoption Framework (CAF) Benefits:**
- Reduced business risk
- Improved environmental, social, and governance (ESG) performance
- Increased revenue
- Increased operational efficiency

**Migration Strategies (6 R's):**
- **Rehost** (lift-and-shift): Move applications without changes
- **Replatform** (lift-tinker-and-shift): Make minor optimizations
- **Repurchase**: Switch to SaaS
- **Refactor/Re-architect**: Redesign using cloud-native features
- **Retire**: Decommission unnecessary applications
- **Retain**: Keep certain applications on-premises

**Migration Tools:**
- AWS Snowball family (for large data transfers)
- Database Migration Service (DMS)
- AWS Application Migration Service

---

### Task Statement 1.4: Understand concepts of cloud economics

**Knowledge of:**
- Aspects of cloud economics
- Cost savings of moving to the cloud

**Skills in:**
- Understanding fixed costs vs. variable costs
- Understanding on-premises environment costs
- Understanding licensing strategies (BYOL vs. included licenses)
- Understanding rightsizing
- Identifying benefits of automation
- Identifying managed AWS services

**Key Concepts:**

**Fixed vs. Variable Costs:**
- **Fixed costs:** On-premises requires upfront capital expenditure (CapEx) for hardware
- **Variable costs:** Cloud uses operational expenditure (OpEx), paying only for what you use

**On-Premises Costs:**
- Hardware purchase and maintenance
- Data center space, power, and cooling
- IT staff for infrastructure management
- Capacity planning and over-provisioning

**Licensing Models:**
- **BYOL (Bring Your Own License):** Use existing licenses in the cloud
- **Included licenses:** Pay for software licenses as part of instance pricing

**Rightsizing:**
- Matching instance types and sizes to workload requirements
- Continuously optimizing to avoid over-provisioning

**Automation Benefits:**
- Infrastructure as Code (CloudFormation, Terraform)
- Automated provisioning and configuration management
- Reduced human error and faster deployments

**Managed Services:**
- Amazon RDS (managed databases)
- Amazon ECS/EKS (managed containers)
- Amazon DynamoDB (managed NoSQL)
- Reduce operational overhead and management burden

---

## Domain 2: Security and Compliance (30%)

### Task Statement 2.1: Understand the AWS shared responsibility model

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

---

### Task Statement 2.2: Understand AWS Cloud security, governance, and compliance concepts

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

---

### Task Statement 2.3: Identify AWS access management capabilities

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

### Task Statement 2.4: Identify components and resources for security

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

---

## Domain 3: Cloud Technology and Services (34%)

### Task Statement 3.1: Define methods of deploying and operating in the AWS Cloud

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

### Task Statement 3.2: Define the AWS global infrastructure

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

### Task Statement 3.3: Identify AWS compute services

**Amazon EC2 (Elastic Compute Cloud):**
- Virtual servers in the cloud
- Full control over operating system
- Various instance types for different workloads

**AWS Lambda:**
- Serverless compute
- Run code without managing servers
- Pay only for compute time used
- Event-driven execution

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

### Task Statement 3.4: Identify AWS database services

**Relational Databases:**

**Amazon RDS (Relational Database Service):**
- Managed relational databases
- Supports: MySQL, PostgreSQL, MariaDB, Oracle, SQL Server
- Automated backups, patching, scaling

**Amazon Aurora:**
- MySQL and PostgreSQL compatible
- 5x performance of MySQL, 3x performance of PostgreSQL
- Automatic scaling, high availability

**Non-Relational Databases:**

**Amazon DynamoDB:**
- NoSQL key-value and document database
- Serverless, automatic scaling
- Single-digit millisecond latency

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

### Task Statement 3.5: Identify AWS network services

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

**AWS Direct Connect:**
- Dedicated network connection from on-premises to AWS
- More consistent network performance

**Elastic Load Balancing:**
- Distribute traffic across multiple targets
- Types: Application Load Balancer, Network Load Balancer, Gateway Load Balancer

**AWS VPN:**
- Secure connection between on-premises and AWS
- Site-to-Site VPN and Client VPN

**Amazon API Gateway:**
- Create, publish, and manage APIs
- Serverless API management

---

### Task Statement 3.6: Identify AWS storage services

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

**Amazon EBS (Elastic Block Store):**
- Block storage for EC2 instances
- Persistent storage
- Types: SSD (gp3, io2) and HDD (st1, sc1)

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

### Task Statement 3.7: Identify AWS AI/ML and analytics services

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

### Task Statement 3.8: Identify services from other in-scope AWS service categories

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

---

## Domain 4: Billing, Pricing, and Support (12%)

### Task Statement 4.1: Compare AWS pricing models

**Pricing Models:**

**On-Demand:**
- Pay by the hour or second
- No long-term commitments
- Good for: short-term, unpredictable workloads

**Reserved Instances:**
- 1-year or 3-year commitment
- Up to 75% discount vs. On-Demand
- Good for: steady-state workloads
- Types: Standard, Convertible, Scheduled

**Savings Plans:**
- Commit to consistent usage ($/hour) for 1 or 3 years
- Up to 72% discount
- More flexible than Reserved Instances
- Types: Compute Savings Plans, EC2 Instance Savings Plans

**Spot Instances:**
- Bid on unused EC2 capacity
- Up to 90% discount
- Can be interrupted by AWS
- Good for: fault-tolerant, flexible workloads

**Dedicated Hosts:**
- Physical server dedicated to your use
- Compliance requirements or BYOL
- Most expensive option

**Free Tier:**
- 12 months free for many services
- Always free for some services (Lambda, DynamoDB limited usage)
- Trials for other services

---

### Task Statement 4.2: Understand resources for billing, budget, and cost management

**Cost Management Tools:**

**AWS Cost Explorer:**
- Visualize and analyze costs
- Forecast future costs
- Identify cost trends
- Filter by service, region, tags

**AWS Budgets:**
- Set custom cost and usage budgets
- Alerts when exceeding thresholds
- Budget types: cost, usage, reservation, Savings Plans

**AWS Cost and Usage Report:**
- Most comprehensive cost data
- Hourly, daily, or monthly reports
- Delivered to S3
- Integrates with analytics tools

**AWS Pricing Calculator:**
- Estimate costs before deployment
- Create cost estimates for use cases

**Consolidated Billing:**
- One bill for multiple AWS accounts
- Volume discounts across accounts
- Part of AWS Organizations

**Cost Allocation Tags:**
- Track costs by project, department, environment
- User-defined and AWS-generated tags

---

### Task Statement 4.3: Identify AWS technical resources and AWS Support options

**Support Plans:**

**Basic Support (Free):**
- 24/7 customer service
- Documentation and whitepapers
- AWS Trusted Advisor (7 core checks)
- AWS Personal Health Dashboard

**Developer Support ($29+/month):**
- Business hours email access to Cloud Support Associates
- General guidance < 24 hours
- System impaired < 12 hours

**Business Support ($100+/month):**
- 24/7 phone, email, and chat access to Cloud Support Engineers
- Production system down < 1 hour
- AWS Trusted Advisor (all checks)
- Infrastructure Event Management (additional fee)

**Enterprise On-Ramp Support ($5,500+/month):**
- Business-critical system down < 30 minutes
- Pool of Technical Account Managers
- Concierge support team
- Cost optimization workshop

**Enterprise Support ($15,000+/month):**
- Business-critical system down < 15 minutes
- Dedicated Technical Account Manager (TAM)
- Infrastructure Event Management (included)
- Operations reviews and tools

**Other Resources:**

**AWS Documentation:**
- User guides, API references, tutorials

**AWS Knowledge Center:**
- FAQs and solutions to common issues

**AWS re:Post:**
- Community-driven Q&A

**AWS Whitepapers:**
- Technical and business content
- Best practices and architectures

**AWS Partner Network (APN):**
- Consulting and technology partners
- AWS Marketplace (third-party software)

**AWS Training and Certification:**
- Digital training (free)
- Classroom training
- Certification programs

**AWS Events:**
- AWS Summit (regional)
- AWS re:Invent (annual conference)

---

## Study Tips

### Key Focus Areas (by weight):
1. **Cloud Technology and Services (34%)** - Largest domain, know the core services
2. **Security and Compliance (30%)** - Shared responsibility model is critical
3. **Cloud Concepts (24%)** - Understand the "why" of cloud
4. **Billing and Pricing (12%)** - Know pricing models and cost tools

### Recommended Study Approach:
1. Start with AWS Well-Architected Framework
2. Understand the shared responsibility model deeply
3. Learn core services in each category (compute, storage, database, networking)
4. Practice with AWS Free Tier
5. Use AWS training resources (Skill Builder)
6. Take practice exams

### Common Pitfalls to Avoid:
- Confusing customer vs. AWS responsibilities
- Not understanding when to use different pricing models
- Mixing up similar services (EBS vs. EFS vs. S3)
- Overlooking the importance of tags for cost allocation

### Next Steps After CCP:
- AWS Certified Solutions Architect Associate (SAA-C03)
- AWS Certified Developer Associate
- AWS Certified SysOps Administrator Associate
