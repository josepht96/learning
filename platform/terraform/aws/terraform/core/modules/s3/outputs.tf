# S3 Module - Outputs

output "bucket_ids" {
  description = "IDs of created S3 buckets"
  value       = { for k, v in aws_s3_bucket.main : k => v.id }
}

output "bucket_arns" {
  description = "ARNs of created S3 buckets"
  value       = { for k, v in aws_s3_bucket.main : k => v.arn }
}

output "bucket_domain_names" {
  description = "Domain names of created S3 buckets"
  value       = { for k, v in aws_s3_bucket.main : k => v.bucket_domain_name }
}

output "bucket_regional_domain_names" {
  description = "Regional domain names of created S3 buckets"
  value       = { for k, v in aws_s3_bucket.main : k => v.bucket_regional_domain_name }
}
