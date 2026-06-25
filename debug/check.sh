indices=(
  operate-post-importer-queue-8.3.0_2024-09-20
  operate-post-importer-queue-8.3.0_2024-09-23
  operate-post-importer-queue-8.3.0_2024-09-22
)

for idx in "${indices[@]}"; do
  echo "=== $idx ==="
  curl -s "http://localhost:9200/$idx/_search?size=20&pretty" | python3 -c "
import sys, json
d = json.load(sys.stdin)
for hit in d['hits']['hits']:
    s = hit['_source']
    print(s.get('intent'), s.get('actionType'), s.get('creationTime'), s.get('processInstanceKey'))
"
done