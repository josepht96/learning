curl -s "http://<source-host>:9200/_component_template/operate_template" > operate_component_template.json
curl -s "http://<source-host>:9200/_component_template/tasklist_template" > tasklist_component_template.json

python3 -c "
import json
with open('operate_component_template.json') as f:
    d = json.load(f)
print(json.dumps(d['component_templates'][0]['component_template']))
" > operate_component_body.json

curl -X PUT "http://localhost:9200/_component_template/operate_template" -H 'Content-Type: application/json' -d @operate_component_body.json