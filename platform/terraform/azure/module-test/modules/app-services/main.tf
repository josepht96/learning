resource "azurerm_app_service_plan" "app_service_plan" {
  name                = var.app_service_plan_name
  location            = var.location
  resource_group_name = var.resource_group_name
  kind                = var.app_service_plan_kind
  sku {
    tier = var.app_service_plan_tier
    size = var.app_service_plan_size
  }
}


resource "azurerm_app_service" "app_service" {
  count               = length(var.app_service_name)
  name                = var.app_service_name[count.index]
  location            = var.location
  resource_group_name = var.resource_group_name
  app_service_plan_id = azurerm_app_service_plan.app_service_plan.id

  # site_config {
  #   always_on       = true
  #   min_tls_version = var.app_service_min_tls_version
  #   ftps_state      = var.app_service_ftps_state
  # }
}

