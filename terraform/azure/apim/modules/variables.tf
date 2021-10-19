########################## General Variables ###################################
variable "resource_group_name" {
  description = "Name of the resource group to be imported."
  type        = string
}

variable "env" {
  description = "Name of the environment"
  type        = string
}

variable "location" {
  description = "Location"
  type        = string
}

variable "tags" {
  description = "The tags to associate with your network and subnets."
  type        = map(string)

  default = {}
}

########################## App Service Variables ###############################
variable "plan_name" {
  description = "Name of the App Service Plan"
  type        = string
}

variable "kind" {
  description = "The kind of the App Service Plan to create. Possible values are Windows (also available as App), Linux, elastic (for Premium Consumption) and FunctionApp (for a Consumption Plan)."
  type        = string
  default     = "Windows"
}

variable "per_site_scaling" {
  description = "Can Apps assigned to this App Service Plan be scaled independently? If set to false apps assigned to this plan will scale to all instances of the plan."
  type        = bool
  default     = false
}

variable "size" {
  description = "Specifies the plan's instance size"
  type        = string
  default     = "S1"
}

variable "tier" {
  description = "Specifies the plan's pricing tier"
  type        = string
  default     = "Standard"
}

variable "app_names" {
  description = "A list of app names inside the plan."
  type        = list(string)
  default     = ["app1"]
}

variable "api_management_ip" {
  description = "The IP Address of the API Management"
  type        = string
  default     = null
}

variable "hybrid_host_name" {
  description = "The hybrid connection host name"
  type        = string
}

variable "hybrid_port_number" {
  description = "The hybrid connection port number"
  type        = number
}

########################## Service Bus Variables ###############################
variable "service_bus_name" {
  description = "Name of the Service Bus Relay"
  type        = string
}

variable "hybrid_connection_name" {
  description = "The hybrid connection name"
  type        = string
}

########################## App Insights Variables ##############################
variable "app_insights_name" {
  description = "Name of the Application Insight"
  type        = string
}

########################## Api Management Variables ############################
variable "api_management_name" {
  description = "Name of the Api Management"
  type        = string
}

variable "api_management_sku" {
  description = "is a string consisting of two parts separated by an underscore(_). The first part is the name, valid values include: Consumption, Developer, Basic, Standard and Premium. The second part is the capacity (e.g. the number of deployed units of the sku), which must be a positive integer (e.g. Developer_1)."
  type        = string
  default     = "Developer_1"
}

variable "publisher_name" {
  description = "API Management Publisher Name"
  type        = string
}

variable "publisher_email" {
  description = "API Management Publisher Email"
  type        = string
}

variable "api_revision" {
  description = "A list of app revisions"
  type        = list(string)
  default     = ["1"]
}

variable "api_display_names" {
  description = "A list of app display names."
  type        = list(string)
  default     = ["app 1"]
}

variable "api_path" {
  description = "A list of app paths."
  type        = list(string)
  default     = ["path"]
}

variable "api_content_value" {
  description = "A list of app content_values."
  type        = list(string)
  default     = ["http://myapi.azurewebsites.net/?format=json"]
}
