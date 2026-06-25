terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "<= 2.67"
    }
  }
  required_version = ">= 1.0.2"
}

provider "azurerm" {
  features {}
}

variable "location" {
  description = "A list of app paths."
  type        = string
  default     = "eastus"
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

resource "azurerm_mssql_database" "test" {
  name           = "acctest-db-d"
  server_id      = azurerm_mssql_server.example.id
  collation      = "SQL_Latin1_General_CP1_CI_AS"
  license_type   = "LicenseIncluded"
  max_size_gb    = 4
  read_scale     = true
  sku_name       = "BC_Gen5_2"
  zone_redundant = true

  # extended_auditing_policy {
  #   storage_endpoint                        = azurerm_storage_account.example.primary_blob_endpoint
  #   storage_account_access_key              = azurerm_storage_account.example.primary_access_key
  #   storage_account_access_key_is_secondary = true
  #   retention_in_days                       = 6
  # }

  tags = {
    foo = "bar"
  }

}
resource "azurerm_mssql_database_extended_auditing_policy" "example" {
  database_id                             = azurerm_mssql_database.test.id
  storage_endpoint                        = azurerm_storage_account.example.primary_blob_endpoint
  storage_account_access_key              = azurerm_storage_account.example.primary_access_key
  storage_account_access_key_is_secondary = false
  retention_in_days                       = 6
}
