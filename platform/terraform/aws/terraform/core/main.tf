# Main Terraform configuration

locals {
  common_tags = {
    Environment = var.environment
    Project     = var.project_name
    ManagedBy   = "Terraform"
  }

  name_prefix = "${var.project_name}-${var.environment}"
}

# VPC Module
module "vpc" {
  source = "./modules/vpc"

  count = var.create_vpc ? 1 : 0

  environment          = var.environment
  vpc_cidr             = var.vpc_cidr
  availability_zones   = var.availability_zones
  public_subnet_cidrs  = var.public_subnet_cidrs
  private_subnet_cidrs = var.private_subnet_cidrs
  name_prefix          = local.name_prefix
  create_igw           = var.create_igw
  create_nat_gateway   = var.create_nat_gateway

  tags = local.common_tags
}

module "iam" {
  source = "./modules/iam"

  count = var.create_iam ? 1 : 0

  environment   = var.environment

  tags = local.common_tags
}


# EC2 Module
module "ec2" {
  source = "./modules/ec2"

  count = var.create_ec2 && var.create_vpc ? 1 : 0

  environment      = var.environment
  instance_type    = var.ec2_instance_type
  instance_count   = var.ec2_instance_count
  vpc_id           = module.vpc[0].vpc_id
  subnet_ids       = module.vpc[0].public_subnet_ids
  name_prefix      = local.name_prefix

  tags = local.common_tags

  depends_on = [module.vpc]
}

# RDS Module
module "rds" {
  source = "./modules/rds"

  count = var.create_rds && var.create_vpc ? 1 : 0

  environment        = var.environment
  instance_class     = var.rds_instance_class
  engine             = var.rds_engine
  engine_version     = var.rds_engine_version
  allocated_storage  = var.rds_allocated_storage
  database_name      = var.rds_database_name
  master_username    = var.rds_username
  master_password    = var.rds_password
  vpc_id             = module.vpc[0].vpc_id
  subnet_ids         = module.vpc[0].private_subnet_ids
  name_prefix        = local.name_prefix

  tags = local.common_tags

  depends_on = [module.vpc]
}

# S3 Module
module "s3" {
  source = "./modules/s3"

  count = var.create_s3_buckets ? 1 : 0

  environment   = var.environment
  bucket_names  = var.s3_bucket_names
  name_prefix   = local.name_prefix

  tags = local.common_tags
}
