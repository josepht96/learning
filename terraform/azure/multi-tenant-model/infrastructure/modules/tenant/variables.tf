variable "sql_server_id" {
  description = "(required) SQL Server name to provision database"
  type        = string
}
variable "tenant_short_name" {
  description = "(required) the abbreviated name of the tenant"
  type        = string
}
variable "prefix" {
  description = "(required) prefix for naming resources"
  type        = string
}
