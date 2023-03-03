curl -X PUT "localhost:9200/_snapshot/my_s3_repository?pretty" -H 'Content-Type: application/json' -d'
{
  "type": "s3",
  "settings": {
    "bucket": "my-bucket",
    "client": "my-alternate-client"
  }
}
'
