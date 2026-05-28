# Domain 1: Cloud Concepts (24%)

## Task Statement 1.1: Define the benefits of the AWS Cloud

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

## Task Statement 1.2: Identify design principles of the AWS Cloud

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

## Task Statement 1.3: Understand the benefits of and strategies for migration to the AWS Cloud

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

> **COMMON USAGE:** Most organizations start with **Rehost** (lift-and-shift) as it's the fastest migration path with lowest risk. Over time, they gradually move to **Replatform** and **Refactor** for specific workloads to take advantage of cloud-native features. **Retire** often yields quick cost savings by identifying unused applications.

**Migration Tools:**

- AWS Snowball family (for large data transfers)
- Database Migration Service (DMS)
- AWS Application Migration Service

---

## Task Statement 1.4: Understand concepts of cloud economics

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

> **COMMON USAGE:** Managed services like **RDS, DynamoDB, and Lambda** are heavily favored over self-managed alternatives (e.g., databases on EC2) because they eliminate patching, backups, and scaling complexity. Organizations typically see 30-50% operational cost savings by using managed services.
