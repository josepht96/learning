output "app-services-url-list" {
  value = azurerm_app_service.app_service[*].default_site_hostname
}
