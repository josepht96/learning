indices=(
  operate-post-importer-queue-8.3.0_2024-09-20
  operate-post-importer-queue-8.3.0_2024-09-23
  operate-post-importer-queue-8.3.0_2024-09-22
)

for idx in "${indices[@]}"; do
  echo "=== $idx ==="
  curl -s "http://localhost:9200/$idx/_count"
  echo ""
  curl -s "http://localhost:9200/$idx/_search?size=5&pretty" | grep -E '"intent"|"actionType"'
done