# Azure Modules: Design and Consumer Plan

## Overview

Three modules in the `terraform-modules` monorepo, versioned independently per the repository's `CLAUDE.md`:

| Module | Directory | Tag format | Purpose |
|---|---|---|---|
| Virtual network | `modules/azure-vnet` | `azure-vnet/vX.Y.Z` | VNet and subnets (with optional delegations and NSGs) |
| Virtual machine | `modules/azure-vm` | `azure-vm/vX.Y.Z` | Linux VM with NIC in a given subnet |
| App Service | `modules/azure-app-service` | `azure-app-service/vX.Y.Z` | Linux Web App with optional VNet integration |

## Composition Pattern

Modules do not reference each other. The **consumer's root module** composes them:

```
resource group (consumer)
        │
        ▼
  azure-vnet ──► subnet_ids["vm"]  ──► azure-vm          (subnet_id)
             └─► subnet_ids["app"] ──► azure-app-service (vnet_integration_subnet_id)
```

Design rules:

- **Pass IDs, not objects.** Downstream modules take a single `subnet_id` string, not a VNet object or subnet map.
- **Resource groups are created by the consumer.** Modules take `resource_group_name` and `location`; they do not create resource groups.
- **Outputs are part of the contract.** Renaming or removing an output is a MAJOR version change.

## Module Contracts

### azure-vnet

**Inputs**

| Name | Type | Required | Default | Description |
|---|---|---|---|---|
| `name` | `string` | yes | | VNet name |
| `resource_group_name` | `string` | yes | | Existing resource group |
| `location` | `string` | yes | | Azure region |
| `address_space` | `list(string)` | yes | | e.g. `["10.10.0.0/16"]` |
| `subnets` | `map(object)` | yes | | See below |
| `tags` | `map(string)` | no | `{}` | Tags applied to all resources |

`subnets` object shape:

```hcl
map(object({
  address_prefixes  = list(string)
  service_endpoints = optional(list(string), [])
  delegation = optional(object({
    name         = string
    service_name = string        # e.g. "Microsoft.Web/serverFarms"
    actions      = list(string)  # e.g. ["Microsoft.Network/virtualNetworks/subnets/action"]
  }))
  create_nsg = optional(bool, true)
}))
```

**Outputs**

| Name | Description |
|---|---|
| `vnet_id` | VNet resource ID |
| `vnet_name` | VNet name |
| `subnet_ids` | Map of subnet name → subnet ID |
| `nsg_ids` | Map of subnet name → NSG ID (for subnets with `create_nsg = true`) |

### azure-vm

Linux VM using `azurerm_linux_virtual_machine`, SSH key authentication only, no public IP by default.

**Inputs**

| Name | Type | Required | Default | Description |
|---|---|---|---|---|
| `name` | `string` | yes | | VM name |
| `resource_group_name` | `string` | yes | | Existing resource group |
| `location` | `string` | yes | | Azure region |
| `subnet_id` | `string` | yes | | Subnet for the NIC |
| `admin_username` | `string` | no | `"azureuser"` | Admin user |
| `admin_ssh_public_key` | `string` | yes | | SSH public key |
| `size` | `string` | no | `"Standard_B2s"` | VM size |
| `image` | `object` | no | Ubuntu 24.04 LTS | `{ publisher, offer, sku, version }` |
| `os_disk` | `object` | no | `{ storage_account_type = "Premium_LRS", disk_size_gb = 30 }` | OS disk settings |
| `create_public_ip` | `bool` | no | `false` | Attach a public IP |
| `tags` | `map(string)` | no | `{}` | Tags |

**Outputs**

| Name | Description |
|---|---|
| `vm_id` | VM resource ID |
| `private_ip_address` | NIC private IP |
| `nic_id` | NIC resource ID |
| `public_ip_address` | Public IP, or `null` |
| `identity_principal_id` | System-assigned managed identity principal ID |

### azure-app-service

Linux Web App (`azurerm_linux_web_app`) with its own service plan, or an existing plan.

**Inputs**

| Name | Type | Required | Default | Description |
|---|---|---|---|---|
| `name` | `string` | yes | | Web app name (globally unique) |
| `resource_group_name` | `string` | yes | | Existing resource group |
| `location` | `string` | yes | | Azure region |
| `service_plan_id` | `string` | no | `null` | Existing plan; if null, the module creates one |
| `sku_name` | `string` | no | `"P0v3"` | SKU for a module-created plan (Basic or higher for VNet integration) |
| `application_stack` | `object` | yes | | e.g. `{ docker_image_name = "nginx:latest", docker_registry_url = "https://index.docker.io" }` or a runtime (`python_version`, `node_version`, ...) |
| `app_settings` | `map(string)` | no | `{}` | App settings |
| `vnet_integration_subnet_id` | `string` | no | `null` | Delegated subnet for outbound VNet integration |
| `public_network_access_enabled` | `bool` | no | `true` | Inbound public access |
| `tags` | `map(string)` | no | `{}` | Tags |

Fixed settings: `https_only = true`, minimum TLS 1.2, system-assigned managed identity enabled.

**Outputs**

| Name | Description |
|---|---|
| `app_id` | Web app resource ID |
| `default_hostname` | Default `*.azurewebsites.net` hostname |
| `service_plan_id` | Service plan ID (created or passed in) |
| `outbound_ip_addresses` | Outbound IPs |
| `identity_principal_id` | System-assigned managed identity principal ID |

### Subnet requirements for App Service VNet integration

