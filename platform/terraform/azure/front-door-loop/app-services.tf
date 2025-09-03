
resource "azurerm_app_service_plan" "app-svc-plan" {
  name                = "plan-name-test"
  location            = var.location
  resource_group_name = data.azurerm_resource_group.resource-group.name
  kind                = "Windows"
  per_site_scaling    = false

  sku {
    size = "P1V2"
    tier = "Premium"
  }
}

resource "azurerm_app_service" "app-svc" {
  count               = 2
  location            = var.location
  name                = var.apps[count.index].name
  resource_group_name = data.azurerm_resource_group.resource-group.name
  app_service_plan_id = azurerm_app_service_plan.app-svc-plan.id
  https_only          = true
  site_config {
    always_on       = true
    min_tls_version = "1.2"
    ftps_state      = "Disabled"
  }

  depends_on = [azurerm_app_service_plan.app-svc-plan]
}
