# AWS Migration Acceleration Program (MAP)

## Overview
AWS Migration Acceleration Program (MAP) is a comprehensive cloud migration program that provides a proven methodology, tools, expertise, and financial support to accelerate your migration to AWS.

## Three-Phase Migration Journey

### Phase 1: Assess
**Goal:** Develop migration business case and readiness assessment

**Key Activities:**
- Evaluate current state of IT portfolio
- Identify migration readiness gaps
- Analyze application portfolio
- Calculate TCO (Total Cost of Ownership)
- Build migration business case
- Perform readiness assessment across 6 perspectives (aligned with CAF)

**Deliverables:**
- Migration Readiness Assessment (MRA)
- Migration business case
- TCO analysis
- Initial migration portfolio assessment

**Key AWS Tools:**
- AWS Migration Evaluator (formerly TSO Logic) - TCO analysis
- AWS Application Discovery Service - Inventory and dependency mapping
- Migration Hub - Track migration progress

**Duration:** Typically 2-4 weeks

---

### Phase 2: Mobilize
**Goal:** Build operational foundation for migration at scale

**Key Activities:**
- Address gaps identified in Assess phase
- Create detailed migration plan
- Build cloud foundation (landing zone)
- Establish operating model
- Develop skills and processes
- Run pilot migrations
- Set up security baselines
- Configure migration tools

**Deliverables:**
- Migration landing zone
- Trained migration team
- Security and compliance framework
- Migration playbooks and runbooks
- Pilot migration results
- Detailed wave planning

**Key AWS Tools:**
- AWS Landing Zone / Control Tower - Multi-account environment
- AWS Service Catalog - Standardized provisioning
- AWS Migration Hub - Orchestration
- CloudEndure Migration / AWS MGN - Rehost migrations
- AWS Database Migration Service (DMS) - Database migrations

**Duration:** Typically 2-3 months

**Success Criteria:**
- Landing zone operational
- Pilot migrations completed successfully
- Team trained and ready
- Processes documented and tested

---

### Phase 3: Migrate and Modernize
**Goal:** Execute migration at scale and modernize workloads

**Key Activities:**
- Execute migration in waves
- Migrate applications using the 7 Rs strategy
- Optimize and modernize workloads
- Decommission on-premises infrastructure
- Iterate and refine processes
- Achieve targeted business outcomes

**Migration Strategies - The 7 Rs:**

1. **Retire**
   - Decommission applications no longer needed
   - Identify through portfolio analysis
   - Reduces TCO immediately
   - No migration effort required

2. **Retain**
   - Keep applications on-premises (for now)
   - Due to compliance, performance, or other constraints
   - Revisit decision periodically
   - Focus resources on higher-value migrations

3. **Rehost** (Lift and Shift)
   - Move applications as-is to AWS
   - Minimal or no changes
   - Fastest migration path
   - Good for: legacy apps, quick wins, data center exit
   - Tools: AWS MGN, CloudEndure

4. **Relocate** (Hypervisor-level lift and shift)
   - Move VMware workloads to VMware Cloud on AWS
   - No conversion or modifications
   - Specific to VMware environments
   - Tools: VMware Cloud on AWS

5. **Repurchase** (Drop and Shop)
   - Move to different product (usually SaaS)
   - Example: Migrate CRM to Salesforce, email to Office 365
   - May require data migration
   - Reduces operational burden

6. **Replatform** (Lift, Tinker, and Shift)
   - Make a few cloud optimizations
   - Don't change core architecture
   - Examples: Move to RDS, use ELB
   - Balance between speed and optimization
   - Tools: AWS Elastic Beanstalk, RDS

7. **Refactor** / Re-architect
   - Reimagine application architecture
   - Built with cloud-native features
   - Driven by business need for scale, performance, features
   - Most expensive but highest long-term value
   - Examples: Monolith to microservices, serverless
   - Tools: Lambda, containers, API Gateway

**Key AWS Services:**
- AWS Application Migration Service (MGN) - Rehost
- AWS Database Migration Service (DMS) - Database migrations
- AWS DataSync - Data transfer
- AWS Snow Family - Large-scale data transfer
- AWS Server Migration Service (SMS) - Legacy, replaced by MGN
- Migration Hub - Centralized tracking

**Best Practices:**
- Migrate in waves (not all at once)
- Start with less complex applications
- Build confidence and expertise
- Use automation wherever possible
- Monitor and optimize continuously
- Celebrate wins to maintain momentum

---

## Migration Strategies Decision Matrix

