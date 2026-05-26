variable "aws_region" {
  description = "AWS region for resources"
  type        = string
  default     = "us-east-1"
}

variable "use_localstack" {
  description = "Use LocalStack instead of AWS"
  type        = bool
  default     = false
}

variable "localstack_endpoint" {
  description = "LocalStack endpoint URL"
  type        = string
  default     = "http://localhost:4566"
}

variable "environment" {
  description = "Environment name (dev, preprod, prod)"
  type        = string

  validation {
    condition     = contains(["dev", "preprod", "prod"], var.environment)
    error_message = "Environment must be dev, preprod, or prod."
  }
}

variable "project_name" {
  description = "Project name for resource naming and tagging"
  type        = string
  default     = "myproject"
}

variable "owner" {
  description = "Owner of the resources"
  type        = string
  default     = "devops-team"
}

# VPC Variables
variable "vpc_cidr" {
  description = "CIDR block for VPC"
  type        = string
  default     = "10.0.0.0/16"
}

variable "availability_zones" {
  description = "Availability zones for the region"
  type        = list(string)
  default     = ["us-east-1a", "us-east-1b"]
}

variable "public_subnet_cidrs" {
  description = "CIDR blocks for public subnets"
  type        = list(string)
  default     = ["10.0.1.0/24", "10.0.2.0/24"]
}

variable "private_subnet_cidrs" {
  description = "CIDR blocks for private subnets"
  type        = list(string)
  default     = ["10.0.10.0/24", "10.0.11.0/24"]
}

# EC2 Variables
variable "ec2_instance_type" {
  description = "EC2 instance type"
  type        = string
  default     = "t3.micro"
}

variable "ec2_instance_count" {
  description = "Number of EC2 instances"
  type        = number
  default     = 1
}

# RDS Variables
variable "rds_instance_class" {
  description = "RDS instance class"
  type        = string
  default     = "db.t3.micro"
}

variable "rds_engine" {
  description = "RDS engine type"
  type        = string
  default     = "postgres"
}

variable "rds_engine_version" {
  description = "RDS engine version"
  type        = string
  default     = "15.4"
}

variable "rds_allocated_storage" {
  description = "RDS allocated storage in GB"
  type        = number
  default     = 20
}

variable "rds_database_name" {
  description = "Name of the database to create"
  type        = string
  default     = "appdb"
}

variable "rds_username" {
  description = "Master username for RDS"
  type        = string
  default     = "admin"
  sensitive   = true
}

variable "rds_password" {
  description = "Master password for RDS"
  type        = string
  default     = "admin"
  sensitive   = true
}

# Module Toggle Variables
variable "create_vpc" {
  description = "Whether to create VPC module"
  type        = bool
  default     = true
}

variable "create_igw" {
  description = "Whether to create Internet Gateway"
  type        = bool
  default     = true
}

variable "create_nat_gateway" {
  description = "Whether to create NAT Gateway and EIP"
  type        = bool
  default     = true
}

variable "create_ec2" {
  description = "Whether to create EC2 module"
  type        = bool
  default     = true
}

variable "create_rds" {
  description = "Whether to create RDS module"
  type        = bool
  default     = true
}

# S3 Variables
variable "create_s3_buckets" {
  description = "Whether to create S3 buckets"
  type        = bool
  default     = true
}

variable "s3_bucket_names" {
  description = "List of S3 bucket names to create"
  type        = list(string)
  default     = ["app-data", "logs", "backups"]
}

# S3 Variables
variable "create_iam" {
  description = "Whether to create IAM config"
  type        = bool
  default     = true
}