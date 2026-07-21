for idx in <index1> <index2> <index3>; do
  echo "=== $idx ==="
  curl -s "http://localhost:9200/$idx/_count"
  echo ""
  curl -s "http://localhost:9200/$idx/_search?size=5&pretty" | grep -E '"intent"|"actionType"'
done