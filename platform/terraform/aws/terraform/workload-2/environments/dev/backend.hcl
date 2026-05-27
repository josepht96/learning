# Backend configuration for dev environment
# Usage: terraform init -backend-config=environments/dev/backend.hcl

bucket         = "terraform-state-000001"
key            = "dev/terraform-workload-1.tfstate"
region         = "us-east-1"
encrypt        = true
use_lockfile   = true
