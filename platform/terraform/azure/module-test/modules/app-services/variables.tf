variable "resource_group_name" {
  description = "rg name"
  type        = string
}

variable "location" {
  description = "(Required) Location"
  type        = string
}

variable "app_service_plan_name" {
  description = "name of the app service plan which will contains all  the app  service"
  type        = string
}

variable "app_service_plan_kind" {
  description = "Kind of Plan (Windows/linux)"
  type        = string
}

variable "app_service_plan_tier" {
  description = "Tier of the app service plan"
  type        = string
}

variable "app_service_plan_size" {
  description = "size of the app service plan"
  type        = string
}

variable "app_service_name" {
  description = "name of the app service plan which will contains all  the app  service"
  type        = list(string)
}




