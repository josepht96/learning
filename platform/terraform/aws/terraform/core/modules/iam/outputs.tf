# ========================================
# Group Outputs
# ========================================

output "admins_group_arn" {
  description = "ARN of the admins IAM group"
  value       = aws_iam_group.admins.arn
}

output "admins_group_name" {
  description = "Name of the admins IAM group"
  value       = aws_iam_group.admins.name
}

output "developers_group_arn" {
  description = "ARN of the developers IAM group"
  value       = aws_iam_group.developers.arn
}

output "developers_group_name" {
  description = "Name of the developers IAM group"
  value       = aws_iam_group.developers.name
}

output "billing_group_arn" {
  description = "ARN of the billing IAM group"
  value       = aws_iam_group.billing.arn
}

output "billing_group_name" {
  description = "Name of the billing IAM group"
  value       = aws_iam_group.billing.name
}

# ========================================
# Role Outputs
# ========================================

output "ec2_s3_role_arn" {
  description = "ARN of the EC2 S3 access role"
  value       = aws_iam_role.ec2_s3_role.arn
}

output "ec2_s3_role_name" {
  description = "Name of the EC2 S3 access role"
  value       = aws_iam_role.ec2_s3_role.name
}

output "ec2_s3_instance_profile_arn" {
  description = "ARN of the EC2 S3 instance profile"
  value       = aws_iam_instance_profile.ec2_s3_profile.arn
}

output "ec2_s3_instance_profile_name" {
  description = "Name of the EC2 S3 instance profile"
  value       = aws_iam_instance_profile.ec2_s3_profile.name
}

output "ci_cd_role_arn" {
  description = "ARN of the CI/CD role"
  value       = aws_iam_role.ci_cd_role.arn
}

output "ci_cd_role_name" {
  description = "Name of the CI/CD role"
  value       = aws_iam_role.ci_cd_role.name
}

output "ci_cd_instance_profile_arn" {
  description = "ARN of the CI/CD instance profile"
  value       = aws_iam_instance_profile.ci_cd_profile.arn
}

output "ci_cd_instance_profile_name" {
  description = "Name of the CI/CD instance profile"
  value       = aws_iam_instance_profile.ci_cd_profile.name
}

# ========================================
# User Outputs
# ========================================

output "admin_user_arns" {
  description = "ARNs of admin users"
  value       = { for k, v in aws_iam_user.admin_users : k => v.arn }
}

output "developer_user_arns" {
  description = "ARNs of developer users"
  value       = { for k, v in aws_iam_user.developer_users : k => v.arn }
}

output "billing_user_arns" {
  description = "ARNs of billing users"
  value       = { for k, v in aws_iam_user.billing_users : k => v.arn }
}

# ========================================
# Policy Outputs
# ========================================

output "cost_explorer_policy_arn" {
  description = "ARN of the Cost Explorer policy"
  value       = aws_iam_policy.cost_explorer_policy.arn
}

output "password_policy_expire_passwords" {
  description = "Whether password expiration is enabled"
  value       = aws_iam_account_password_policy.strict.expire_passwords
}

# ========================================
# Summary Output
# ========================================

output "iam_setup_summary" {
  description = "Summary of IAM resources created"
  value = {
    groups = {
      admins     = aws_iam_group.admins.arn
      developers = aws_iam_group.developers.arn
      billing    = aws_iam_group.billing.arn
    }
    roles = {
      ec2_s3_role = aws_iam_role.ec2_s3_role.arn
      ci_cd_role  = aws_iam_role.ci_cd_role.arn
    }
    instance_profiles = {
      ec2_s3_profile = aws_iam_instance_profile.ec2_s3_profile.arn
      ci_cd_profile  = aws_iam_instance_profile.ci_cd_profile.arn
    }
    user_count = {
      admins     = length(var.admin_users)
      developers = length(var.developer_users)
      billing    = length(var.billing_users)
    }
  }
}
