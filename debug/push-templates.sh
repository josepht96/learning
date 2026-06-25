#!/bin/bash
# restore_templates.sh
# Pulls index templates from a source environment and applies them to sandbox

SOURCE_HOST="http://<source-host>:9200"
DEST_HOST="http://localhost:9200"

operate_templates=(
  "operate-template-1"
  "operate-template-2"
  # ...add all operate template names here
)

tasklist_templates=(
  "tasklist-template-1"
  "tasklist-template-2"
  # ...add all tasklist template names here
)

apply_template() {
  local name="$1"
  echo "=== Applying template: $name ==="

  curl -s "$SOURCE_HOST/_index_template/$name" -o "/tmp/${name}.json"

  python3 -c "
import json
with open('/tmp/${name}.json') as f:
    d = json.load(f)
body = d['index_templates'][0]['index_template']
with open('/tmp/${name}_body.json', 'w') as f:
    json.dump(body, f)
"

  curl -s -X PUT "$DEST_HOST/_index_template/$name" -H 'Content-Type: application/json' -d @"/tmp/${name}_body.json"
  echo ""
}

echo "### Restoring Operate templates ###"
for t in "${operate_templates[@]}"; do
  apply_template "$t"
done

echo "### Restoring Tasklist templates ###"
for t in "${tasklist_templates[@]}"; do
  apply_template "$t"
done

echo "Done."