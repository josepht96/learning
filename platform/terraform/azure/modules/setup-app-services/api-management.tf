resource "azurerm_api_management" "api-management" {
  name                 = var.api_management_name
  location             = var.location
  resource_group_name  = data.azurerm_resource_group.resource-group.name
  publisher_name       = var.publisher_name
  publisher_email      = var.publisher_email
  sku_name             = var.api_management_sku
  virtual_network_type = "Internal"
  virtual_network_configuration {
    subnet_id = azurerm_subnet.subnet.id
  }
  tags = var.tags
  depends_on = [azurerm_resource_group.rg]
}

resource "azurerm_api_management_api" "api" {
  count               = length(var.api_display_names)
  name                = var.app_names[count.index]
  resource_group_name = data.azurerm_resource_group.resource-group.name
  api_management_name = azurerm_api_management.api-management.name
  revision            = var.api_revision[count.index]
  display_name        = var.api_display_names[count.index]
  path                = var.api_path[count.index]
  protocols           = ["https"]

  import {
    content_format = "openapi+json-link"
    content_value  = var.api_content_value[count.index]
  }
}

resource "azurerm_api_management_logger" "api-management-logging" {
  name                = "${var.api_management_name}-logger"
  api_management_name = azurerm_api_management.api-management.name
  resource_group_name = data.azurerm_resource_group.resource-group.name
  resource_id         = azurerm_application_insights.app-insights.id

  application_insights {
    instrumentation_key = azurerm_application_insights.app-insights.instrumentation_key
  }
}
