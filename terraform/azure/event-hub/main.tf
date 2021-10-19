terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "<= 2.64"
    }
  }
  required_version = ">= 0.14.9"
}

provider "azurerm" {
  features {}
}

variable "eventHubs" {
  description = "(Required) name of the eventhub"
  type = list(object({
    name                 = string
    namespaceNameVersion = string
    partition_count      = number
    retention            = number
    authorization_name   = string
    listen               = bool
    send                 = bool
    manage               = bool
  }))
}
data "azurerm_resource_group" "resource-group" {
  name = "test-resource-group"
}

resource "azurerm_eventhub_namespace" "eventhub-namespace" {
  name                = "somerandomehnamespace"
  location            = data.azurerm_resource_group.resource-group.location
  resource_group_name = data.azurerm_resource_group.resource-group.name
  sku                 = "Standard"
  capacity            = 1
}

resource "azurerm_eventhub" "event-hub" {
  count               = length(var.eventHubs)
  name                = "${var.eventHubs[count.index].name}"
  namespace_name      = azurerm_eventhub_namespace.eventhub-namespace.name
  resource_group_name = data.azurerm_resource_group.resource-group.name
  partition_count     = var.eventHubs[count.index].partition_count
  message_retention   = var.eventHubs[count.index].retention
}

resource "azurerm_eventhub_authorization_rule" "eventhub-authorization-rule" {
  count               = length(var.eventHubs)
  name                = "navi"
  resource_group_name = data.azurerm_resource_group.resource-group.name
  namespace_name      = azurerm_eventhub_namespace.eventhub-namespace.name
  eventhub_name       = "${var.eventHubs[count.index].name}"
  listen              = var.eventHubs[count.index].listen
  send                = var.eventHubs[count.index].send
  manage              = var.eventHubs[count.index].manage
}


