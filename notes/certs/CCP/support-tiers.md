# AWS Support Tiers

## Overview

AWS offers five support tiers to meet different customer needs, ranging from basic documentation access to 24/7 technical support with dedicated Technical Account Managers (TAMs).

## Support Tier Comparison

### 1. Basic Support

**Cost**: Free (included with all AWS accounts)

**Features**:

- Customer Service & Communities access 24/7
- Access to documentation, whitepapers, and support forums
- AWS Trusted Advisor (7 core checks):
  - S3 bucket permissions
  - Security groups - specific ports unrestricted
  - IAM use
  - MFA on root account
  - EBS public snapshots
  - RDS public snapshots
  - Service limits
- AWS Personal Health Dashboard
- No technical support cases

**Best For**: Experimenting with AWS, non-production workloads

---

### 2. Developer Support

**Cost**: Greater of $29/month or 3% of monthly AWS usage

**Features**:

- All Basic Support features
- Business hours email access to Cloud Support Associates
- Unlimited support cases (1 primary contact)
- Response times:
  - General guidance: < 24 hours
  - System impaired: < 12 hours
- No third-party software support

**Best For**: Development and testing environments, early-stage startups

---

### 3. Business Support

**Cost**: Greater of $100/month or tiered pricing:

- 10% of monthly usage for first $0-$10K
- 7% for $10K-$80K
- 5% for $80K-$250K
- 3% over $250K

**Features**:

- All Developer Support features
- 24/7 phone, email, and chat access to Cloud Support Engineers
- Unlimited contacts/unlimited cases
- Full Trusted Advisor checks (all best practice recommendations)
- AWS Support API access
- Response times:
  - General guidance: < 24 hours
  - System impaired: < 12 hours
  - Production system impaired: < 4 hours
  - Production system down: < 1 hour
- Third-party software support (interoperability & configuration)
- Infrastructure Event Management (additional fee)

**Best For**: Production workloads, multiple services in use, need 24/7 support

---

### 4. Enterprise On-Ramp Support

**Cost**: Greater of $5,500/month or tiered pricing:

- 10% for first $0-$150K
- 7% for $150K-$500K
- 5% for $500K-$1M
- 3% over $1M

**Features**:

- All Business Support features
- Pool of Technical Account Managers (TAMs)
- Concierge Support Team (billing and account assistance)
- Infrastructure Event Management, Well-Architected & Operations Reviews
- Response times:
  - Business-critical system down: < 30 minutes
  - All other response times same as Business
- Access to AWS Incident Detection and Response (additional fee)
- Consultative architectural guidance
- Cost optimization workshops

**Best For**: Production and/or business-critical workloads, growing companies

---

### 5. Enterprise Support

**Cost**: Greater of $15,000/month or tiered pricing:

- 10% for first $0-$150K
- 7% for $150K-$500K
- 5% for $500K-$1M
- 3% over $1M

**Features**:

- All Enterprise On-Ramp features
- Dedicated Technical Account Manager (TAM)
- Application architecture guidance
- White-glove case routing
- Management business reviews
- Response times:
  - Business-critical system down: < 15 minutes
  - All other response times same as Business
- Operations reviews and tools recommendations
- Proactive guidance and best practices
- Training and game days
- Support automation workflows

**Best For**: Mission-critical workloads, enterprise-scale operations

---

## Key Differences Summary

| Feature | Basic | Developer | Business | Enterprise On-Ramp | Enterprise |
|---------|-------|-----------|----------|-------------------|------------|
| **Cost** | Free | $29+/mo | $100+/mo | $5,500+/mo | $15,000+/mo |
| **Technical Support** | None | Email only | 24/7 Phone/Chat/Email | 24/7 Phone/Chat/Email | 24/7 Phone/Chat/Email |
| **Support Cases** | No | Unlimited (1 contact) | Unlimited | Unlimited | Unlimited |
| **Trusted Advisor** | 7 checks | 7 checks | All checks | All checks | All checks |
| **TAM** | No | No | No | Pool of TAMs | Dedicated TAM |
| **Critical Response** | N/A | N/A | < 1 hour | < 30 minutes | < 15 minutes |
| **Third-party Support** | No | No | Yes | Yes | Yes |

---

## Response Time SLAs (CCP Exam Focus)

### General Guidance

- Developer: < 24 hours
- Business/Enterprise: < 24 hours

### System Impaired

- Developer: < 12 hours
- Business/Enterprise: < 12 hours

### Production System Impaired

- Business/Enterprise: < 4 hours

### Production System Down

- Business/Enterprise: < 1 hour

### Business-Critical System Down

- Enterprise On-Ramp: < 30 minutes
- Enterprise: < 15 minutes

---

## AWS Trusted Advisor

Available checks vary by support tier:

**Basic/Developer (7 Core Checks)**:

- Cost Optimization: None at this tier
- Performance: None at this tier
- Security: Basic security checks
- Fault Tolerance: Basic checks
- Service Limits: All service limit checks

**Business/Enterprise (All Checks)**:

- 50+ checks across all categories
- Cost optimization recommendations
- Performance improvements
- Security best practices
- Fault tolerance enhancements
- Service limit monitoring

---

## AWS Personal Health Dashboard

- Available to ALL support tiers
- Provides alerts and remediation guidance when AWS experiences events that may impact you
- Personalized view of AWS service health
- Proactive notifications about scheduled activities

---

## Exam Tips

1. **Know the response times** - This is frequently tested
2. **Understand TAM availability** - Only Enterprise On-Ramp (pool) and Enterprise (dedicated)
3. **Trusted Advisor differences** - Basic/Developer get 7 checks, Business+ get all checks
4. **Basic is free** - All AWS accounts include Basic support
5. **Business is the minimum for 24/7 support** - Developer is business hours only
6. **Third-party software support** - Requires Business or higher
7. **Critical response times**:
   - Business: 1 hour for production down
   - Enterprise On-Ramp: 30 minutes for business-critical
   - Enterprise: 15 minutes for business-critical
