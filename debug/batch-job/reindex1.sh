for idx in $(curl -s "http://localhost:9200/_cat/indices/*batch-operation*?h=index"); do
  echo "=== $idx ==="
  curl -s "http://localhost:9200/$idx/_mapping" | python3 -c "
import sys, json
d = json.load(sys.stdin)
for index_name, body in d.items():
    props = body['mappings']['properties']
    id_type = props.get('id', {}).get('type')
    print(f\"id field type: {id_type}\")
"
done