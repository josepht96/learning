curl -X PUT "http://localhost:9200/operate-batch-operation-1.0.0_/_mapping" -H 'Content-Type: application/json' -d '{
  "properties": {
    "errors": {
      "properties": {
        "message": {"type": "text"},
        "partitionId": {"type": "integer"},
        "type": {"type": "keyword"}
      }
    },
    "operationsCompletedCount": {"type": "long"},
    "operationsFailedCount": {"type": "long"},
    "state": {"type": "keyword"}
  }
}'