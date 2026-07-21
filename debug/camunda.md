curl -s "http://<es-host>:9200/operate-post-importer-queue-8.3.0_/_mapping?pretty"

curl -s "http://localhost:9200/_cat/indices/*operate-post-importer-queue*?v"


[2026-06-25 14:51:26.651] [pool-4-thread-1] ERROR
	io.camunda.search.schema.IndexSchemaValidator - Index name: operate-post-importer-queue-8.3.0_. Unsupported index changes have been introduced. Data migration is required. Changes found: [PropertyDifference[name=creationTime, leftValue=IndexMappingProperty{name='creationTime', typeDefinition={format=date_time || epoch_millis[String], type=date[String]}}, rightValue=IndexMappingProperty{name='creationTime', typeDefinition={type=date[String]}}], PropertyDifference[name=actionType, leftValue=IndexMappingProperty{name='actionType', typeDefinition={type=keyword[String]}}, rightValue=IndexMappingProperty{name='actionType', typeDefinition={type=text[String], fields={keyword={type=keyword, ignore_above=256}}[LinkedHashMap]}}], PropertyDifference[name=partitionId, leftValue=IndexMappingProperty{name='partitionId', typeDefinition={type=integer[String]}}, rightValue=IndexMappingProperty{name='partitionId', typeDefinition={type=long[String]}}], PropertyDifference[name=id, leftValue=IndexMappingProperty{name='id', typeDefinition={type=keyword[String]}}, rightValue=IndexMappingProperty{name='id', typeDefinition={type=text[String], fields={keyword={type=keyword, ignore_above=256}}[LinkedHashMap]}}], PropertyDifference[name=intent, leftValue=IndexMappingProperty{name='intent', typeDefinition={type=keyword[String]}}, rightValue=IndexMappingProperty{name='intent', typeDefinition={type=text[String], fields={keyword={type=keyword, ignore_above=256}}[LinkedHashMap]}}]]
[2026-06-25 14:51:26.652] [pool-4-thread-1] WARN 
	io.camunda.zeebe.util.retry.RetryDecorator - Retrying operation for 'init schema': attempt 22. Message: Index name: operate-post-importer-queue-8.3.0_. Unsupported index changes have been introduced. Data migration is required. Changes found: [PropertyDifference[name=creationTime, leftValue=IndexMappingProperty{name='creationTime', typeDefinition={format=date_time || epoch_millis[String], type=date[String]}}, rightValue=IndexMappingProperty{name='creationTime', typeDefinition={type=date[String]}}], PropertyDifference[name=actionType, leftValue=IndexMappingProperty{name='actionType', typeDefinition={type=keyword[String]}}, rightValue=IndexMappingProperty{name='actionType', typeDefinition={type=text[String], fields={keyword={type=keyword, ignore_above=256}}[LinkedHashMap]}}], PropertyDifference[name=partitionId, leftValue=IndexMappingProperty{name='partitionId', typeDefinition={type=integer[String]}}, rightValue=IndexMappingProperty{name='partitionId', typeDefinition={type=long[String]}}], PropertyDifference[name=id, leftValue=IndexMappingProperty{name='id', typeDefinition={type=keyword[String]}}, rightValue=IndexMappingProperty{name='id', typeDefinition={type=text[String], fields={keyword={type=keyword, ignore_above=256}}[LinkedHashMap]}}], PropertyDifference[name=intent, leftValue=IndexMappingProperty{name='intent', typeDefinition={type=keyword[String]}}, rightValue=IndexMappingProperty{name='intent', typeDefinition={type=text[String], fields={keyword={type=keyword, ignore_above=256}}[LinkedHashMap]}}]].

