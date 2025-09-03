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

data "azurerm_resource_group" "resource-group" {
  name = "default"
}
variable "vnets" {
  type = map(object({
    name = string
    tags = map(string)
  }))
}
variable "tags" {
  type = map(any)
  default = {
    default  = "default"
    default2 = "default2"
  }
}
resource "azurerm_virtual_network" "vnet" {
  for_each            = var.vnets
  name                = each.value.name
  address_space       = ["10.0.0.0/16"]
  location            = "eastus"
  resource_group_name = data.azurerm_resource_group.resource-group.name
  tags                = merge(var.tags, each.value.tags)
}