- Dedicated subnet; no other resources (VMs, private endpoints) in it.
- Delegated to `Microsoft.Web/serverFarms`.
- `/26` recommended to allow scaling; smaller subnets limit instance count.

## Consumer Usage

### Pattern A: Single root module (one state)

Use when one team owns the network and workloads together.

```hcl
terraform {
  required_version = ">= 1.6"
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 4.0"
    }
  }
}

provider "azurerm" {
  features {}
  subscription_id = var.subscription_id   # required in azurerm 4.x
}

locals {
  modules_repo = "git::https://<git-host>/<org>/terraform-modules.git"
  location     = "eastus2"
  tags         = { environment = "dev", owner = "team-x" }
}

resource "azurerm_resource_group" "this" {
  name     = "rg-team-x-dev"
  location = local.location
  tags     = local.tags
}

module "vnet" {
  source = "${local.modules_repo}//modules/azure-vnet?ref=azure-vnet/v0.1.0"

  name                = "vnet-team-x-dev"
  resource_group_name = azurerm_resource_group.this.name
  location            = local.location
  address_space       = ["10.10.0.0/16"]

  subnets = {
    vm = {
      address_prefixes = ["10.10.1.0/24"]
    }
    app = {
      address_prefixes = ["10.10.2.0/26"]
      delegation = {
        name         = "appservice"
        service_name = "Microsoft.Web/serverFarms"
        actions      = ["Microsoft.Network/virtualNetworks/subnets/action"]
      }
    }
  }

  tags = local.tags
}

module "vm" {
  source = "${local.modules_repo}//modules/azure-vm?ref=azure-vm/v0.2.0"

  name                 = "vm-team-x-dev"
  resource_group_name  = azurerm_resource_group.this.name
  location             = local.location
  subnet_id            = module.vnet.subnet_ids["vm"]
  admin_ssh_public_key = file("~/.ssh/id_ed25519.pub")
  tags                 = local.tags
}

module "app" {
  source = "${local.modules_repo}//modules/azure-app-service?ref=azure-app-service/v0.1.1"

  name                       = "app-team-x-dev"
  resource_group_name        = azurerm_resource_group.this.name
  location                   = local.location
  vnet_integration_subnet_id = module.vnet.subnet_ids["app"]

  application_stack = {
    docker_image_name   = "nginx:latest"
    docker_registry_url = "https://index.docker.io"
  }

  tags = local.tags
}
```

> **Note:** Terraform does not allow variables or interpolation in module `source`. The `local.modules_repo` usage above is for readability in this document only; in real code, write each `source` as a literal string.

Each module is pinned to its own version (`azure-vnet/v0.1.0`, `azure-vm/v0.2.0`, `azure-app-service/v0.1.1`) and can be upgraded independently.

### Pattern B: Separate states (platform-owned network)

Use when a platform team manages the network and app teams deploy workloads in their own state. App teams look up subnet IDs with a data source instead of a module output:

```hcl
data "azurerm_subnet" "vm" {
  name                 = "vm"
  virtual_network_name = "vnet-shared-dev"
  resource_group_name  = "rg-network-dev"
}

module "vm" {
  source = "git::https://<git-host>/<org>/terraform-modules.git//modules/azure-vm?ref=azure-vm/v0.2.0"

  name                 = "vm-team-x-dev"
  resource_group_name  = "rg-team-x-dev"
  location             = "eastus2"
  subnet_id            = data.azurerm_subnet.vm.id
  admin_ssh_public_key = file("~/.ssh/id_ed25519.pub")
}
```

Data sources are preferred over `terraform_remote_state`: app teams need only read access to the network resources, not to the platform team's state file. Subnet names therefore become a contract between the platform and app teams; document them and treat renames as breaking.

## Build Notes for the Agent

Follow the repository `CLAUDE.md` for structure, versioning, tagging, and CI. Additionally:

1. **Provider:** `azurerm >= 4.0` (minimum constraint in `versions.tf`). No `provider` blocks in modules.
2. **Build order:** `azure-vnet` first, then `azure-vm` and `azure-app-service`. Each module must be independently testable; tests and examples create their own resource group.
3. **Implement contracts exactly** as defined above. Any deviation (renamed input or output, different type) must be raised with the user before implementation.
4. **Validation blocks:** CIDR format for `address_space` and `address_prefixes`; non-empty `admin_ssh_public_key`; `sku_name` not in the Free/Shared tiers when `vnet_integration_subnet_id` is set.
5. **Security defaults:** no public IP on VMs by default; password authentication disabled; HTTPS-only and TLS 1.2 minimum on App Service; managed identity enabled on VM and App Service.
6. **Tests:** `terraform test` with `command = plan` for unit-style checks (no Azure resources created). Apply-based tests are optional and must be clearly separated, since they need Azure credentials and incur cost.
7. **Examples:** each module gets `examples/basic`. Add a repository-level `consumer-example/` implementing Pattern A, with literal `source` strings pinned to released tags.
8. **Initial versions:** all three modules start at `0.1.0`.

## Acceptance Criteria

- [ ] Three modules exist with the inputs and outputs defined above.
- [ ] `terraform validate` and `terraform test` pass for each module and example.
- [ ] `consumer-example/` composes all three modules from released tags at independent versions and passes `terraform init` and `terraform validate`.
- [ ] App Service VNet integration works against a delegated subnet created by `azure-vnet`.
- [ ] Module READMEs document inputs, outputs, and usage with the correct tag format.