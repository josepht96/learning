indices=(
  operate-post-importer-queue-8.3.0_2024-09-20
  operate-post-importer-queue-8.3.0_2024-09-23
  operate-post-importer-queue-8.3.0_2024-09-22
)

curl -X DELETE "http://localhost:9200/$(IFS=,; echo "${indices[*]}")"