locals {
  common_tags = {
    Environment = var.environment
    Project     = var.project_name
    ManagedBy   = "Terraform"
  }

  name_prefix = "${var.project_name}-${var.environment}"
}

data "terraform_remote_state" "common" {
  backend = "s3"
  config = {
    bucket = "terraform-state-000001"
    key    = "dev/terraform.tfstate"
    region = var.aws_region
  }
}

# module "data_bucket" {
#   source      = "../core/modules/s3"

#   environment   = var.environment
#   bucket_names  = var.s3_bucket_names
#   name_prefix   = local.name_prefix

#   tags = local.common_tags
# }