variable "app_service_url_list" {
  description = "name of the app service plan which will contains all  the app  service"
  type        = list(string)
}

variable "resource_group_name" {
  description = "name of the app service plan which will contains all  the app  service"
  type        = string
}

variable "front_door_name" {
  description = "name of the app service plan which will contains all  the app  service"
  type        = string
}
variable "backend_pool" {
    description = "name of the app service plan which will contains all  the app  service"
  type        = string
}
