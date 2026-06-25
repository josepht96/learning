indices=(
  operate-post-importer-queue-8.3.0_2024-09-20
  operate-post-importer-queue-8.3.0_2024-09-23
  operate-post-importer-queue-8.3.0_2024-09-22
)

for idx in "${indices[@]}"; do
  echo "=== $idx ==="
  echo -n "CREATED count: "
  curl -s "http://localhost:9200/$idx/_count" -H 'Content-Type: application/json' -d '{"query":{"term":{"intent":"CREATED"}}}' | python3 -c "import sys,json; print(json.load(sys.stdin)['count'])"
  echo -n "RESOLVED count: "
  curl -s "http://localhost:9200/$idx/_count" -H 'Content-Type: application/json' -d '{"query":{"term":{"intent":"RESOLVED"}}}' | python3 -c "import sys,json; print(json.load(sys.stdin)['count'])"
  echo ""
done