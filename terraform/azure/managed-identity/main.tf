data "azurerm_resource_group" "resource-group" {
  name = var.resource_group_name
}

resource "azurerm_user_assigned_identity" "test-managed-identity" {
  resource_group_name = data.azurerm_resource_group.resource-group.name
  location            = "eastus"

  name = "test-MI"
}
