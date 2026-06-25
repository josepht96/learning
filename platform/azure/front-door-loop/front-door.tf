variable "front_door_name" {
  type    = string
  default = "front-door-test-dajsdsad"
}

locals {
  backend_pool_default = "backend-pool"
}

resource "azurerm_frontdoor" "front-door" {
  name                                         = var.front_door_name
  resource_group_name                          = data.azurerm_resource_group.resource-group.name
  enforce_backend_pools_certificate_name_check = false

  frontend_endpoint {
    name      = var.front_door_name
    host_name = "${var.front_door_name}.azurefd.net"
  }

  backend_pool_load_balancing {
    name = "backend-load-balancer"
  }

  backend_pool_health_probe {
    name = "backend-health-probe"
  }

  dynamic "backend_pool" {
    for_each = [
      for app in var.apps : {
        i = app
      }
    ]
    content {
      name = "${backend_pool.value.i.short}-${local.backend_pool_default}"
      backend {
        host_header = backend_pool.value.i.url
        address     = backend_pool.value.i.url
        http_port   = 80
        https_port  = 443
        weight      = 100
      }
      load_balancing_name = "backend-load-balancer"
      health_probe_name   = "backend-health-probe"
    }
  }

  dynamic "routing_rule" {
    for_each = [
      for app in var.apps : {
        #mapping the app obj to i
        i = app
      }
    ]
    content {
      name               = "${routing_rule.value.i.short}-routing-rule"
      accepted_protocols = ["Https"]
      patterns_to_match  = ["/${routing_rule.value.i.short}"]
      frontend_endpoints = [var.front_door_name]
      forwarding_configuration {
        forwarding_protocol = "MatchRequest"
        backend_pool_name   = "${routing_rule.value.i.short}-${local.backend_pool_default}"
      }
    }
  }
}
# routing_rule {
#   name               = "routing-rule-1"
#   accepted_protocols = ["Https"]
#   patterns_to_match  = ["/test"]
#   frontend_endpoints = [var.front_door_name]
#   forwarding_configuration {
#     forwarding_protocol = "MatchRequest"
#     backend_pool_name   = local.backend_pool_name
#   }
# }
# routing_rule {
#   name               = "routing-rule-2"
#   accepted_protocols = ["Https"]
#   patterns_to_match  = ["/test2"]
#   frontend_endpoints = [var.front_door_name]
#   forwarding_configuration {
#     forwarding_protocol = "MatchRequest"
#     backend_pool_name   = local.backend_pool_name2
#   }

# backend_pool {
#   name = local.backend_pool_name2
#   backend {
#     host_header = azurerm_app_service.app-svc[1].default_site_hostname
#     address     = azurerm_app_service.app-svc[1].default_site_hostname
#     http_port   = 80
#     https_port  = 443
#     weight      = 100
#   }

#   load_balancing_name = "backend-load-balancer"
#   health_probe_name   = "backend-health-probe"
# }



