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

variable "resource_group_name" {
  type    = string
  default = "test-resource-group"
}

variable "location" {
  type    = string
  default = "eastus"
}

data "azurerm_resource_group" "resource-group" {
  name = var.resource_group_name
}

variable "app_names" {
  description = "A list of app paths."
  type        = list(string)
  default     = ["app-name-1232321", "app-name-1232sd321"]
}

# data "azurerm_user_assigned_identity" "mi_data" {
#   resource_group_name = data.azurerm_resource_group.resource-group.name
#   name                = "test-MI"
# }

output "app-names" {
  value = data.azurerm_user_assigned_identity.mi_data.id
}

resource "azurerm_app_service_plan" "app-svc-plan" {
  name                = "plan-name-test"
  location            = var.location
  resource_group_name = data.azurerm_resource_group.resource-group.name
  kind                = "Windows"
  per_site_scaling    = false

  sku {
    size = "P1V2"
    tier = "Premium"
  }
}

resource "azurerm_app_service" "app-svc" {
  count               = 2
  location            = var.location
  name                = var.app_names[count.index]
  resource_group_name = data.azurerm_resource_group.resource-group.name
  app_service_plan_id = azurerm_app_service_plan.app-svc-plan.id
  https_only          = true
  site_config {
    always_on       = true
    min_tls_version = "1.2"
    ftps_state      = "Disabled"
  }

  identity {
    type         = "UserAssigned"
    identity_ids = list(data.azurerm_user_assigned_identity.mi_data.id)
  }

  depends_on = [azurerm_app_service_plan.app-svc-plan]
}

