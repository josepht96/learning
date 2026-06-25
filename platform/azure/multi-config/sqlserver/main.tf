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

resource "azurerm_mssql_server" "sql-server" {
  name                         = "sql-server-test-11353"
  resource_group_name          = data.azurerm_resource_group.resource-group.name 
  location                     = data.azurerm_resource_group.resource-group.location
  version                      = "12.0"
  administrator_login          = "testuser"
  administrator_login_password = "TestPassword123!"
  lifecycle {
    ignore_changes = [
      # Ignore changes to tags, e.g. because a management agent
      # updates these based on some ruleset managed elsewhere.
      #tags,
    ]
  }
}
