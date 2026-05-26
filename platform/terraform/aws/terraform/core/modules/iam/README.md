# AWS IAM Terraform Module

This Terraform module creates a comprehensive AWS IAM setup including users, groups, roles, and policies following AWS best practices.

## Features

- **IAM Groups** with appropriate managed and custom policies
- **IAM Users** with console access (no programmatic access keys)
- **IAM Roles** for EC2 instances with specific permissions
- **Password Policy** enforcement at the account level
- **No hardcoded secrets** - passwords are set on first login

## Resources Created

### Groups

1. **admins** - Full AWS access via AdministratorAccess managed policy
2. **developers** - Read-only access plus EC2 and S3 management capabilities
3. **billing** - Access to billing information and Cost Explorer

### Users

1. **joe** - Member of admins group
2. **dev-user-1** - Member of developers group
3. **billing-user-1** - Member of billing group

All users are created with console access and must reset their password on first login.

### Roles

1. **ec2-s3-role** - For EC2 instances to access S3 (includes instance profile)
2. **ci-cd-role** - For CI/CD server to push to ECR and deploy to ECS (includes instance profile)

### Policies

- **Account Password Policy** - Enforces strong password requirements
- **Cost Explorer Policy** - Custom policy for billing group to access cost information

## Usage

### Basic Usage

```hcl
module "iam" {
  source = "./modules/iam"

  environment = "prod"
  aws_region  = "us-east-1"

  admin_users     = ["joe"]
  developer_users = ["dev-user-1"]
  billing_users   = ["billing-user-1"]

  tags = {
    Project = "MyProject"
    Team    = "Platform"
  }
}
```

### Custom Password Policy

```hcl
module "iam" {
  source = "./modules/iam"

  environment = "prod"
  aws_region  = "us-east-1"

  password_policy = {
    minimum_password_length        = 16
    require_lowercase_characters   = true
    require_uppercase_characters   = true
    require_numbers                = true
    require_symbols                = true
    allow_users_to_change_password = true
    max_password_age               = 60
    password_reuse_prevention      = 24
  }
}
```

### Multiple Users

```hcl
module "iam" {
  source = "./modules/iam"

  environment = "dev"

  admin_users     = ["joe", "admin-user-2"]
  developer_users = ["dev-user-1", "dev-user-2", "dev-user-3"]
  billing_users   = ["billing-user-1", "finance-admin"]
}
```

## Accessing Outputs

```hcl
# Group ARNs
output "admin_group" {
  value = module.iam.admins_group_arn
}

# Role ARNs
output "ec2_role" {
  value = module.iam.ec2_s3_role_arn
}

# Instance Profile for EC2
output "ec2_profile" {
  value = module.iam.ec2_s3_instance_profile_name
}

# Summary
output "iam_summary" {
  value = module.iam.iam_setup_summary
}
```

## Using Roles with EC2 Instances

### EC2 with S3 Access

```hcl
resource "aws_instance" "app_server" {
  ami                    = "ami-12345678"
  instance_type          = "t3.micro"
  iam_instance_profile   = module.iam.ec2_s3_instance_profile_name

  tags = {
    Name = "App Server"
  }
}
```

### CI/CD Server

```hcl
resource "aws_instance" "ci_cd_server" {
  ami                    = "ami-12345678"
  instance_type          = "t3.small"
  iam_instance_profile   = module.iam.ci_cd_instance_profile_name

  tags = {
    Name = "CI/CD Server"
  }
}
```

## User Password Setup

After Terraform creates the users:

1. **Root/Admin user** generates temporary passwords:
   ```bash
   aws iam create-login-profile --user-name joe --password TempPassword123! --password-reset-required
   ```

2. **Users login** to AWS Console using:
   - Account ID or alias
   - Username
   - Temporary password

3. **Users are prompted** to change password on first login

## Security Best Practices

This module implements several security best practices:

- No programmatic access keys created in Terraform
- Strong password policy enforcement
- Password rotation required (90 days)
- Password reuse prevention (12 previous passwords)
- Least privilege principle (groups have only necessary permissions)
- Use of AWS managed policies where available
- Custom inline policies only where necessary
- All resources tagged for tracking and cost allocation

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| aws_region | AWS region | `string` | `"us-east-1"` | no |
| environment | Environment name | `string` | `"dev"` | no |
| admin_users | List of admin users | `list(string)` | `["joe"]` | no |
| developer_users | List of developer users | `list(string)` | `["dev-user-1"]` | no |
| billing_users | List of billing users | `list(string)` | `["billing-user-1"]` | no |
| password_policy | Password policy configuration | `object` | See variables.tf | no |
| tags | Additional tags | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| admins_group_arn | ARN of admins group |
| developers_group_arn | ARN of developers group |
| billing_group_arn | ARN of billing group |
| ec2_s3_role_arn | ARN of EC2 S3 role |
| ec2_s3_instance_profile_name | Name of EC2 S3 instance profile |
| ci_cd_role_arn | ARN of CI/CD role |
| ci_cd_instance_profile_name | Name of CI/CD instance profile |
| iam_setup_summary | Complete summary of all IAM resources |

## Requirements

- Terraform >= 1.0
- AWS Provider >= 5.0
- AWS credentials with IAM administrative permissions

## Notes

- The password policy is applied at the AWS account level and affects all IAM users
- Login profiles are created with lifecycle ignore_changes to prevent issues after initial creation
- Instance profiles are automatically created for roles that need to be assumed by EC2
- The CI/CD role includes PassRole permission limited to ECS tasks only

## License

MIT

## Author

Created with Terraform best practices in mind.
