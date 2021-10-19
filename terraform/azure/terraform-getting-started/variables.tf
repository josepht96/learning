variable "location" {
    type = string
    default = "eastus"

}
variable "prefix" {
  type    = string
  default = "my"
}
variable "tags" {
  type = map

  default = {
    Environment = "Terraform GS"
    Dept        = "Engineering"
    Team = "DevOps"
  }
}
variable "sku" {
    default = {
        westus2 = "16.04-LTS"
        eastus = "18.04-LTS"
    }
}
