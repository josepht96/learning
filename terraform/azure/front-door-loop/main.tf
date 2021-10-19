# Configure the Azure provider
terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = ">= 2.26"
    }
  }
  required_version = ">= 0.14.9"
}

provider "azurerm" {
  features {}
}

variable "location" {
  type    = string
  default = "eastus"
}

variable "apps" {
  type = list(object({
    name  = string
    url   = string
    short = string
  }))
}

data "azurerm_resource_group" "resource-group" {
  name = "test-resource-group"
}

