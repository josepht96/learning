resource "azurerm_frontdoor" "front-door" {
  name                                         = var.front_door_name
  resource_group_name                          = var.resource_group_name
  enforce_backend_pools_certificate_name_check = false

  routing_rule {
    name               = "routing-rule-1"
    accepted_protocols = ["Https"]
    patterns_to_match  = ["/*"]
    frontend_endpoints = [var.front_door_name]
    forwarding_configuration {
      forwarding_protocol = "MatchRequest"
      backend_pool_name   = var.backend_pool
    }
  }

  backend_pool_load_balancing {
    name = "backend-load-balancer"
  }

  backend_pool_health_probe {
    name = "backend-health-probe"
  }

  backend_pool {
    name = var.backend_pool
    # for app in azurerm_app_service.app-svc :
    #For now just these two, can add additional backends for each api
    #Have not found way to auto iterate through apis
    backend {
      host_header = var.app_service_url_list[0]
      address     = var.app_service_url_list[0]
      http_port   = 80
      https_port  = 443
    }
      backend {
      host_header = var.app_service_url_list[1]
      address     = var.app_service_url_list[1]
      http_port   = 80
      https_port  = 443
    }

    load_balancing_name = "backend-load-balancer"
    health_probe_name   = "backend-health-probe"
  }

  frontend_endpoint {
    name      = var.front_door_name
    host_name = "${var.front_door_name}.azurefd.net"
  }
}
