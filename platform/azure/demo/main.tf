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
    type = string
    default = "eastus"
}
variable "resource_group_name" {
    type = string
    default = "eastus"
}
variable "tags" {
  type = map

  default = {
    Environment = "Terraform GS"
    Dept        = "Engineering"
    Team = "DevOps"
  }
}

resource "azurerm_resource_group" "rg" {
  name     = var.resource_group_name
  location = var.location

  tags = {
    Environment = lookup(var.tags, "Environment")
    Team        = lookup(var.tags, "Team")
  }
}


data "azurerm_resource_group" "resource-group" {
  name       = var.resource_group_name
  depends_on = [azurerm_resource_group.rg]
}

# Create a virtual network
resource "azurerm_virtual_network" "vnet" {
  name                = "myTFVnet"
  address_space       = ["10.0.0.0/16"]
  location            = var.location
  resource_group_name = azurerm_resource_group.rg.name
}