# Basic IAM Module Example

This example demonstrates the basic usage of the IAM module with default settings.

## Usage

1. Configure AWS credentials:
   ```bash
   export AWS_ACCESS_KEY_ID="your-access-key"
   export AWS_SECRET_ACCESS_KEY="your-secret-key"
   export AWS_REGION="us-east-1"
   ```

2. Initialize Terraform:
   ```bash
   terraform init
   ```

3. Review the plan:
   ```bash
   terraform plan
   ```

4. Apply the configuration:
   ```bash
   terraform apply
   ```

## What Gets Created

This example creates:

- 3 IAM groups (admins, developers, billing)
- 3 IAM users (joe, dev-user-1, billing-user-1)
- 2 IAM roles (ec2-s3-role, ci-cd-role)
- 2 instance profiles (for the roles above)
- Account-wide password policy
- Custom Cost Explorer policy
- 1 example EC2 instance with S3 access

## Outputs

After applying, you'll see:

- Summary of all IAM resources
- ARNs of all groups and roles
- Instance profile names for EC2 usage

## Setting User Passwords

After resources are created, set temporary passwords for users:

```bash
# For admin user
aws iam create-login-profile --user-name joe --password TempPass123! --password-reset-required

# For developer user
aws iam create-login-profile --user-name dev-user-1 --password TempPass123! --password-reset-required

# For billing user
aws iam create-login-profile --user-name billing-user-1 --password TempPass123! --password-reset-required
```

Users will be required to change their password on first login.

## Testing

### Test Admin Access
```bash
# Admin user should have full access
aws sts get-caller-identity
aws s3 ls
aws ec2 describe-instances
```

### Test Developer Access
```bash
# Developer should have read access plus EC2/S3 management
aws s3 ls
aws ec2 describe-instances
aws s3 mb s3://test-bucket-name
```

### Test Billing Access
```bash
# Billing user should access cost information
aws ce get-cost-and-usage --time-period Start=2024-01-01,End=2024-01-31 --granularity MONTHLY --metrics BlendedCost
```

### Test EC2 Role
SSH into the EC2 instance and run:
```bash
# Should be able to access S3
aws s3 ls
aws sts get-caller-identity
```

## Clean Up

```bash
terraform destroy
```

Note: You may need to manually delete any S3 buckets created during testing before running destroy.
