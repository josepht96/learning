# ========================================
# Account-Wide Password Policy
# ========================================

resource "aws_iam_account_password_policy" "strict" {
  minimum_password_length        = var.password_policy.minimum_password_length
  require_lowercase_characters   = var.password_policy.require_lowercase_characters
  require_uppercase_characters   = var.password_policy.require_uppercase_characters
  require_numbers                = var.password_policy.require_numbers
  require_symbols                = var.password_policy.require_symbols
  allow_users_to_change_password = var.password_policy.allow_users_to_change_password
  max_password_age               = var.password_policy.max_password_age
  password_reuse_prevention      = var.password_policy.password_reuse_prevention
}

# ========================================
# Custom IAM Policies
# ========================================

# Cost Explorer Access Policy
resource "aws_iam_policy" "cost_explorer_policy" {
  name        = "${var.environment}-cost-explorer-policy"
  description = "Grants access to AWS Cost Explorer and billing information"
  policy      = data.aws_iam_policy_document.cost_explorer_policy.json

  tags = merge(
    var.tags,
    {
      Name        = "${var.environment}-cost-explorer-policy"
      Purpose     = "Cost Explorer Access"
      Environment = var.environment
    }
  )
}

data "aws_iam_policy_document" "cost_explorer_policy" {
  statement {
    sid    = "CostExplorerReadAccess"
    effect = "Allow"

    actions = [
      "ce:GetCostAndUsage",
      "ce:GetCostAndUsageWithResources",
      "ce:GetCostCategories",
      "ce:GetCostForecast",
      "ce:GetDimensionValues",
      "ce:GetReservationCoverage",
      "ce:GetReservationPurchaseRecommendation",
      "ce:GetReservationUtilization",
      "ce:GetSavingsPlansCoverage",
      "ce:GetSavingsPlansPurchaseRecommendation",
      "ce:GetSavingsPlansUtilization",
      "ce:GetSavingsPlansUtilizationDetails",
      "ce:GetTags",
      "ce:GetUsageForecast",
      "ce:ListCostCategoryDefinitions",
      "ce:DescribeReport",
      "ce:DescribeNotificationSubscription"
    ]

    resources = ["*"]
  }

  statement {
    sid    = "BudgetsReadAccess"
    effect = "Allow"

    actions = [
      "budgets:ViewBudget",
      "budgets:DescribeBudgetAction",
      "budgets:DescribeBudgetActionsForAccount",
      "budgets:DescribeBudgetActionsForBudget",
      "budgets:DescribeBudgetNotificationsForAccount"
    ]

    resources = ["*"]
  }

  statement {
    sid    = "CURReadAccess"
    effect = "Allow"

    actions = [
      "cur:DescribeReportDefinitions",
      "cur:GetClassicReport",
      "cur:GetClassicReportPreferences",
      "cur:GetUsageReport",
      "cur:ValidateReportDestination"
    ]

    resources = ["*"]
  }
}
