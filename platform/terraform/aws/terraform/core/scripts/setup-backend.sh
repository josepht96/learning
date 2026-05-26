#!/bin/bash
# Script to create S3 bucket and DynamoDB table for Terraform backend

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Check if environment is provided
if [ -z "$1" ]; then
    echo -e "${RED}Error: Environment not specified${NC}"
    echo "Usage: $0 <environment>"
    echo "Example: $0 dev"
    echo "Valid environments: dev, preprod, prod"
    exit 1
fi

ENVIRONMENT=$1
PROJECT_NAME="myproject"
REGION="us-east-1"

# Validate environment
if [[ ! "$ENVIRONMENT" =~ ^(dev|preprod|prod)$ ]]; then
    echo -e "${RED}Error: Invalid environment '$ENVIRONMENT'${NC}"
    echo "Valid environments: dev, preprod, prod"
    exit 1
fi

# Configuration
BUCKET_NAME="${PROJECT_NAME}-terraform-state-${ENVIRONMENT}"
DYNAMODB_TABLE="${PROJECT_NAME}-terraform-lock-${ENVIRONMENT}"

echo -e "${GREEN}=== Setting up Terraform backend for ${ENVIRONMENT} ===${NC}"
echo "Bucket: $BUCKET_NAME"
echo "DynamoDB Table: $DYNAMODB_TABLE"
echo "Region: $REGION"
echo ""

# Check if AWS CLI is installed
if ! command -v aws &> /dev/null; then
    echo -e "${RED}Error: AWS CLI is not installed${NC}"
    exit 1
fi

# Check AWS credentials
if ! aws sts get-caller-identity &> /dev/null; then
    echo -e "${RED}Error: AWS credentials not configured${NC}"
    exit 1
fi

echo -e "${YELLOW}Current AWS Identity:${NC}"
aws sts get-caller-identity
echo ""

read -p "Continue with backend setup? (y/n) " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "Aborted"
    exit 0
fi

# Create S3 bucket for state
echo -e "${GREEN}Creating S3 bucket: $BUCKET_NAME${NC}"

# Check if bucket exists
if aws s3 ls "s3://$BUCKET_NAME" 2>&1 | grep -q 'NoSuchBucket'; then
    # Create bucket
    if [ "$REGION" == "us-east-1" ]; then
        aws s3api create-bucket \
            --bucket "$BUCKET_NAME" \
            --region "$REGION"
    else
        aws s3api create-bucket \
            --bucket "$BUCKET_NAME" \
            --region "$REGION" \
            --create-bucket-configuration LocationConstraint="$REGION"
    fi
    echo -e "${GREEN}✓ Bucket created${NC}"
else
    echo -e "${YELLOW}Bucket already exists${NC}"
fi

# Enable versioning
echo "Enabling versioning..."
aws s3api put-bucket-versioning \
    --bucket "$BUCKET_NAME" \
    --versioning-configuration Status=Enabled
echo -e "${GREEN}✓ Versioning enabled${NC}"

# Enable encryption
echo "Enabling encryption..."
aws s3api put-bucket-encryption \
    --bucket "$BUCKET_NAME" \
    --server-side-encryption-configuration '{
        "Rules": [{
            "ApplyServerSideEncryptionByDefault": {
                "SSEAlgorithm": "AES256"
            },
            "BucketKeyEnabled": true
        }]
    }'
echo -e "${GREEN}✓ Encryption enabled${NC}"

# Block public access
echo "Blocking public access..."
aws s3api put-public-access-block \
    --bucket "$BUCKET_NAME" \
    --public-access-block-configuration \
        "BlockPublicAcls=true,IgnorePublicAcls=true,BlockPublicPolicy=true,RestrictPublicBuckets=true"
echo -e "${GREEN}✓ Public access blocked${NC}"

# Add lifecycle policy for old versions
echo "Adding lifecycle policy..."
aws s3api put-bucket-lifecycle-configuration \
    --bucket "$BUCKET_NAME" \
    --lifecycle-configuration '{
        "Rules": [{
            "Id": "delete-old-versions",
            "Status": "Enabled",
            "NoncurrentVersionExpiration": {
                "NoncurrentDays": 90
            }
        }]
    }'
echo -e "${GREEN}✓ Lifecycle policy added${NC}"

# Create DynamoDB table for state locking
echo -e "${GREEN}Creating DynamoDB table: $DYNAMODB_TABLE${NC}"

# Check if table exists
if aws dynamodb describe-table --table-name "$DYNAMODB_TABLE" --region "$REGION" &> /dev/null; then
    echo -e "${YELLOW}Table already exists${NC}"
else
    aws dynamodb create-table \
        --table-name "$DYNAMODB_TABLE" \
        --attribute-definitions AttributeName=LockID,AttributeType=S \
        --key-schema AttributeName=LockID,KeyType=HASH \
        --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 \
        --region "$REGION" \
        --tags Key=Environment,Value="$ENVIRONMENT" Key=Purpose,Value=TerraformStateLock

    echo "Waiting for table to be active..."
    aws dynamodb wait table-exists --table-name "$DYNAMODB_TABLE" --region "$REGION"
    echo -e "${GREEN}✓ Table created${NC}"
fi

# Enable point-in-time recovery for production
if [ "$ENVIRONMENT" == "prod" ]; then
    echo "Enabling point-in-time recovery for production..."
    aws dynamodb update-continuous-backups \
        --table-name "$DYNAMODB_TABLE" \
        --point-in-time-recovery-specification PointInTimeRecoveryEnabled=true \
        --region "$REGION"
    echo -e "${GREEN}✓ Point-in-time recovery enabled${NC}"
fi

echo ""
echo -e "${GREEN}=== Backend setup complete! ===${NC}"
echo ""
echo "Next steps:"
echo "1. Update environments/${ENVIRONMENT}/backend.hcl if needed"
echo "2. Initialize Terraform:"
echo "   cd terraform"
echo "   terraform init -backend-config=environments/${ENVIRONMENT}/backend.hcl"
echo ""
echo -e "${YELLOW}Backend Configuration:${NC}"
echo "  Bucket: $BUCKET_NAME"
echo "  DynamoDB Table: $DYNAMODB_TABLE"
echo "  Region: $REGION"
