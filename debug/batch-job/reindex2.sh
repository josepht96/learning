old_indices=(
operate-batch-operation-1.0.0_2023-10-11
operate-batch-operation-1.0.0_2023-09-26
operate-batch-operation-1.0.0_2023-09-27
operate-batch-operation-1.0.0_2023-10-10
operate-batch-operation-1.0.0_2023-09-28
operate-batch-operation-1.0.0_2023-09-29
operate-batch-operation-1.0.0_2023-09-12
operate-batch-operation-1.0.0_2023-09-15
operate-batch-operation-1.0.0_2023-09-19
operate-batch-operation-1.0.0_2023-10-09
operate-batch-operation-1.0.0_2023-10-06
operate-batch-operation-1.0.0_2023-08-30
operate-batch-operation-1.0.0_2023-08-31
)

for idx in "${old_indices[@]}"; do
  echo "=== Reindexing $idx ==="

  curl -s -X PUT "http://localhost:9200/${idx}-fixed" -H 'Content-Type: application/json' -d '{
    "mappings": {
      "dynamic": "strict",
      "properties": {
        "endDate": {"type": "date", "format": "date_time || epoch_millis"},
        "errors": {
          "properties": {
            "message": {"type": "text"},
            "partitionId": {"type": "integer"},
            "type": {"type": "keyword"}
          }
        },
        "id": {"type": "keyword"},
        "instancesCount": {"type": "long"},
        "name": {"type": "keyword"},
        "operationsCompletedCount": {"type": "long"},
        "operationsFailedCount": {"type": "long"},
        "operationsFinishedCount": {"type": "long"},
        "operationsTotalCount": {"type": "long"},
        "startDate": {"type": "date", "format": "date_time || epoch_millis"},
        "state": {"type": "keyword"},
        "type": {"type": "keyword"},
        "username": {"type": "keyword"}
      }
    }
  }'

  curl -s -X POST "http://localhost:9200/_reindex" -H 'Content-Type: application/json' -d "{
    \"source\": {\"index\": \"$idx\"},
    \"dest\": {\"index\": \"${idx}-fixed\"}
  }"

  curl -s -X DELETE "http://localhost:9200/$idx"
  curl -s -X PUT "http://localhost:9200/$idx" -H 'Content-Type: application/json' -d '{
    "mappings": {
      "dynamic": "strict",
      "properties": {
        "endDate": {"type": "date", "format": "date_time || epoch_millis"},
        "errors": {
          "properties": {
            "message": {"type": "text"},
            "partitionId": {"type": "integer"},
            "type": {"type": "keyword"}
          }
        },
        "id": {"type": "keyword"},
        "instancesCount": {"type": "long"},
        "name": {"type": "keyword"},
        "operationsCompletedCount": {"type": "long"},
        "operationsFailedCount": {"type": "long"},
        "operationsFinishedCount": {"type": "long"},
        "operationsTotalCount": {"type": "long"},
        "startDate": {"type": "date", "format": "date_time || epoch_millis"},
        "state": {"type": "keyword"},
        "type": {"type": "keyword"},
        "username": {"type": "keyword"}
      }
    }
  }'
  curl -s -X POST "http://localhost:9200/_reindex" -H 'Content-Type: application/json' -d "{
    \"source\": {\"index\": \"${idx}-fixed\"},
    \"dest\": {\"index\": \"$idx\"}
  }"
  curl -s -X DELETE "http://localhost:9200/${idx}-fixed"

  echo "Done with $idx"
done