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

