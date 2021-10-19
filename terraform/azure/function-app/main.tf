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

resource "azurerm_app_service_plan" "app-service-plan" {
  name                = "someappservice123213"
  location            = data.azurerm_resource_group.resource-group.location
  resource_group_name = data.azurerm_resource_group.resource-group.name
  kind                = "FunctionApp"
  sku {
    tier = "Standard"
    size = "S1"
  }
}
resource "azurerm_app_configuration" "app-config" {
  name                = "appconfig12323"
  location            = data.azurerm_resource_group.resource-group.location
  resource_group_name = data.azurerm_resource_group.resource-group.name
  sku                 = "standard"
}

output "connstring" {
    value = azurerm_app_configuration.app-config.primary_read_key
}
resource "azurerm_storage_account" "example" {
  name                     = "storageaccountname1232"
  location            = data.azurerm_resource_group.resource-group.location
  resource_group_name = data.azurerm_resource_group.resource-group.name
  account_tier             = "Standard"
  account_replication_type = "GRS"

  tags = {
    environment = "staging"
  }
}
resource "azurerm_function_app" "azure-function" {
  name = "somefunctionappreee132"
  location            = data.azurerm_resource_group.resource-group.location
  resource_group_name = data.azurerm_resource_group.resource-group.name
  app_service_plan_id        = azurerm_app_service_plan.app-service-plan.id
  storage_account_name       = azurerm_storage_account.example.name
  storage_account_access_key = azurerm_storage_account.example.primary_access_key
  app_settings = {
    "AppConfigConnectionString" = azurerm_app_configuration.app-config.primary_read_key[0].connection_string
  }
}
