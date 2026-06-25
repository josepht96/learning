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