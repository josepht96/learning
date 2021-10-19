
resource "azurerm_relay_namespace" "service-bus-relay" {
  name                = var.service_bus_name
  location            = var.location
  resource_group_name = data.azurerm_resource_group.resource-group.name
  sku_name            = "Standard" ### At this time the only supported value is Standard
  tags                = var.tags
}

############## Removing pending updates from Microsoft on issues

# resource "azurerm_relay_hybrid_connection" "service-bus-hybrid-connection" {
#   name                          = var.hybrid_connection_name
#   resource_group_name           = data.azurerm_resource_group.resource-group.name
#   relay_namespace_name          = azurerm_relay_namespace.service-bus-relay.name
#   requires_client_authorization = true
#   user_metadata                 = "[{\"key\":\"endpoint\",\"value\":\"${var.hybrid_host_name}:${var.hybrid_port_number}\"}]"
# }
