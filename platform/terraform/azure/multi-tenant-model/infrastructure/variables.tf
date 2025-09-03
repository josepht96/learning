variable "resource_group" {
  description = "resource group to apply infrastructure to"
  type        = string
}
variable "tenants" {
  description = "Configuration for all of the tenants"
  type = map(object({}))

}