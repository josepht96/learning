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

## Domain Breakdown

The exam content is divided into four domains. Each domain has been moved to a separate file:

- [Domain 1: Design Secure Architectures (30%)](domain_1.md)
- [Domain 2: Design Resilient Architectures (26%)](domain_2.md)
- [Domain 3: Design High-Performing Architectures (24%)](domain_3.md)
- [Domain 4: Design Cost-Optimized Architectures (20%)](domain_4.md)

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
