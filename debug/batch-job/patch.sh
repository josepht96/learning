indices=(
  "operate-batch-operation-1.0.0_"
  # add other affected dated indices here
)

for idx in "${indices[@]}"; do
  echo "=== Updating $idx ==="
  curl -X PUT "http://localhost:9200/$idx/_mapping" -H 'Content-Type: application/json' -d '{
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
  echo ""
done