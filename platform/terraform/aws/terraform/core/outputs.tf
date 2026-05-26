# Outputs

# VPC Outputs
output "vpc_id" {
  description = "ID of the VPC"
  value       = var.create_vpc ? module.vpc[0].vpc_id : null
}

output "public_subnet_ids" {
  description = "IDs of public subnets"
  value       = var.create_vpc ? module.vpc[0].public_subnet_ids : []
}

output "private_subnet_ids" {
  description = "IDs of private subnets"
  value       = var.create_vpc ? module.vpc[0].private_subnet_ids : []
}

# EC2 Outputs
output "ec2_instance_ids" {
  description = "IDs of EC2 instances"
  value       = var.create_ec2 ? module.ec2[0].instance_ids : []
}

output "ec2_public_ips" {
  description = "Public IPs of EC2 instances"
  value       = var.create_ec2 ? module.ec2[0].public_ips : []
}

# RDS Outputs
output "rds_endpoint" {
  description = "RDS instance endpoint"
  value       = var.create_rds ? module.rds[0].db_instance_endpoint : null
  sensitive   = true
}

output "rds_database_name" {
  description = "Name of the database"
  value       = var.create_rds ? module.rds[0].db_name : null
}

# S3 Outputs
output "s3_bucket_ids" {
  description = "IDs of created S3 buckets"
  value       = var.create_s3_buckets ? module.s3[*].bucket_ids : []
}

output "s3_bucket_arns" {
  description = "ARNs of created S3 buckets"
  value       = var.create_s3_buckets ? module.s3[*].bucket_arns : []
}
