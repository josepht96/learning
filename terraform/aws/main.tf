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

resource "aws_s3_bucket" "primary" {
  bucket = "primary-${var.project}"

  tags = {
    Name        = "Primary bucket"
    Environment = "Dev"
  }
}