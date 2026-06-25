for idx in $(curl -s "http://localhost:9200/_cat/indices/*batch-operation*?h=index"); do
  echo "=== $idx ==="
  curl -s "http://localhost:9200/$idx/_mapping" | grep -A1 '"id"'
done