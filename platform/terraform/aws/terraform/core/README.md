# Terraform AWS Infrastructure

This repository contains Terraform code for provisioning AWS infrastructure across multiple environments (dev, preprod, prod) using best practices.

## Project Structure

```
terraform/
├── backend.tf                  # Backend configuration
├── versions.tf                 # Terraform and provider version constraints
├── provider.tf                 # AWS provider configuration
├── main.tf                     # Root module - orchestrates all resources
├── variables.tf                # Input variables
├── outputs.tf                  # Output values
├── environments/               # Environment-specific configurations
│   ├── dev/
│   │   ├── backend.hcl        # Dev backend config
│   │   └── terraform.tfvars   # Dev variable values
│   ├── preprod/
│   │   ├── backend.hcl        # PreProd backend config
│   │   └── terraform.tfvars   # PreProd variable values
│   └── prod/
│       ├── backend.hcl        # Prod backend config
│       └── terraform.tfvars   # Prod variable values
└── modules/                    # Reusable modules
    ├── vpc/                   # VPC, subnets, NAT, IGW
    ├── ec2/                   # EC2 instances with security groups
    ├── rds/                   # RDS PostgreSQL with encryption
    └── s3/                    # S3 buckets with best practices
```

## Resources Provisioned

### VPC Module
- VPC with DNS support
- Public and private subnets across multiple AZs
- Internet Gateway for public subnets
- NAT Gateway for private subnet internet access
- Route tables and associations
- VPC Flow Logs to CloudWatch

### EC2 Module
- EC2 instances with latest Amazon Linux 2023
- Security groups with SSH, HTTP, HTTPS access
- IAM role with SSM access for management
- CloudWatch agent for monitoring
- IMDSv2 enforced for enhanced security
- EBS encryption enabled

### RDS Module
- PostgreSQL RDS instance
- DB subnet group across private subnets
- Security group restricting access to VPC
- Automated backups with configurable retention
- Enhanced monitoring
- Performance Insights enabled
- Encryption at rest
- Secrets Manager integration for credentials
- Multi-AZ support for production

### S3 Module
- Multiple S3 buckets (app-data, logs, backups)
- Versioning enabled
- Server-side encryption (AES256 or KMS)
- Public access blocked
- Lifecycle policies for cost optimization
- Access logging
- TLS-only bucket policies

## Prerequisites

1. **AWS CLI** configured with appropriate credentials
2. **Terraform** >= 1.5.0
3. **AWS Account** with necessary permissions
4. **S3 Bucket** and **DynamoDB Table** for Terraform state (see Setup below)

## Initial Setup

### 1. Create Backend Infrastructure

Before using Terraform, you need to create the S3 buckets and DynamoDB tables for state management. Run the setup script:

```bash
cd terraform
./scripts/setup-backend.sh <environment>
```

Where `<environment>` is one of: `dev`, `preprod`, or `prod`

This creates:
- S3 bucket for Terraform state
- DynamoDB table for state locking
- Enables versioning and encryption

### 2. Update Configuration

Edit the environment-specific files to match your requirements:
- `environments/<env>/backend.hcl` - Backend bucket names
- `environments/<env>/terraform.tfvars` - Resource configurations

### 3. Set Sensitive Variables

Set the RDS password via environment variable:

```bash
export TF_VAR_rds_password="your-secure-password"
```

Or use a `.tfvars` file (not committed to git):

```bash
echo 'rds_password = "your-secure-password"' > secrets.tfvars
```

## Usage

### Initialize Terraform

Initialize with environment-specific backend:

```bash
cd terraform
terraform init -backend-config=environments/dev/backend.hcl
```

### Plan Infrastructure

Preview changes:

```bash
terraform plan -var-file=environments/dev/terraform.tfvars
```

### Apply Infrastructure

Deploy resources:

```bash
terraform apply -var-file=environments/dev/terraform.tfvars
```

### Destroy Infrastructure

Remove all resources:

```bash
terraform destroy -var-file=environments/dev/terraform.tfvars
```

## Environment-Specific Deployments

### Development

```bash
# Initialize
terraform init -backend-config=environments/dev/backend.hcl

# Plan
terraform plan -var-file=environments/dev/terraform.tfvars

# Apply
terraform apply -var-file=environments/dev/terraform.tfvars
```

### Pre-Production

```bash
# Initialize
terraform init -backend-config=environments/preprod/backend.hcl -reconfigure

# Plan
terraform plan -var-file=environments/preprod/terraform.tfvars

# Apply
terraform apply -var-file=environments/preprod/terraform.tfvars
```

