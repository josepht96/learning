# S3 Module - Variables

variable "environment" {
  description = "Environment name"
  type        = string
}

variable "bucket_names" {
  description = "List of bucket names to create (will be prefixed with name_prefix)"
  type        = list(string)
}

variable "name_prefix" {
  description = "Prefix for bucket names"
  type        = string
}

variable "tags" {
  description = "Common tags for all resources"
  type        = map(string)
  default     = {}
}

variable "enable_versioning" {
  description = "Enable versioning on buckets"
  type        = bool
  default     = true
}

variable "kms_key_id" {
  description = "KMS key ID for encryption (leave empty for AES256)"
  type        = string
  default     = null
}

variable "log_retention_days" {
  description = "Number of days to retain logs"
  type        = number
  default     = 90
}

variable "enable_notifications" {
  description = "Enable S3 event notifications"
  type        = bool
  default     = false
}
