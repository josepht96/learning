resource "azurerm_mssql_database" "test" {
  name           = format("%s-%s-db", var.prefix, var.tenant_short_name)
  server_id      = var.sql_server_id
  collation      = "SQL_Latin1_General_CP1_CI_AS"
  license_type   = "LicenseIncluded"
  max_size_gb    = 4
  read_scale     = true
  sku_name       = "BC_Gen5_2"

  tags = {
    foo = "bar"
  }

}
# resource "azurerm_mssql_database_extended_auditing_policy" "example" {
#   database_id                             = azurerm_mssql_database.test.id
#   storage_endpoint                        = azurerm_storage_account.example.primary_blob_endpoint
#   storage_account_access_key              = azurerm_storage_account.example.primary_access_key
#   storage_account_access_key_is_secondary = false
#   retention_in_days                       = 6
# }