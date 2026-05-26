# Makefile for Terraform operations

.PHONY: help init-dev init-preprod init-prod plan-dev plan-preprod plan-prod apply-dev apply-preprod apply-prod destroy-dev destroy-preprod destroy-prod fmt validate clean

# Default environment
ENV ?= dev

help:
	@echo "Terraform Infrastructure Management"
	@echo ""
	@echo "Usage:"
	@echo "  make init-<env>      Initialize Terraform for environment"
	@echo "  make plan-<env>      Plan changes for environment"
	@echo "  make apply-<env>     Apply changes for environment"
	@echo "  make destroy-<env>   Destroy infrastructure for environment"
	@echo "  make fmt             Format Terraform files"
	@echo "  make validate        Validate Terraform configuration"
	@echo ""
	@echo "Environments: dev, preprod, prod"
	@echo ""
	@echo "Examples:"
	@echo "  make init-dev"
	@echo "  make plan-dev"
	@echo "  make apply-prod"

# Development
init-dev:
	terraform init -backend-config=environments/dev/backend.hcl

plan-dev:
	terraform plan -var-file=environments/dev/terraform.tfvars

apply-dev:
	terraform apply -var-file=environments/dev/terraform.tfvars

destroy-dev:
	terraform destroy -var-file=environments/dev/terraform.tfvars

# Pre-Production
init-preprod:
	terraform init -backend-config=environments/preprod/backend.hcl -reconfigure

plan-preprod:
	terraform plan -var-file=environments/preprod/terraform.tfvars

apply-preprod:
	terraform apply -var-file=environments/preprod/terraform.tfvars

destroy-preprod:
	terraform destroy -var-file=environments/preprod/terraform.tfvars

# Production
init-prod:
	terraform init -backend-config=environments/prod/backend.hcl -reconfigure

plan-prod:
	terraform plan -var-file=environments/prod/terraform.tfvars

apply-prod:
	@echo "⚠️  WARNING: You are about to apply changes to PRODUCTION"
	@read -p "Type 'yes' to continue: " confirm && [ "$$confirm" = "yes" ]
	terraform apply -var-file=environments/prod/terraform.tfvars

destroy-prod:
	@echo "⚠️  WARNING: You are about to DESTROY PRODUCTION infrastructure"
	@read -p "Type 'destroy-prod' to continue: " confirm && [ "$$confirm" = "destroy-prod" ]
	terraform destroy -var-file=environments/prod/terraform.tfvars

# Common operations
fmt:
	terraform fmt -recursive

validate:
	terraform validate

clean:
	rm -rf .terraform
	rm -f .terraform.lock.hcl
	rm -f terraform.tfstate*

# Setup backend
setup-backend-dev:
	./scripts/setup-backend.sh dev

setup-backend-preprod:
	./scripts/setup-backend.sh preprod

setup-backend-prod:
	./scripts/setup-backend.sh prod
