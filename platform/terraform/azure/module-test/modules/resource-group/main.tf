
resource "azurerm_resource_group" "rg" {
  name     = var.resource_group_name
  location = var.location
}

data "azurerm_resource_group" "resource-group" {
  name = azurerm_resource_group.rg.name
}