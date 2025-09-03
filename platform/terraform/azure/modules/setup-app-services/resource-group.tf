resource "azurerm_resource_group" "rg" {
  name     = var.resource_group_name
  location = var.location

  tags = var.tags
}

#returns data object from the resource group matching the name
data "azurerm_resource_group" "resource-group" {
  name       = var.resource_group_name
  depends_on = [azurerm_resource_group.rg]
}

