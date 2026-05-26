# S3 Module - Main Configuration

# S3 Buckets
resource "aws_s3_bucket" "main" {
  for_each = toset(var.bucket_names)

  bucket = "${var.name_prefix}-${each.value}"

  tags = merge(
    var.tags,
    {
      Name = "${var.name_prefix}-${each.value}"
    }
  )
}

# Enable versioning
resource "aws_s3_bucket_versioning" "main" {
  for_each = aws_s3_bucket.main

  bucket = each.value.id

  versioning_configuration {
    status = var.enable_versioning ? "Enabled" : "Disabled"
  }
}

# Enable server-side encryption
resource "aws_s3_bucket_server_side_encryption_configuration" "main" {
  for_each = aws_s3_bucket.main

  bucket = each.value.id

  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm     = var.kms_key_id != null ? "aws:kms" : "AES256"
      kms_master_key_id = var.kms_key_id
    }
    bucket_key_enabled = var.kms_key_id != null ? true : false
  }
}

# Block public access
resource "aws_s3_bucket_public_access_block" "main" {
  for_each = aws_s3_bucket.main

  bucket = each.value.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

# Enable logging for non-log buckets
resource "aws_s3_bucket_logging" "main" {
  for_each = { for k, v in aws_s3_bucket.main : k => v if !strcontains(k, "logs") }

  bucket = each.value.id

  target_bucket = aws_s3_bucket.main["logs"].id
  target_prefix = "${each.key}/"
}

# Lifecycle rules
resource "aws_s3_bucket_lifecycle_configuration" "main" {
  for_each = aws_s3_bucket.main

  bucket = each.value.id

  rule {
    id     = "transition-old-versions"
    status = "Enabled"

    filter {}

    noncurrent_version_transition {
      noncurrent_days = 30
      storage_class   = "STANDARD_IA"
    }

    noncurrent_version_transition {
      noncurrent_days = 90
      storage_class   = "GLACIER"
    }

    noncurrent_version_expiration {
      noncurrent_days = 365
    }
  }

  rule {
    id     = "delete-incomplete-multipart-uploads"
    status = "Enabled"

    filter {}

    abort_incomplete_multipart_upload {
      days_after_initiation = 7
    }
  }

  dynamic "rule" {
    for_each = strcontains(each.key, "logs") ? [1] : []

    content {
      id     = "expire-logs"
      status = "Enabled"

      filter {}

      expiration {
        days = var.log_retention_days
      }
    }
  }
}

# Bucket policies (example - adjust as needed)
resource "aws_s3_bucket_policy" "main" {
  for_each = aws_s3_bucket.main

  bucket = each.value.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "EnforcedTLS"
        Effect = "Deny"
        Principal = "*"
        Action = "s3:*"
        Resource = [
          each.value.arn,
          "${each.value.arn}/*"
        ]
        Condition = {
          Bool = {
            "aws:SecureTransport" = "false"
          }
        }
      }
    ]
  })
}

# S3 bucket notifications (optional - example for logs bucket)
resource "aws_s3_bucket_notification" "logs" {
  count = contains(var.bucket_names, "logs") && var.enable_notifications ? 1 : 0

  bucket = aws_s3_bucket.main["logs"].id

  # Example: Send notification to SNS when new logs are uploaded
  # Uncomment and configure as needed
  # topic {
  #   topic_arn = aws_sns_topic.bucket_notifications.arn
  #   events    = ["s3:ObjectCreated:*"]
  # }
}
