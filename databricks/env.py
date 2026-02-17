dbutils.widgets.text("env", "preprod")
env = dbutils.widgets.get("env")

configs = {
    "preprod": {
        "bucket": "my-bucket-preprod",
        "catalog": "preprod_catalog",
        "oracle_host": "preprod-oracle.company.com",
        "s3_landing": "s3://preprod-landing/data/",
        "secret_scope": "preprod-scope"
    },
    "prod": {
        "bucket": "my-bucket-prod",
        "catalog": "prod_catalog",
        "oracle_host": "prod-oracle.company.com",
        "s3_landing": "s3://prod-landing/data/",
        "secret_scope": "prod-scope"
    }
}

config = configs[env]