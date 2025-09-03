provider "azurerm" {
  features {}
}

terraform {

}

data "azurerm_api_management" "api-management" {
  name = "apim-test-bdservices-test-eastus"
  resource_group_name = "test-resource-group"
}

output "private_ip" {
    value = data.azurerm_api_management.api-management.private_ip_addresses[*]
}
output "public" {
    value = data.azurerm_api_management.api-management.public_ip_addresses[*]
}

