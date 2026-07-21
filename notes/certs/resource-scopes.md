# AWS Resource Scopes

A reference guide for understanding at what scope common AWS services operate.

## Global Services

These services are not tied to a specific region and operate globally:

- **IAM** (Identity and Access Management) - Users, groups, roles, and policies
- **Route 53** - DNS service
- **CloudFront** - Content Delivery Network (CDN)
- **WAF** (Web Application Firewall) - When used with CloudFront
- **Shield Standard** - DDoS protection
- **Organizations** - Multi-account management
- **IAM Identity Center** (formerly AWS SSO)

## Regional Services

These services are created and managed within a specific AWS region:

### Compute

- **EC2** - Virtual machines (instances are AZ-specific, but AMIs are regional)
- **Lambda** - Serverless compute
- **Elastic Beanstalk** - Application deployment service
- **ECS** (Elastic Container Service) - Container orchestration
- **EKS** (Elastic Kubernetes Service) - Managed Kubernetes
- **Fargate** - Serverless container compute
- **Batch** - Batch computing jobs

### Storage

- **S3** - Object storage (buckets are regional, but names are globally unique)
- **EFS** (Elastic File System) - Network file system
- **FSx** - Managed file systems (Windows, Lustre, NetApp, OpenZFS)
- **Storage Gateway** - Hybrid cloud storage
- **Backup** - Centralized backup service

### Database

- **RDS** - Relational databases (instances are AZ-specific, but managed regionally)
- **DynamoDB** - NoSQL database with optional global tables
- **ElastiCache** - In-memory caching (Redis/Memcached)
- **Neptune** - Graph database
- **DocumentDB** - MongoDB-compatible database
- **Redshift** - Data warehouse
- **Aurora** - MySQL/PostgreSQL-compatible database
- **QLDB** (Quantum Ledger Database)
- **Timestream** - Time series database

### Networking

- **VPC** - Virtual Private Cloud
- **API Gateway** - API management (Regional and Edge-optimized endpoints available)
- **Direct Connect** - Dedicated network connection (connections to specific regions)
- **Transit Gateway** - Network transit hub
- **PrivateLink** - Private connectivity to services
- **VPN** - Virtual Private Network connections
- **Load Balancers** (ALB, NLB, CLB) - Regional but distribute across AZs
- **Global Accelerator** - Uses global network but terminates in regions

### Security & Identity

- **KMS** (Key Management Service) - Encryption keys (regional, with multi-region keys available)
- **Secrets Manager** - Secret storage and rotation
- **Certificate Manager (ACM)** - SSL/TLS certificates
- **Cognito** - User authentication and authorization
- **WAF** - When used with ALB or API Gateway (regional)
- **Shield Advanced** - Enhanced DDoS protection

### Management & Governance

- **CloudWatch** - Monitoring and observability
- **CloudTrail** - API logging (logs are regional, but can enable for all regions)
- **Config** - Resource configuration tracking
- **Systems Manager** - Operational management
- **CloudFormation** - Infrastructure as Code
- **Service Catalog** - IT service catalog
- **Auto Scaling** - Automatic scaling

### Analytics

- **Athena** - Serverless query service
- **EMR** (Elastic MapReduce) - Big data processing
- **Kinesis** - Real-time data streaming
- **Glue** - ETL service
- **QuickSight** - Business intelligence
- **Data Pipeline** - Data workflow orchestration
- **Lake Formation** - Data lake management

### Application Integration

- **SQS** (Simple Queue Service) - Message queuing
- **SNS** (Simple Notification Service) - Pub/sub messaging
- **Step Functions** - Workflow orchestration
- **EventBridge** - Event bus service
- **AppSync** - GraphQL API service
- **MQ** - Managed message broker

### Developer Tools

- **CodeCommit** - Source control
- **CodeBuild** - Build service
- **CodeDeploy** - Deployment automation
- **CodePipeline** - CI/CD pipeline
- **Cloud9** - Cloud IDE
- **X-Ray** - Application tracing and debugging

## Availability Zone (AZ) Specific Resources

These resources are tied to a specific availability zone within a region:

- **EC2 Instances** - Launched in a specific AZ
- **EBS Volumes** - Exist in a single AZ (can snapshot and copy across AZs/regions)
- **Subnets** - Each subnet exists in exactly one AZ
- **RDS Instances** - Primary instance in one AZ (can have standby in another for Multi-AZ)
- **ElastiCache Nodes** - Individual cache nodes in specific AZs
- **Redshift Nodes** - Cluster nodes in a single AZ

## Special Cases & Important Notes

### Multi-AZ Deployments

These services can be configured for high availability across multiple AZs within a region:

- **RDS Multi-AZ** - Synchronous replication to standby in another AZ
- **EFS** - Automatically stores data redundantly across multiple AZs
- **S3** - Standard storage class replicates across multiple AZs
- **ELB** (All types) - Can distribute traffic across multiple AZs
- **Aurora** - Replicates data across 3 AZs automatically
- **DynamoDB** - Replicates across 3 AZs in a region

### Cross-Region Capabilities

- **S3 Cross-Region Replication (CRR)** - Replicate objects to another region
- **DynamoDB Global Tables** - Multi-region, multi-active database
- **Aurora Global Database** - Cross-region replication with <1 second latency
- **RDS Cross-Region Read Replicas** - Asynchronous replication across regions
- **CloudFormation StackSets** - Deploy stacks across multiple regions
- **Route 53** - Can route to resources in any region
- **ECR** - Container images can be replicated across regions

### Name Uniqueness Scopes

- **S3 bucket names** - Must be globally unique
- **DynamoDB table names** - Unique within a region
- **RDS instance identifiers** - Unique within a region
- **Lambda function names** - Unique within a region
- **IAM user names** - Globally unique within an AWS account
- **VPC names** - Unique within a region (actually tags, VPC IDs are unique)

## Exam Tips

1. **IAM is global** - Users/roles/policies are not region-specific
2. **S3 buckets are regional** - Despite having globally unique names
3. **EBS volumes cannot move between AZs** - Must snapshot and restore
4. **Most services are regional** - When in doubt, assume regional
5. **Understand Multi-AZ vs Multi-Region** - Multi-AZ is for HA, Multi-Region is for DR
6. **Route 53 is global** - Can route to any regional resource
7. **CloudFront is global** - But origin can be regional
8. **Know which resources are AZ-locked** - EC2 instances, EBS, subnets
