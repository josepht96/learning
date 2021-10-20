terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = ">= 2.26"
    }
  }

  #run terraform init \
  # -backend-config "storage_account_name=somestorageacccount" \
  # -backend-config "container_name=statefiles" \
  # -backend-config "key=some.tfstate" \
  # -backend-config "sas_token=somesastoken123213"

  # Create service principal -
  # az ad sp create-for-rbac --role="Contributor" --scopes="/subscriptions/somesubid"
  
  #Setting environment variables to use SP without login
  #appId is the client_id defined above.
  #password is the client_secret defined above.
  #tenant is the tenant_id defined above.

  #Bash
  #$ export ARM_CLIENT_ID="00000000-0000-0000-0000-000000000000"
  #$ export ARM_CLIENT_SECRET="00000000-0000-0000-0000-000000000000"
  #$ export ARM_SUBSCRIPTION_ID="00000000-0000-0000-0000-000000000000"
  #$ export ARM_TENANT_ID="00000000-0000-0000-0000-000000000000"

  #Powershell
  #$Env:ARM_CLIENT_ID = "00000000-0000-0000-0000-000000000000"
  #$Env:ARM_CLIENT_SECRET = "00000000-0000-0000-0000-000000000000"
  #$Env:ARM_SUBSCRIPTION_ID = "00000000-0000-0000-0000-000000000000"
  #$Env:ARM_TENANT_ID = "00000000-0000-0000-0000-000000000000"


  backend "azurerm" {}
  required_version = ">= 0.14.9"
}

provider "azurerm" {
  features {}
}

data "azurerm_resource_group" "resource-group" {
  name = "default"
}
variable "vnets" {
  type = map(object({
    name = string
  }))
}

resource "azurerm_virtual_network" "vnet" {
  for_each            = var.vnets
  name                = each.value.name
  address_space       = ["10.0.0.0/16"]
  location            = "eastus"
  resource_group_name = data.azurerm_resource_group.resource-group.name
}
