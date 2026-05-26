terraform init -reconfigure -backend-config=environments/localstack/backend.hcl 
terraform plan -var-file=environments/localstack/terraform.tfvars
terraform apply -var-file=environments/localstack/terraform.tfvars
aws s3 rm s3://terraform-state-000001/dev/terraform.tfstate.tflock


# AWS CLI Commands for LocalStack

This document contains useful AWS CLI commands to interact with LocalStack resources.

## Setup

Set the LocalStack endpoint as an environment variable for easier command usage:

```bash
export AWS_ENDPOINT=http://localhost:4566
export AWS_DEFAULT_REGION=us-east-1
```

Or use the `--endpoint-url` flag with each command.

## S3 Commands

### List all S3 buckets
```bash
aws --endpoint-url=http://localhost:4566 s3 ls
```

### List contents of a specific bucket
```bash
aws --endpoint-url=http://localhost:4566 s3 ls s3://myproject-dev-app-data
aws --endpoint-url=http://localhost:4566 s3 ls s3://terraform-state
```

### Get bucket details
```bash
aws --endpoint-url=http://localhost:4566 s3api list-buckets
```

## VPC Commands

### List all VPCs
```bash
aws --endpoint-url=http://localhost:4566 ec2 describe-vpcs
```

### Get VPC details (formatted)
```bash
aws --endpoint-url=http://localhost:4566 ec2 describe-vpcs \
  --query 'Vpcs[*].[VpcId,CidrBlock,State,Tags[?Key==`Name`].Value|[0]]' \
  --output table
```

### List all subnets
```bash
aws --endpoint-url=http://localhost:4566 ec2 describe-subnets
```

### List subnets (formatted)
```bash
aws --endpoint-url=http://localhost:4566 ec2 describe-subnets \
  --query 'Subnets[*].[SubnetId,VpcId,CidrBlock,AvailabilityZone,Tags[?Key==`Name`].Value|[0]]' \
  --output table
```

### List Internet Gateways
```bash
aws --endpoint-url=http://localhost:4566 ec2 describe-internet-gateways
```

### List NAT Gateways
```bash
aws --endpoint-url=http://localhost:4566 ec2 describe-nat-gateways
```

### List Route Tables
```bash
aws --endpoint-url=http://localhost:4566 ec2 describe-route-tables
```

### List Security Groups
```bash
aws --endpoint-url=http://localhost:4566 ec2 describe-security-groups
```

## EC2 Commands

### List all EC2 instances
```bash
aws --endpoint-url=http://localhost:4566 ec2 describe-instances
```

### List EC2 instances (formatted)
```bash
aws --endpoint-url=http://localhost:4566 ec2 describe-instances \
  --query 'Reservations[*].Instances[*].[InstanceId,InstanceType,State.Name,PublicIpAddress,Tags[?Key==`Name`].Value|[0]]' \
  --output table
```

### Get instance status
```bash
aws --endpoint-url=http://localhost:4566 ec2 describe-instance-status
```

## RDS Commands

### List all RDS instances
```bash
aws --endpoint-url=http://localhost:4566 rds describe-db-instances
```

### List RDS instances (formatted)
```bash
aws --endpoint-url=http://localhost:4566 rds describe-db-instances \
  --query 'DBInstances[*].[DBInstanceIdentifier,DBInstanceClass,Engine,DBInstanceStatus,Endpoint.Address]' \
  --output table
```

### List RDS subnet groups
```bash
aws --endpoint-url=http://localhost:4566 rds describe-db-subnet-groups
```

## General Resource Listing

### List all resources (using ResourceGroupsTaggingAPI)
```bash
aws --endpoint-url=http://localhost:4566 resourcegroupstaggingapi get-resources
```

### List resources by tag
```bash
aws --endpoint-url=http://localhost:4566 resourcegroupstaggingapi get-resources \
  --tag-filters Key=Environment,Values=dev
```

## Useful Aliases

Add these to your `~/.bashrc` or `~/.zshrc`:

```bash
# LocalStack alias
alias awslocal='aws --endpoint-url=http://localhost:4566'

# Usage examples:
# awslocal s3 ls
# awslocal ec2 describe-vpcs
# awslocal rds describe-db-instances
```

## Quick Check Script

```bash
#!/bin/bash
# Quick check of all LocalStack resources

ENDPOINT="http://localhost:4566"

echo "=== S3 Buckets ==="
aws --endpoint-url=$ENDPOINT s3 ls

echo -e "\n=== VPCs ==="
aws --endpoint-url=$ENDPOINT ec2 describe-vpcs --query 'Vpcs[*].[VpcId,CidrBlock]' --output table

echo -e "\n=== Subnets ==="
aws --endpoint-url=$ENDPOINT ec2 describe-subnets --query 'Subnets[*].[SubnetId,CidrBlock]' --output table

echo -e "\n=== EC2 Instances ==="
aws --endpoint-url=$ENDPOINT ec2 describe-instances --query 'Reservations[*].Instances[*].[InstanceId,State.Name]' --output table

echo -e "\n=== RDS Instances ==="
aws --endpoint-url=$ENDPOINT rds describe-db-instances --query 'DBInstances[*].[DBInstanceIdentifier,DBInstanceStatus]' --output table
```
