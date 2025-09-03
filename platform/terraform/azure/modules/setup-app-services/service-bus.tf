resource "azurerm_relay_namespace" "service-bus-relay" {
  name                = var.service_bus_name
  location            = var.location
  resource_group_name = data.azurerm_resource_group.resource-group.name
  sku_name            = "Standard" ### At this time the only supported value is Standard
  tags                = var.tags
}
