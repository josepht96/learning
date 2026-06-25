curl -s "http://localhost:9200/_cat/indices/*batch-operation*?v"
curl -s "http://localhost:9200/operate-batch-operation-1.0.0_/_mapping"

curl -s "http://localhost:9200/operate-migration-steps-repository/_search?size=100&pretty" -H 'Content-Type: application/json' -d '{
  "query": {"match": {"indexName": "batch-operation"}}
}'

curl -s "http://localhost:9200/_cat/indices/*migration*?v"