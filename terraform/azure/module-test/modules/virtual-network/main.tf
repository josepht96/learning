
resource "azurerm_virtual_network" "vnet" {
  name                = var.vn_name
  address_space       = ["10.0.0.0/16"]
  location            = var.location
  resource_group_name = azurerm_resource_group.rg.name
}
resource "azurerm_subnet" "subnet" {
  name                 = var.subnet_name
  resource_group_name  = azurerm_resource_group.rg.name
  virtual_network_name = azurerm_virtual_network.vnet.name
  address_prefixes     = ["10.0.1.0/24"]
}

data "azurerm_subnet" "azure-return-subnet-data" {
  name = azurerm_subnet.subnet.name
  resource_group_name = data.azurerm_resource_group.resource-group.name
}

