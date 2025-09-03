resource "azurerm_app_service_plan" "app-svc-plan" {
  name                = var.plan_name
  location            = var.location
  resource_group_name = data.azurerm_resource_group.resource-group.name
  kind                = var.kind
  per_site_scaling    = var.per_site_scaling

  sku {
    size = var.size
    tier = var.tier
  }
  tags = var.tags
}

resource "azurerm_application_insights" "app-insights" {
  name                = var.app_insights_name
  location            = var.location
  resource_group_name = data.azurerm_resource_group.resource-group.name
  application_type    = "web"
  tags                = var.tags
}

resource "azurerm_app_service" "app-svc" {
  count               = length(var.app_names)
  location            = var.location
  name                = var.app_names[count.index]
  resource_group_name = data.azurerm_resource_group.resource-group.name
  app_service_plan_id = azurerm_app_service_plan.app-svc-plan.id
  https_only          = true
  site_config {
    always_on       = true
    min_tls_version = "1.2"
    ftps_state      = "Disabled"
    ip_restriction = [
      {
        action                    = "Deny"
        ip_address                = "Any"
        name                      = "Deny All"
        priority                  = 500
        headers                   = null
        service_tag               = null
        subnet_id                 = null
        virtual_network_subnet_id = azurerm_subnet.subnet.id
      }
    ]
  }
  tags = var.tags

  depends_on = [azurerm_app_service_plan.app-svc-plan]
}
# resource "azurerm_app_service_virtual_network_swift_connection" "vnet-integration-connection" {
#   count          = length(var.app_names)
#   app_service_id = azurerm_app_service.app-svc[count.index].id
#   subnet_id      = azurerm_subnet.subnet.id
#   depends_on     = [azurerm_app_service.app-svc]
# }
