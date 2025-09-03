terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  required_version = ">= 1.2.0"
}

provider "aws" {
  region  = var.location
}

module "my_s3_bucket" {
  source = "./modules/s3"
  project = "josepht96"
}