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

