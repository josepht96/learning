curl -X PUT "https://your-elasticsearch-host:9200/_slm/policy/winlogsnapshots" \
  -H "Content-Type: application/json" \
  -u username:password \
  -d '{
    "retention": {
      "expire_after": "90d"
    }
  }'