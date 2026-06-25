terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "<= 2.64"
    }
  }
  required_version = ">= 1.0.2"
}
provider "azurerm" {
  features {    
  }
}
data "azurerm_resource_group" "resource-group" {
  name = "test-resource-group"
}

data "azurerm_client_config" "current" {}

resource "azurerm_key_vault" "kv" {
  name                = "fancykeyvault12324"
  location            = data.azurerm_resource_group.resource-group.location
  resource_group_name = data.azurerm_resource_group.resource-group.name
  tenant_id           = data.azurerm_client_config.current.tenant_id
  sku_name            = "standard"
}

resource "azurerm_key_vault_access_policy" "sp-manage" {
  key_vault_id       = azurerm_key_vault.kv.id
  tenant_id          = data.azurerm_client_config.current.tenant_id
  object_id          = data.azurerm_client_config.current.object_id
  secret_permissions = ["set", "delete", "get", "list"]
}
resource "azurerm_key_vault_secret" "testsecret" {
  name         = "somefancysecret12321"
  key_vault_id = azurerm_key_vault.kv.id
  value        = "sdsdsds"

  depends_on = [azurerm_key_vault_access_policy.sp-manage]
}
output "somestuff" {
    value = azurerm_key_vault_secret.testsecret.versionless_id
}