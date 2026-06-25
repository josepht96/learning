#!/bin/bash
# apply_templates_from_file.sh
# Reads operate_templates.json and tasklist_templates.json (from _index_template export format)
# and PUTs each template individually to localhost

DEST_HOST="http://localhost:9200"

apply_templates_from_file() {
  local file="$1"
  echo "### Processing $file ###"

  python3 -c "
import json
with open('$file') as f:
    d = json.load(f)
for t in d['index_templates']:
    print(t['name'])
" | while read -r name; do
    echo "=== Applying template: $name ==="

    python3 -c "
import json
with open('$file') as f:
    d = json.load(f)
for t in d['index_templates']:
    if t['name'] == '$name':
        print(json.dumps(t['index_template']))
        break
" > "/tmp/${name}_body.json"

    curl -s -X PUT "$DEST_HOST/_index_template/$name" -H 'Content-Type: application/json' -d @"/tmp/${name}_body.json"
    echo ""
  done
}

apply_templates_from_file "operate_templates.json"
apply_templates_from_file "tasklist_templates.json"

echo "Done."