### Production

```bash
# Initialize
terraform init -backend-config=environments/prod/backend.hcl -reconfigure

# Plan
terraform plan -var-file=environments/prod/terraform.tfvars

# Apply (with extra confirmation)
terraform apply -var-file=environments/prod/terraform.tfvars
```

## Best Practices Implemented

### Security
- ✅ All data encrypted at rest (EBS, RDS, S3)
- ✅ IMDSv2 required for EC2 instances
- ✅ Security groups with least privilege
- ✅ Private subnets for databases
- ✅ S3 public access blocked
- ✅ TLS-only policies
- ✅ Secrets stored in AWS Secrets Manager
- ✅ VPC Flow Logs enabled

### High Availability
- ✅ Multi-AZ deployments for production
- ✅ Resources distributed across availability zones
- ✅ Auto-scaling capable architecture
- ✅ Automated backups for databases

### Monitoring & Observability
- ✅ CloudWatch agent on EC2 instances
- ✅ RDS Enhanced Monitoring
- ✅ RDS Performance Insights
- ✅ VPC Flow Logs
- ✅ S3 access logging

### Cost Optimization
- ✅ Lifecycle policies for S3 objects
- ✅ gp3 volumes for better price/performance
- ✅ Appropriately sized instances per environment
- ✅ Auto-scaling capabilities

### Infrastructure as Code
- ✅ DRY principle with reusable modules
- ✅ Environment-specific configurations
- ✅ Remote state with locking
- ✅ Version-controlled infrastructure
- ✅ Consistent tagging strategy

## Outputs

After applying, Terraform will output:
- VPC and subnet IDs
- EC2 instance IDs and IP addresses
- RDS endpoint (sensitive)
- S3 bucket names and ARNs

View outputs:

```bash
terraform output
```

View sensitive outputs:

```bash
terraform output -json | jq
```

## Accessing Resources

### EC2 Instances

Use SSM Session Manager (no SSH keys needed):

```bash
aws ssm start-session --target <instance-id>
```

Or traditional SSH:

```bash
ssh -i your-key.pem ec2-user@<public-ip>
```

### RDS Database

Retrieve credentials from Secrets Manager:

```bash
aws secretsmanager get-secret-value --secret-id myproject-dev-db-password --query SecretString --output text | jq
```

Connect from EC2 instance:

```bash
psql -h <rds-endpoint> -U admin -d appdb
```

## Troubleshooting

### State Lock Issues

If state is locked:

```bash
terraform force-unlock <lock-id>
```

### Backend Reconfiguration

When switching environments:

```bash
terraform init -backend-config=environments/<env>/backend.hcl -reconfigure
```

### Validate Configuration

```bash
terraform validate
```

### Format Code

```bash
terraform fmt -recursive
```

## Customization

### Adding New Resources

1. Create new module in `modules/` directory
2. Reference module in `main.tf`
3. Add required variables to `variables.tf`
4. Update environment-specific tfvars files

### Modifying Existing Resources

1. Update module configuration
2. Run `terraform plan` to preview changes
3. Apply changes with `terraform apply`

## Security Considerations

1. **Never commit sensitive data** - Use environment variables or AWS Secrets Manager
2. **Review IAM permissions** - Follow principle of least privilege
3. **Enable MFA Delete** on S3 state buckets for production
4. **Rotate credentials regularly**
5. **Use VPN or bastion hosts** for accessing private resources
6. **Enable CloudTrail** for API audit logging
7. **Set up AWS Config** for compliance monitoring

## Maintenance

### State File Backup

The S3 backend automatically versions state files, but consider:

```bash
aws s3 sync s3://myproject-terraform-state-prod s3://myproject-terraform-state-backup-prod
```

### Terraform Upgrades

Before upgrading Terraform:

```bash
terraform version
terraform init -upgrade
terraform plan
```

## Cost Estimation

Use AWS Cost Calculator or Infracost:

```bash
# Install infracost
brew install infracost

# Generate cost estimate
infracost breakdown --path .
```

## Support

For issues or questions:
1. Check Terraform documentation: https://www.terraform.io/docs
2. AWS documentation: https://docs.aws.amazon.com
3. Review CloudWatch logs for runtime issues

## License

[Your License Here]

## Contributing

1. Create feature branch
2. Make changes
3. Test in dev environment
4. Submit pull request
5. Deploy to preprod for validation
6. Deploy to production after approval
