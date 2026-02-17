dbutils.widgets.text("env", "preprod")

%run /Repos/shared/env_config

# Now config is available
print(f"s3://{config['bucket']}")
