terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "2.56.0"
    }
  }
}

provider "azurerm" {
  features {}
}

data "azurerm_resource_group" "resource-group" {
  name = "test-resource-group"
}
resource "azurerm_storage_account" "example" {
  name                     = "examplesaasdasd12"
  resource_group_name      = data.azurerm_resource_group.resource-group.name
  location                 = data.azurerm_resource_group.resource-group.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
}

resource "azurerm_mssql_server" "example" {
  name                         = "example-sqlserverasdasdsad2"
  resource_group_name          = data.azurerm_resource_group.resource-group.name
  location                     = data.azurerm_resource_group.resource-group.location
  version                      = "12.0"
  administrator_login          = "4dm1n157r470r"
  administrator_login_password = "4-v3ry-53cr37-p455w0rd"
}

module "tenant" {
  for_each = var.tenants
  source   = "./modules/tenant"

  prefix        = local.prefix
  tenant_short_name         = each.key

}