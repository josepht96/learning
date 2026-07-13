# ========================================
# IAM Roles
# ========================================

# EC2 S3 Role - Assumable by EC2 instances for S3 access
resource "aws_iam_role" "ec2_s3_role" {
  name               = "${var.environment}-ec2-s3-role"
  description        = "Role for EC2 instances to access S3 buckets"
  assume_role_policy = data.aws_iam_policy_document.ec2_assume_role.json

  tags = merge(
    var.tags,
    {
      Name        = "${var.environment}-ec2-s3-role"
      Purpose     = "EC2 S3 Access"
      Environment = var.environment
    }
  )
}

# Attach S3 Full Access managed policy to ec2-s3-role
resource "aws_iam_role_policy_attachment" "ec2_s3_access" {
  role       = aws_iam_role.ec2_s3_role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonS3FullAccess"
}

# Instance Profile for EC2 to assume the role
resource "aws_iam_instance_profile" "ec2_s3_profile" {
  name = "${var.environment}-ec2-s3-profile"
  role = aws_iam_role.ec2_s3_role.name

  tags = merge(
    var.tags,
    {
      Name        = "${var.environment}-ec2-s3-profile"
      Environment = var.environment
    }
  )
}

# CI/CD Role - Assumable by EC2 instances for ECR and ECS access
resource "aws_iam_role" "ci_cd_role" {
  name               = "${var.environment}-ci-cd-role"
  description        = "Role for CI/CD server to push to ECR and deploy to ECS"
  assume_role_policy = data.aws_iam_policy_document.ec2_assume_role.json

  tags = merge(
    var.tags,
    {
      Name        = "${var.environment}-ci-cd-role"
      Purpose     = "CI/CD Deployment"
      Environment = var.environment
    }
  )
}

# Attach inline policy for ECR and ECS permissions
resource "aws_iam_role_policy" "ci_cd_ecr_ecs_policy" {
  name   = "${var.environment}-ci-cd-ecr-ecs-policy"
  role   = aws_iam_role.ci_cd_role.id
  policy = data.aws_iam_policy_document.ci_cd_policy.json
}

# Instance Profile for CI/CD EC2 instance
resource "aws_iam_instance_profile" "ci_cd_profile" {
  name = "${var.environment}-ci-cd-profile"
  role = aws_iam_role.ci_cd_role.name

  tags = merge(
    var.tags,
    {
      Name        = "${var.environment}-ci-cd-profile"
      Environment = var.environment
    }
  )
}

# ========================================
# Trust Policies (Assume Role Policies)
# ========================================

# EC2 assume role policy - allows EC2 service to assume the role
data "aws_iam_policy_document" "ec2_assume_role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}

# ========================================
# IAM Policy Documents
# ========================================

# CI/CD Policy - ECR and ECS permissions
data "aws_iam_policy_document" "ci_cd_policy" {
  # ECR Push Permissions
  statement {
    sid    = "ECRPushAccess"
    effect = "Allow"

    actions = [
      "ecr:GetAuthorizationToken",
      "ecr:BatchCheckLayerAvailability",
      "ecr:GetDownloadUrlForLayer",
      "ecr:BatchGetImage",
      "ecr:PutImage",
      "ecr:InitiateLayerUpload",
      "ecr:UploadLayerPart",
      "ecr:CompleteLayerUpload",
      "ecr:DescribeRepositories",
      "ecr:ListImages",
      "ecr:DescribeImages"
    ]

    resources = ["*"]
  }

  # ECS Deployment Permissions
  statement {
    sid    = "ECSDeployAccess"
    effect = "Allow"

    actions = [
      "ecs:UpdateService",
      "ecs:DescribeServices",
      "ecs:DescribeTaskDefinition",
      "ecs:DescribeTasks",
      "ecs:ListTasks",
      "ecs:ListServices",
      "ecs:RegisterTaskDefinition",
      "ecs:DeregisterTaskDefinition",
      "ecs:RunTask",
      "ecs:StopTask",
      "ecs:DescribeClusters",
      "ecs:ListClusters"
    ]

    resources = ["*"]
  }

  # IAM PassRole for ECS Task Execution
  statement {
    sid    = "IAMPassRole"
    effect = "Allow"

    actions = [
      "iam:PassRole"
    ]

    resources = ["*"]

    condition {
      test     = "StringLike"
      variable = "iam:PassedToService"
      values   = ["ecs-tasks.amazonaws.com"]
    }
  }

  # CloudWatch Logs for deployment monitoring
  statement {
    sid    = "CloudWatchLogsAccess"
    effect = "Allow"

    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents",
      "logs:DescribeLogStreams"
    ]

    resources = ["*"]
  }
}
