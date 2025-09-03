
resource "azurerm_api_management" "api-management" {
  name                 = var.api_management_name
  location             = var.location
  resource_group_name  = data.azurerm_resource_group.resource-group.name
  publisher_name       = var.publisher_name
  publisher_email      = var.publisher_email
  sku_name             = var.api_management_sku

  tags = var.tags
}

resource "azurerm_api_management_api" "api" {
  count               = length(var.api_display_names)
  name                = var.app_names[count.index]
  resource_group_name = data.azurerm_resource_group.resource-group.name
  api_management_name = azurerm_api_management.api-management.name
  revision            = var.api_revision[count.index]
  display_name        = var.api_display_names[count.index]
  path                = var.api_path[count.index]
  protocols           = ["https"]

  # import {
  #   content_format = "openapi+json-link"
  #   content_value  = var.api_content_value[count.index]
  # }
}

resource "azurerm_api_management_logger" "api-management-logging" {
  name                = "${var.api_management_name}-logger"
  api_management_name = azurerm_api_management.api-management.name
  resource_group_name = data.azurerm_resource_group.resource-group.name
  resource_id         = azurerm_application_insights.app-insights.id

  application_insights {
    instrumentation_key = azurerm_application_insights.app-insights.instrumentation_key
  }
}

# resource "azurerm_api_management_policy" "api-management-policy" {
#   api_management_id = azurerm_api_management.api-management.id
#   xml_content       = <<XML
# <policies>
#     <inbound>
#         <base />
#         <set-header name="Request-ID" exists-action="skip">
#             <value>@{
#                 var guidBinary = new byte[16];
#                 Array.Copy(Guid.NewGuid().ToByteArray(), 0, guidBinary, 0, 10);
#                 long time = DateTime.Now.Ticks;
#                 byte[] bytes = new byte[6];
#                 unchecked
#                 {
#                        bytes[5] = (byte)(time >> 40);
#                        bytes[4] = (byte)(time >> 32);
#                        bytes[3] = (byte)(time >> 24);
#                        bytes[2] = (byte)(time >> 16);
#                        bytes[1] = (byte)(time >> 8);
#                        bytes[0] = (byte)(time);
#                 }
#                 Array.Copy(bytes, 0, guidBinary, 10, 6);
#                 return new Guid(guidBinary).ToString();
#             }</value>
#         </set-header>
#     </inbound>
#     <backend>
#         <base />
#     </backend>
#     <outbound>
#         <base />
#     </outbound>
#     <on-error>
#         <base />
#     </on-error>
# </policies>
# XML
# }
