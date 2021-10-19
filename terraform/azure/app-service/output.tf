# #This is the variable you need for the app pool
output "app-names" {
    value = azurerm_app_service.app-svc[*].default_site_hostname
}
