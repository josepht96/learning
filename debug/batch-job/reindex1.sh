#!/bin/bash
# Find all batch-operation indices with the old text-type id field
old_indices=()

for idx in $(curl -s "http://localhost:9200/_cat/indices/*batch-operation*?h=index"); do
  id_type=$(curl -s "http://localhost:9200/$idx/_mapping" | python3 -c "
import sys, json
d = json.load(sys.stdin)
for index_name, body in d.items():
    print(body['mappings']['properties'].get('id', {}).get('type', ''))
")
  if [ "$id_type" == "text" ]; then
    old_indices+=("$idx")
  fi
done

echo "Found ${#old_indices[@]} indices to fix"
printf '%s\n' "${old_indices[@]}"