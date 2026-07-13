# Example usage of the IAM module

terraform {
  required_version = ">= 1.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

# Basic IAM setup with default configuration
module "iam" {
  source = "../../"

  environment = "dev"
  aws_region  = "us-east-1"

  # Users can be customized
  admin_users     = ["joe"]
  developer_users = ["dev-user-1"]
  billing_users   = ["billing-user-1"]

  # Optional: Customize password policy
  password_policy = {
    minimum_password_length        = 14
    require_lowercase_characters   = true
    require_uppercase_characters   = true
    require_numbers                = true
    require_symbols                = true
    allow_users_to_change_password = true
    max_password_age               = 90
    password_reuse_prevention      = 12
  }

  tags = {
    Project     = "IAM-Setup"
    Team        = "Platform"
    CostCenter  = "Engineering"
    Terraform   = "true"
  }
}

# Example EC2 instance using the EC2-S3 role
resource "aws_instance" "example" {
  ami                    = "ami-0c55b159cbfafe1f0"  # Amazon Linux 2 (update for your region)
  instance_type          = "t3.micro"
  iam_instance_profile   = module.iam.ec2_s3_instance_profile_name

  tags = {
    Name = "Example Server with S3 Access"
  }
}

# Example outputs
output "iam_summary" {
  description = "Summary of all IAM resources created"
  value       = module.iam.iam_setup_summary
}

output "ec2_instance_profile" {
  description = "Instance profile name for EC2 instances"
  value       = module.iam.ec2_s3_instance_profile_name
}

output "ci_cd_instance_profile" {
  description = "Instance profile name for CI/CD server"
  value       = module.iam.ci_cd_instance_profile_name
}

output "all_group_arns" {
  description = "All IAM group ARNs"
  value = {
    admins     = module.iam.admins_group_arn
    developers = module.iam.developers_group_arn
    billing    = module.iam.billing_group_arn
  }
}

output "all_role_arns" {
  description = "All IAM role ARNs"
  value = {
    ec2_s3 = module.iam.ec2_s3_role_arn
    ci_cd  = module.iam.ci_cd_role_arn
  }
}