[2026-06-25 14:25:47.691] [pool-3-thread-1] WARN 
	io.camunda.zeebe.util.retry.RetryDecorator - Retrying operation for 'init schema': attempt 33. Message: Ambiguous schema update. Multiple indices for mapping 'batch-operation' has different fields. Differences: '[IndexMappingDifference[equal=false, entriesOnlyOnLeft=[IndexMappingProperty{name='state', typeDefinition={type=keyword[String]}}, IndexMappingProperty{name='operationsCompletedCount', typeDefinition={type=long[String]}}, IndexMappingProperty{name='operationsFailedCount', typeDefinition={type=long[String]}}, IndexMappingProperty{name='errors', typeDefinition={type=object[String], properties={partitionId={type=integer}, type={type=keyword}, message={type=text}}[LinkedHashMap]}}], entriesOnlyOnRight=[], entriesInCommon=[IndexMappingProperty{name='operationsTotalCount', typeDefinition={type=long[String]}}, IndexMappingProperty{name='operationsFinishedCount', typeDefinition={type=long[String]}}, IndexMappingProperty{name='startDate', typeDefinition={format=date_time || epoch_millis[String], type=date[String]}}, IndexMappingProperty{name='username', typeDefinition={type=keyword[String]}}, IndexMappingProperty{name='instancesCount', typeDefinition={type=long[String]}}, IndexMappingProperty{name='type', typeDefinition={type=keyword[String]}}, IndexMappingProperty{name='name', typeDefinition={type=keyword[String]}}, IndexMappingProperty{name='endDate', typeDefinition={format=date_time || epoch_millis[String], type=date[String]}}, IndexMappingProperty{name='id', typeDefinition={type=keyword[String]}}], entriesDiffering=[]], IndexMappingDifference[equal=false, entriesOnlyOnLeft=[IndexMappingProperty{name='state', typeDefinition={type=keyword[String]}}, IndexMappingProperty{name='operationsCompletedCount', typeDefinition={type=long[String]}}, IndexMappingProperty{name='name', typeDefinition={type=keyword[String]}}, IndexMappingProperty{name='operationsFailedCount', typeDefinition={type=long[String]}}, IndexMappingProperty{name='errors', typeDefinition={type=object[String], properties={partitionId={type=integer}, type={type=keyword}, message={type=text}}[LinkedHashMap]}}], entriesOnlyOnRight=[], entriesInCommon=[IndexMappingProperty{name='operationsTotalCount', typeDefinition={type=long[String]}}, IndexMappingProperty{name='operationsFinishedCount', typeDefinition={type=long[String]}}, IndexMappingProperty{name='instancesCount', typeDefinition={type=long[String]}}], entriesDiffering=[PropertyDifference[name=endDate, leftValue=IndexMappingProperty{name='endDate', typeDefinition={format=date_time || epoch_millis[String], type=date[String]}}, rightValue=IndexMappingProperty{name='endDate', typeDefinition={type=date[String]}}], PropertyDifference[name=type, leftValue=IndexMappingProperty{name='type', typeDefinition={type=keyword[String]}}, rightValue=IndexMappingProperty{name='type', typeDefinition={type=text[String], fields={keyword={type=keyword, ignore_above=256}}[LinkedHashMap]}}], PropertyDifference[name=username, leftValue=IndexMappingProperty{name='username', typeDefinition={type=keyword[String]}}, rightValue=IndexMappingProperty{name='username', typeDefinition={type=text[String], fields={keyword={type=keyword, ignore_above=256}}[LinkedHashMap]}}], PropertyDifference[name=id, leftValue=IndexMappingProperty{name='id', typeDefinition={type=keyword[String]}}, rightValue=IndexMappingProperty{name='id', typeDefinition={type=text[String], fields={keyword={type=keyword, ignore_above=256}}[LinkedHashMap]}}], PropertyDifference[name=startDate, leftValue=IndexMappingProperty{name='startDate', typeDefinition={format=date_time || epoch_millis[String], type=date[String]}}, rightValue=IndexMappingProperty{name='startDate', typeDefinition={type=date[String]}}]]]]'.


for idx in $(curl -s "http://localhost:9200/_cat/indices/*operate-post-importer-queue*?h=index"); do
  echo "=== $idx ==="
  curl -s "http://localhost:9200/$idx/_mapping" | grep -A2 '"creationTime"'
done

for idx in $(curl -s "http://localhost:9200/_cat/indices/*operate-post-importer-queue*?h=index"); do
  echo "=== $idx ==="
  curl -s "http://localhost:9200/$idx/_mapping" | python3 -m json.tool | md5sum
done

#!/bin/bash
# identify_bad_mappings.sh
# Finds operate-post-importer-queue indices whose mapping differs from the most common one

declare -A mapping_hashes
declare -A index_to_hash

for idx in $(curl -s "http://localhost:9200/_cat/indices/*operate-post-importer-queue*?h=index"); do
  hash=$(curl -s "http://localhost:9200/$idx/_mapping" | python3 -c "
import sys, json
d = json.load(sys.stdin)
# strip out the index name key so only the mapping body is hashed
for k in d:
    print(json.dumps(d[k]['mappings'], sort_keys=True))
" | md5sum | awk '{print $1}')
  
  index_to_hash["$idx"]=$hash
  mapping_hashes["$hash"]=$((${mapping_hashes["$hash"]:-0} + 1))
done

# Find the most common hash (assume that's the "good"/majority mapping)
majority_hash=""
majority_count=0
for hash in "${!mapping_hashes[@]}"; do
  if [ "${mapping_hashes[$hash]}" -gt "$majority_count" ]; then
    majority_count=${mapping_hashes[$hash]}
    majority_hash=$hash
  fi
done

echo "Majority mapping hash: $majority_hash (found in $majority_count indices)"
echo ""
echo "=== Indices that DIFFER from majority mapping ==="
for idx in "${!index_to_hash[@]}"; do
  if [ "${index_to_hash[$idx]}" != "$majority_hash" ]; then
    echo "$idx"
  fi
done

