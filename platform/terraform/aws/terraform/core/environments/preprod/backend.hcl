# Backend configuration for preprod environment
# Usage: terraform init -backend-config=environments/preprod/backend.hcl

bucket         = "myproject-terraform-state-preprod"
key            = "preprod/terraform.tfstate"
region         = "us-east-1"
encrypt        = true
use_lockfile   = true
