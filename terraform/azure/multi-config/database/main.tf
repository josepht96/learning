terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "<= 2.64"
    }
  }
 #required_version = ">= 1.0.2"
}

provider "azurerm" {
  features {
    
  }
}

data "azurerm_resource_group" "resource-group" {
  name = "test-resource-group"
}

data "azurerm_mssql_server" "common" {
    name = "sql-server-test-11353"
    resource_group_name = data.azurerm_resource_group.resource-group.name
}
resource "azurerm_mssql_database" "sql-dbs" {
  #the resources index key is the unformatted key value, ie each.key, not
  #the formatted value used for 'name'
  name           = "test-database-12312e"
  server_id      = data.azurerm_mssql_server.common.id
  collation      = "SQL_Latin1_General_CP1_CI_AS"
  license_type   = "LicenseIncluded"
  max_size_gb    = 2
  sku_name       = "Basic"
  read_scale     = false
  zone_redundant = false
}
