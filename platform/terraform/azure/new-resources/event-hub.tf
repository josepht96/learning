# variable "event_hubs" {
#   type = list(object({
#     name              = string
#     partition_count   = number
#     message_retention = number
#   }))
# }


# resource "azurerm_eventhub_namespace" "eventhub-namespace" {
#   name                = "eventhub-namespace-asdasda231"
#   resource_group_name = data.azurerm_resource_group.resource-group.name
#   location            = data.azurerm_resource_group.resource-group.location
#   sku                 = "Standard"
#   capacity            = 4
# }

# resource "azurerm_eventhub" "example" {
#   count               = length(var.event_hubs)
#   name                = var.event_hubs[count.index].name
#   namespace_name      = azurerm_eventhub_namespace.eventhub-namespace.name
#   resource_group_name = data.azurerm_resource_group.resource-group.name
#   partition_count     = var.event_hubs[count.index].partition_count
#   message_retention   = var.event_hubs[count.index].message_retention
# }

# #I think we actually want 2 rules, listen and send
# resource "azurerm_eventhub_authorization_rule" "example" {
#   count               = length(azurerm_eventhub.example)
#   name                = "authorization"
#   namespace_name      = azurerm_eventhub_namespace.eventhub-namespace.name
#   eventhub_name       = azurerm_eventhub.example[count.index].name
#   resource_group_name = data.azurerm_resource_group.resource-group.name
#   listen              = true
#   send                = false
#   manage              = false
# }
