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

