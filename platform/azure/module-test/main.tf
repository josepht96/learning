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

# module "test-rg-module" {
#   source              = "./modules/resource-group"
#   resource_group_name = var.resource_group_name
#   location            = var.location
# }
module "test-appservice-module" {
  source                = "./modules/app-services"
  resource_group_name   = "test-resource-group"
  location              = "eastus"
  app_service_plan_name = "test-123213-plan-name"
  app_service_plan_kind = "Windows"
  app_service_plan_tier = "Premium"
  app_service_plan_size = "P1V2"

  app_service_name = [
    "test-app-ConditionAPI-dev-eastus-001",
    "test-app-EncounterAPI-dev-eastus-001"
  ]
}
# module "test-front-door" {
#   source               = "./modules/front-door"
#   app_service_url_list = module.test-appservice-module.app-services-url-list
#   resource_group_name  = var.resource_group_name
#   front_door_name      = "some-front-door-name-1232"
#   backend_pool         = "backend-pool-name1232"
# }
