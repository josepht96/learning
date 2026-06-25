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

resource "azurerm_resource_group" "rg" {
  name     = "test-rg"
  location = "eastus"
}

resource "azurerm_virtual_network" "vnet" {
  name                = "vn-test"
  address_space       = ["10.0.0.0/16"]
  location            = "eastus"
  resource_group_name = azurerm_resource_group.rg.name
}
resource "azurerm_subnet" "subnet" {
  name                 = "subnet-test"
  resource_group_name  = azurerm_resource_group.rg.name
  virtual_network_name = azurerm_virtual_network.vnet.name
  address_prefixes     = ["10.0.1.0/24"]
}

data "azurerm_resource_group" "resource-group" {
  name = azurerm_resource_group.rg.name
}

data "azurerm_subnet" "azure-return-subnet-data" {
  name = azurerm_subnet.subnet.name
  resource_group_name = azurerm_resource_group.rg.name
}

