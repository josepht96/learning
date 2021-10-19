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

resource "azurerm_virtual_network" "vnet" {
  name                = "test-virtual-network-2"
  address_space       = ["10.0.0.0/16"]
  location            = "eastus"
  resource_group_name = data.azurerm_resource_group.resource-group.name
}
resource "azurerm_subnet" "app-service-subnet" {
  name                 = "default"
  resource_group_name  = data.azurerm_resource_group.resource-group.name
  virtual_network_name = azurerm_virtual_network.vnet.name
  address_prefixes     = ["10.0.1.0/24"]
  delegation {
    name = "delegation"

    service_delegation {
      name    = "Microsoft.Web/serverFarms"
      actions = ["Microsoft.Network/virtualNetworks/subnets/join/action"]
    }
  }
  //service_endpoints = lookup(var.subnet_service_endpoint, var.subnet_name, null)
}