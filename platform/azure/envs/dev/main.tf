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

############################ East US  ##########################################

module "appsvc-eastus" {
  source                 = "../../../modules/setup-app-services"
  resource_group_name    = "test-resource-group2"
  plan_name              = "plan-name-test"
  env                    = "dev"
  location               = "eastus"
  kind                   = "Windows"
  per_site_scaling       = false
  size                   = "P1V2"
  tier                   = "Premium"
  virtual_network        = "virtual-network-test"
  subnet_name            = "subnet-test-name"
  network_security_group = "network-security-group-test"
  #   apim_subnet_id         = module.network.vnet_subnet_ids_eastus[6]
  #   app_subnet_id          = module.network.vnet_subnet_ids_eastus[5]
  hybrid_host_name       = "SomeHybridName"
  hybrid_port_number     = 1433
  hybrid_connection_name = "HybridConn"
  front_door_name = "front-door-name12321sd3"
  backend_pool = "backend-pool-123213"
  app_names = [
    "app-one-12321312323232312",
    "app-two-23223232321212",
    "app-three-12342323252345"
  ]

  service_bus_name  = "service-bus-test"
  app_insights_name = "app-insights-test"

  api_management_name = "apim-test-mgmt-123125433"
  api_management_sku  = "Developer_1"
  publisher_name      = "Thomas, Joe"
  publisher_email     = "Joe.Thomas@allscripts.com"

  api_revision = [
    "v1",
    "v1",
    "v1",
  ]
  api_display_names = [
    "One API",
    "Two API",
    "Three API"
  ]
  api_path = [
    "one",
    "two",
    "three"
  ]

  api_content_value = [
    "somecontent12232312321.com",
    "somecontent12d3122323321.com",
    "somecontent12d31223233321.com",

  ]

  tags = {
    "organization" = "Test Org",
    "sub-org"      = "Test Sub",
    "department"   = "Department of Test",
    "project"      = "Testing",
  }

}
