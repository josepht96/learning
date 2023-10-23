resource "aws_s3_bucket" "primary-move" {
  bucket = "primary-${var.project}"

  tags = {
    Name        = "Primary bucket"
    Environment = "Dev"
  }
}