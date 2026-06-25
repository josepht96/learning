curl -s "http://localhost:9200/_cat/indices/*batch-operation*?v"
curl -s "http://localhost:9200/operate-batch-operation-1.0.0_/_mapping"

curl -s "http://localhost:9200/operate-migration-steps-repository/_search?size=100&pretty" -H 'Content-Type: application/json' -d '{
  "query": {"match": {"indexName": "batch-operation"}}
}'

curl -s "http://localhost:9200/_cat/indices/*migration*?v"

curl -s "http://localhost:9200/operate-migration-steps-repository-1.1.0_/_search?size=100&pretty" -H 'Content-Type: application/json' -d '{
  "query": {"match": {"indexName": "batch-operation"}}
}'
[2026-06-25 16:44:00.440] [pool-4-thread-1] WARN 
	io.camunda.zeebe.util.retry.RetryDecorator - Retrying operation for 'init schema': attempt 10. Message: Index name: operate-batch-operation-1.0.0_. Unsupported index changes have been introduced. Data migration is required. Changes found: [PropertyDifference[name=endDate, leftValue=IndexMappingProperty{name='endDate', typeDefinition={format=date_time || epoch_millis[String], type=date[String]}}, rightValue=IndexMappingProperty{name='endDate', typeDefinition={type=date[String]}}], PropertyDifference[name=type, leftValue=IndexMappingProperty{name='type', typeDefinition={type=keyword[String]}}, rightValue=IndexMappingProperty{name='type', typeDefinition={type=text[String], fields={keyword={type=keyword, ignore_above=256}}[LinkedHashMap]}}], PropertyDifference[name=username, leftValue=IndexMappingProperty{name='username', typeDefinition={type=keyword[String]}}, rightValue=IndexMappingProperty{name='username', typeDefinition={type=text[String], fields={keyword={type=keyword, ignore_above=256}}[LinkedHashMap]}}], PropertyDifference[name=id, leftValue=IndexMappingProperty{name='id', typeDefinition={type=keyword[String]}}, rightValue=IndexMappingProperty{name='id', typeDefinition={type=text[String], fields={keyword={type=keyword, ignore_above=256}}[LinkedHashMap]}}], PropertyDifference[name=startDate, leftValue=IndexMappingProperty{name='startDate', typeDefinition={format=date_time || epoch_millis[String], type=date[String]}}, rightValue=IndexMappingProperty{name='startDate', typeDefinition={type=date[String]}}]].