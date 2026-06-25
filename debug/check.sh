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