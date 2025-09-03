resource "azurerm_storage_account" "storage-account" {
  name                     = "teststoreaccount1996"
  resource_group_name      = data.azurerm_resource_group.resource-group.name
  location                 = data.azurerm_resource_group.resource-group.location
  account_tier             = "Standard"
  account_replication_type = "GRS"

  tags = {
    environment = "test"
  }
}
resource "azurerm_storage_container" "storage_container_azure_webjobs" {
  count                 = length(var.containers)
  name                  = var.containers[count.index]
  storage_account_name  = azurerm_storage_account.storage-account.name
  container_access_type = "private"
}

resource "azurerm_storage_queue" "storage_queue" {
  name                  = "test-queue"
  storage_account_name  = azurerm_storage_account.storage-account.name
}