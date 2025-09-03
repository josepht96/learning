
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

############## Removing pending updates from Microsoft on issues

# resource "azurerm_app_service_hybrid_connection" "app-svc-hybrid-connection" {
#   count               = length(var.app_names)
#   app_service_name    = var.app_names[count.index]
#   resource_group_name = data.azurerm_resource_group.resource-group.name
#   relay_id            = azurerm_relay_hybrid_connection.service-bus-hybrid-connection.id
#   hostname            = var.hybrid_host_name
#   port                = var.hybrid_port_number

#   depends_on = [azurerm_relay_hybrid_connection.service-bus-hybrid-connection]
# }

resource "azurerm_app_service" "app-svc" {
  count               = length(var.app_names)
  location            = var.location
  name                = var.app_names[count.index]
  resource_group_name = data.azurerm_resource_group.resource-group.name
  app_service_plan_id = azurerm_app_service_plan.app-svc-plan.id
  https_only          = true
  site_config {
    always_on       = false
    min_tls_version = "1.2"
    ftps_state      = "Disabled"
    ip_restriction = [
      {
        action                    = "Allow"
        ip_address                = "${azurerm_api_management.api-management.public_ip_addresses[0]}/30"
        name                      = "Restrict to APIM"
        priority                  = 1
        headers                   = null
        service_tag               = null
        subnet_id                 = null
        virtual_network_subnet_id = null
      }
    ]
  }
  tags = var.tags

  depends_on = [azurerm_app_service_plan.app-svc-plan]
}

# resource "azurerm_monitor_autoscale_setting" "autoscalling" {
#   name                = "AutoscaleSetting"
#   resource_group_name = data.azurerm_resource_group.resource-group.name
#   location            = data.azurerm_resource_group.resource-group.location
#   target_resource_id  = azurerm_app_service_plan.app-svc-plan.id

#   profile {
#     name = "defaultProfile"

#     capacity {
#       default = 1
#       minimum = 1
#       maximum = 10
#     }

#     rule {
#       metric_trigger {
#         metric_name        = "CpuPercentage"
#         metric_resource_id = azurerm_app_service_plan.app-svc-plan.id
#         time_grain         = "PT1M"
#         statistic          = "Average"
#         time_window        = "PT5M"
#         time_aggregation   = "Average"
#         operator           = "GreaterThan"
#         threshold          = 80
#       }

#       scale_action {
#         direction = "Increase"
#         type      = "ChangeCount"
#         value     = "1"
#         cooldown  = "PT1M"
#       }
#     }
#     rule {
#       metric_trigger {
#         metric_name        = "MemoryPercentage"
#         metric_resource_id = azurerm_app_service_plan.app-svc-plan.id
#         time_grain         = "PT1M"
#         statistic          = "Average"
#         time_window        = "PT5M"
#         time_aggregation   = "Average"
#         operator           = "GreaterThan"
#         threshold          = 80
#       }

#       scale_action {
#         direction = "Increase"
#         type      = "ChangeCount"
#         value     = "1"
#         cooldown  = "PT1M"
#       }
#     }

#     rule {
#       metric_trigger {
#         metric_name        = "CpuPercentage"
#         metric_resource_id = azurerm_app_service_plan.app-svc-plan.id
#         time_grain         = "PT1M"
#         statistic          = "Average"
#         time_window        = "PT5M"
#         time_aggregation   = "Average"
#         operator           = "LessThan"
#         threshold          = 60
#       }

#       scale_action {
#         direction = "Decrease"
#         type      = "ChangeCount"
#         value     = "1"
#         cooldown  = "PT1M"
#       }
#     }
#     rule {
#       metric_trigger {
#         metric_name        = "MemoryPercentage"
#         metric_resource_id = azurerm_app_service_plan.app-svc-plan.id
#         time_grain         = "PT1M"
#         statistic          = "Average"
#         time_window        = "PT5M"
#         time_aggregation   = "Average"
#         operator           = "LessThan"
#         threshold          = 60
#       }

#       scale_action {
#         direction = "Decrease"
#         type      = "ChangeCount"
#         value     = "1"
#         cooldown  = "PT1M"
#       }
#     }
#   }

#   lifecycle {
#     ignore_changes = [
#       profile
#     ]
#   }
# }
