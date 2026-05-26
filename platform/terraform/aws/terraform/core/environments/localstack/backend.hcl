# Backend configuration for LocalStack
# Usage: terraform init -backend-config=environments/localstack/backend.hcl

bucket                      = "terraform-state"
key                         = "local/terraform.tfstate"
region                      = "us-east-1"
encrypt                     = true
use_lockfile                = true
skip_credentials_validation = true
skip_metadata_api_check     = true
skip_requesting_account_id  = true
use_path_style              = true
access_key                  = "test"
secret_key                  = "test"
endpoints = {
  s3 = "http://localhost:4566"
}
