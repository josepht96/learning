variable "aws_region" {
  description = "AWS region to create resources in"
  type        = string
  default     = "us-east-1"
}

variable "environment" {
  description = "Environment name (e.g., dev, staging, prod)"
  type        = string
  default     = "dev"
}

variable "admin_users" {
  description = "List of admin user names"
  type        = list(string)
  default     = ["joe"]
}

variable "developer_users" {
  description = "List of developer user names"
  type        = list(string)
  default     = ["dev-user-1"]
}

variable "billing_users" {
  description = "List of billing user names"
  type        = list(string)
  default     = ["billing-user-1"]
}

variable "password_policy" {
  description = "IAM account password policy configuration"
  type = object({
    minimum_password_length        = number
    require_lowercase_characters   = bool
    require_uppercase_characters   = bool
    require_numbers                = bool
    require_symbols                = bool
    allow_users_to_change_password = bool
    max_password_age               = number
    password_reuse_prevention      = number
  })
  default = {
    minimum_password_length        = 14
    require_lowercase_characters   = true
    require_uppercase_characters   = true
    require_numbers                = true
    require_symbols                = true
    allow_users_to_change_password = true
    max_password_age               = 90
    password_reuse_prevention      = 12
  }
}

variable "tags" {
  description = "Additional tags to apply to all resources"
  type        = map(string)
  default     = {}
}