| Strategy | Speed | Cost Savings | Complexity | When to Use |
|----------|-------|--------------|------------|-------------|
| Retire | Fastest | Highest | Lowest | Unused/redundant apps |
| Retain | N/A | None | None | Compliance, not ready |
| Rehost | Fast | Moderate | Low | Quick wins, DC exit |
| Relocate | Fast | Moderate | Low | VMware environments |
| Repurchase | Medium | Variable | Medium | Legacy licensing, SaaS available |
| Replatform | Medium | Good | Medium | Optimize without re-architect |
| Refactor | Slow | Best (long-term) | High | Need cloud-native benefits |

---

## Key Migration Patterns

### Database Migration Patterns
- **Homogeneous:** Same DB engine (Oracle to Oracle on RDS)
- **Heterogeneous:** Different DB engine (Oracle to Aurora PostgreSQL)
- **Schema Conversion Tool (SCT):** Convert schemas for heterogeneous migrations
- **DMS:** Replicate data with minimal downtime
- **Continuous replication:** Keep databases in sync during migration

### Application Migration Patterns
- **Big Bang:** Migrate everything at once (high risk)
- **Phased:** Migrate in planned waves (recommended)
- **Parallel Running:** Run both old and new systems
- **Strangler Fig:** Gradually replace pieces of the application

---

## Migration Factory Approach

**Concept:** Industrialized, repeatable process for migrating at scale

**Key Elements:**
- Standardized playbooks and runbooks
- Automation and tooling
- Dedicated migration teams
- Wave-based planning
- Centralized tracking and governance
- Continuous improvement

**Benefits:**
- Faster migrations
- Predictable outcomes
- Lower costs
- Risk reduction
- Knowledge retention

---

## AWS MAP Funding

AWS provides financial support through MAP:
- **Migration credits** for AWS services
- **Professional services funding** for AWS partners
- **Training credits** for skill development
- Requires business case and commitment

**Typical Structure:**
- Funding based on migration commitment
- Released in phases as milestones achieved
- Must be used within specified timeframe

---

## Critical Success Factors

1. **Executive Sponsorship:** C-level support and commitment
2. **Clear Business Case:** Articulated benefits and ROI
3. **Migration Team:** Dedicated resources with right skills
4. **Partner Engagement:** Leverage AWS partners when needed
5. **Landing Zone:** Secure, scalable foundation
6. **Automation:** Tools and scripts for repeatability
7. **Change Management:** Address people and process
8. **Security and Compliance:** Built-in from the start

---

## Common Migration Challenges

| Challenge | Solution |
|-----------|----------|
| Lack of cloud skills | Training, hiring, partners |
| Legacy dependencies | Modernize or containerize |
| Compliance concerns | AWS compliance programs, security controls |
| Downtime constraints | Use live migration tools (DMS, MGN) |
| Data transfer | AWS DataSync, Snow Family, Direct Connect |
| Application discovery | AWS Application Discovery Service |
| Cost uncertainty | Migration Evaluator, budgets, monitoring |

---

## Migration Tools Summary

| Tool | Purpose | Use Case |
|------|---------|----------|
| Migration Evaluator | TCO analysis | Business case development |
| Application Discovery Service | Inventory and dependencies | Portfolio analysis |
| Migration Hub | Central tracking | Orchestration and visibility |
| Application Migration Service (MGN) | Automated rehost | Lift-and-shift migrations |
| Database Migration Service (DMS) | Database replication | Database migrations |
| DataSync | Automated data transfer | Large data movements |
| Snow Family | Physical data transfer | Petabyte-scale transfers |
| Schema Conversion Tool (SCT) | Database schema conversion | Heterogeneous DB migrations |

---

## Post-Migration Optimization

After migration, focus on:
- **Right-sizing:** Adjust instance sizes
- **Reserved Instances/Savings Plans:** Lock in discounts
- **Modernization:** Refactor to cloud-native
- **Automation:** Implement Infrastructure as Code
- **Monitoring:** CloudWatch, Cost Explorer
- **Security:** Continuous compliance and hardening
- **Well-Architected Review:** Optimize across pillars

---

## Common Exam Topics

- Know the three phases: Assess, Mobilize, Migrate & Modernize
- Memorize the 7 Rs and when to use each
- Understand migration tools and their purposes
- Know the difference between homogeneous and heterogeneous migrations
- Understand migration factory approach
- Recognize scenarios for different migration strategies
- Know which tools to use for different migration types

---

## Quick Reference - The 7 Rs

**Remember:** Retired Relatives Rarely Relax, Resting Peacefully Requires Refactoring

1. **Retire** - Turn it off
2. **Retain** - Keep it on-prem (for now)
3. **Rehost** - Lift and shift
4. **Relocate** - VMware to VMware Cloud
5. **Repurchase** - Move to SaaS
6. **Replatform** - Lift, tinker, shift
7. **Refactor** - Rebuild for cloud
