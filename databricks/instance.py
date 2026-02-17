workspace_url = spark.conf.get("spark.databricks.workspaceUrl")

if "dev" in workspace_url:
    env = "dev"
elif "preprod" in workspace_url:
    env = "preprod"
else:
    env = "prod"