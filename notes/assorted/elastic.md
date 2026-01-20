curl "localhost:9200/_cat/nodes?v&h=name,heap.percent,heap.current,heap.max,ram.percent"
curl -X POST "localhost:9200/_cache/clear?fielddata=true"
curl -X POST "localhost:9200/_cache/clear?query=true"
curl "localhost:9200/_nodes/stats/breaker?pretty"