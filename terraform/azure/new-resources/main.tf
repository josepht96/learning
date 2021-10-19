terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = ">= 2.26"
    }
  }
  required_version = ">= 0.14.9"
}

provider "azurerm" {
  features {}
}

#app config
#event hub
#function
#storage
#blob storage

variable "location" {
  description = "A list of app paths."
  type        = string
  default     = "eastus"
}
variable "containers" {
  type    = list(string)
  default = ["testcont1", "testcont2"]
}

data "azurerm_resource_group" "resource-group" {
  name = "test-resource-group"
}




