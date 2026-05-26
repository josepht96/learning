# ========================================
# IAM Groups
# ========================================

# Admins Group - Full AWS Access
resource "aws_iam_group" "admins" {
  name = "${var.environment}-admins"
  path = "/"
}

resource "aws_iam_group_policy_attachment" "admins_policy" {
  group      = aws_iam_group.admins.name
  policy_arn = "arn:aws:iam::aws:policy/AdministratorAccess"
}

# Developers Group - Read-only + EC2/S3 Management
resource "aws_iam_group" "developers" {
  name = "${var.environment}-developers"
  path = "/"
}

resource "aws_iam_group_policy_attachment" "developers_readonly" {
  group      = aws_iam_group.developers.name
  policy_arn = "arn:aws:iam::aws:policy/ReadOnlyAccess"
}

resource "aws_iam_group_policy_attachment" "developers_ec2" {
  group      = aws_iam_group.developers.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2FullAccess"
}

resource "aws_iam_group_policy_attachment" "developers_s3" {
  group      = aws_iam_group.developers.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonS3FullAccess"
}

# Billing Group - Billing and Cost Explorer Access
resource "aws_iam_group" "billing" {
  name = "${var.environment}-billing"
  path = "/"
}

resource "aws_iam_group_policy_attachment" "billing_policy" {
  group      = aws_iam_group.billing.name
  policy_arn = "arn:aws:iam::aws:policy/job-function/Billing"
}

resource "aws_iam_group_policy_attachment" "billing_cost_explorer" {
  group      = aws_iam_group.billing.name
  policy_arn = aws_iam_policy.cost_explorer_policy.arn
}
