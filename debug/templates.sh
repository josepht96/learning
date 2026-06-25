curl -s "http://localhost:9200/_index_template/*operate*" > operate_templates.json
curl -s "http://localhost:9200/_index_template/*tasklist*" > tasklist_templates.json

python3 -c "
import json
with open('operate_templates.json') as f:
    data = json.load(f)
for t in data['index_templates']:
    print(t['name'])
"

curl -s "http://<source-host>:9200/_index_template/<name>" | python3 -c "
import sys, json
d = json.load(sys.stdin)['index_templates'][0]
print(json.dumps(d['index_template']))
" > template_body.json

curl -X PUT "http://localhost:9200/_index_template/<name>" -H 'Content-Type: application/json' -d @template_body.json

curl -s "http://localhost:9200/_index_template?filter_path=index_templates.name"