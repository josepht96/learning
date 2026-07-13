# Backend configuration for Terraform state
# This should be configured per environment

terraform {
  backend "s3" {
    # These values should be overridden via backend config file
    # Usage: terraform init -backend-config=environments/dev/backend.hcl
    bucket         = ""  # Will be set via backend config
    key            = "terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
    use_lockfile   = true  # Use S3 native locking instead of DynamoDB
  }
}
