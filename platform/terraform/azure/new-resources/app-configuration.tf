variable "config_keys" {
  type = list(object({
    key   = string
    value = string
  }))
}
variable "tags" {
  description = "Tag for each resource"
  type        = map(string)
}
locals {
  tag_list = formatlist("%s:%s", keys(var.tags), values(var.tags))
}
resource "azurerm_app_configuration" "appconf" {
  name                = "appConf1-1231232"
  resource_group_name = data.azurerm_resource_group.resource-group.name
  location            = data.azurerm_resource_group.resource-group.location
}

resource "azurerm_template_deployment" "example-arm" {
  name                = "example-deployment"
  resource_group_name = data.azurerm_resource_group.resource-group.name

  template_body = file("./appkeys.json")
  parameters = {
    configStoreName = azurerm_app_configuration.appconf.name
    location        = data.azurerm_resource_group.resource-group.location
    #these values are lists within an object, the ARM template is expecting
    #an array, so pass as single string, convert in ARM to array
    keyValueNames   = join(",", var.config_keys[*].key)
    keyValueValues  = join(",", var.config_keys[*].value)
    #tags = join(",", local.tag_list)
  } 
  deployment_mode = "Incremental"
}
