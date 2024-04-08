eyJ2ZXIiOiI4LjEzLjIiLCJhZHIiOlsiMTcyLjE5LjAuMjo5MjAwIl0sImZnciI6ImY3YWY3MzJhMjA2OWQwZDg5Mjk4YTQ3MTFiNWI1OGJmOWEwOTE4ZmU2ZmUzNDNkODczYmE1ZWVjOTEyZGI2YzEiLCJrZXkiOiJJejJtdm80Qk0tWUVfT3JrUnhnQTp0UjZybENRZlJOVzJFWVlFTU9WOEF3In0=

elastic
nhkozSJX_XROveGd-FNo


U2oyNXZvNEJNLVlFX09ya2xSaVQ6T2hKRlYwSG9SanljQjVfN1NIQk1SUQ==

export ES_URL="https://localhost:9200"
export API_KEY="U2oyNXZvNEJNLVlFX09ya2xSaVQ6T2hKRlYwSG9SanljQjVfN1NIQk1SUQ=="


curl -k -X POST "${ES_URL}/_bulk?pretty&pipeline=ent-search-generic-ingestion" \
  -H "Authorization: ApiKey "${API_KEY}"" \
  -H "Content-Type: application/json" \
  -d'
{ "index" : { "_index" : "test" } }
{"name": "Snow Crash", "author": "Neal Stephenson", "release_date": "1992-06-01", "page_count": 470, "_extract_binary_content": true, "_reduce_whitespace": true, "_run_ml_inference": true}
{ "index" : { "_index" : "test" } }
{"name": "Revelation Space", "author": "Alastair Reynolds", "release_date": "2000-03-15", "page_count": 585, "_extract_binary_content": true, "_reduce_whitespace": true, "_run_ml_inference": true}
{ "index" : { "_index" : "test" } }
{"name": "1984", "author": "George Orwell", "release_date": "1985-06-01", "page_count": 328, "_extract_binary_content": true, "_reduce_whitespace": true, "_run_ml_inference": true}
{ "index" : { "_index" : "test" } }
{"name": "Fahrenheit 451", "author": "Ray Bradbury", "release_date": "1953-10-15", "page_count": 227, "_extract_binary_content": true, "_reduce_whitespace": true, "_run_ml_inference": true}
{ "index" : { "_index" : "test" } }
{"name": "Brave New World", "author": "Aldous Huxley", "release_date": "1932-06-01", "page_count": 268, "_extract_binary_content": true, "_reduce_whitespace": true, "_run_ml_inference": true}
{ "index" : { "_index" : "test" } }
{"name": "The Handmaid'"'"'s Tale", "author": "Margaret Atwood", "release_date": "1985-06-01", "page_count": 311, "_extract_binary_content": true, "_reduce_whitespace": true, "_run_ml_inference": true}
'