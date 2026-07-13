# ========================================
# IAM Users
# ========================================

# Admin Users
resource "aws_iam_user" "admin_users" {
  for_each = toset(var.admin_users)

  name = each.value
  path = "/"

  tags = merge(
    var.tags,
    {
      Name        = each.value
      Group       = "admins"
      Environment = var.environment
    }
  )
}

resource "aws_iam_user_group_membership" "admin_membership" {
  for_each = toset(var.admin_users)

  user = aws_iam_user.admin_users[each.key].name
  groups = [
    aws_iam_group.admins.name
  ]
}

# Developer Users
resource "aws_iam_user" "developer_users" {
  for_each = toset(var.developer_users)

  name = each.value
  path = "/"

  tags = merge(
    var.tags,
    {
      Name        = each.value
      Group       = "developers"
      Environment = var.environment
    }
  )
}

resource "aws_iam_user_group_membership" "developer_membership" {
  for_each = toset(var.developer_users)

  user = aws_iam_user.developer_users[each.key].name
  groups = [
    aws_iam_group.developers.name
  ]
}

# Billing Users
resource "aws_iam_user" "billing_users" {
  for_each = toset(var.billing_users)

  name = each.value
  path = "/"

  tags = merge(
    var.tags,
    {
      Name        = each.value
      Group       = "billing"
      Environment = var.environment
    }
  )
}

resource "aws_iam_user_group_membership" "billing_membership" {
  for_each = toset(var.billing_users)

  user = aws_iam_user.billing_users[each.key].name
  groups = [
    aws_iam_group.billing.name
  ]
}

# ========================================
# User Login Profiles (Console Access)
# ========================================

# Note: Login profiles require password setup.
# Passwords can be set by users on first login or by administrators
# These resources create console access capability without hardcoding passwords
resource "aws_iam_user_login_profile" "admin_login" {
  for_each = toset(var.admin_users)

  user                    = aws_iam_user.admin_users[each.key].name
  password_reset_required = true

  lifecycle {
    ignore_changes = [
      password_length,
      password_reset_required,
      pgp_key,
    ]
  }
}

resource "aws_iam_user_login_profile" "developer_login" {
  for_each = toset(var.developer_users)

  user                    = aws_iam_user.developer_users[each.key].name
  password_reset_required = true

  lifecycle {
    ignore_changes = [
      password_length,
      password_reset_required,
      pgp_key,
    ]
  }
}

resource "aws_iam_user_login_profile" "billing_login" {
  for_each = toset(var.billing_users)

  user                    = aws_iam_user.billing_users[each.key].name
  password_reset_required = true

  lifecycle {
    ignore_changes = [
      password_length,
      password_reset_required,
      pgp_key,
    ]
  }
}
