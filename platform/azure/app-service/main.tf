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

variable "app_names" {
  description = "A list of app paths."
  type        = list(string)
  default     = ["app-name-1232321"]
}
# resource "azurerm_virtual_network" "vnet" {
#   name                = "vn-test"
#   address_space       = ["10.0.0.0/16"]
#   location            = "eastus"
#   resource_group_name = azurerm_resource_group.rg.name
# }
# resource "azurerm_subnet" "subnet" {
#   name                 = "subnet-test"
#   resource_group_name  = azurerm_resource_group.rg.name
#   virtual_network_name = azurerm_virtual_network.vnet.name
#   address_prefixes     = ["10.0.1.0/24"]
# }

data "azurerm_resource_group" "resource-group" {
  name = "test-resource-group"
}

resource "azurerm_app_service_plan" "app-svc-plan" {
  name                = "plan-name-test"
  location            = "eastus"
  resource_group_name = data.azurerm_resource_group.resource-group.name
  kind                = "Windows"
  per_site_scaling    = false

  sku {
    size = "P1V2"
    tier = "Standard"
  }
}
resource "azurerm_app_service" "app-svc" {
  count               = length(var.app_names)
  location            = "eastus"
  name                = var.app_names[count.index]
  resource_group_name = data.azurerm_resource_group.resource-group.name
  app_service_plan_id = azurerm_app_service_plan.app-svc-plan.id
  https_only          = true
  site_config {
    always_on       = true
    min_tls_version = "1.2"
    ftps_state      = "Disabled"
  }

  depends_on = [azurerm_app_service_plan.app-svc-plan]
}