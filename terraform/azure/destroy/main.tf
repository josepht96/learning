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

resource "azurerm_storage_account" "storage-account" {
  name                     = "teststoreaccount1996"
  resource_group_name      = "test-resource-group"
  location                 = "eastus"
  account_tier             = "Standard"
  account_replication_type = "GRS"

  tags = {
    environment = "test"
  }
}