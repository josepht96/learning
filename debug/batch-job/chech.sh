indices=(
  # paste actual names from the command above
)

for idx in "${indices[@]}"; do
  echo "$idx: $(curl -s "http://localhost:9200/$idx/_count" | python3 -c "import sys,json; print(json.load(sys.stdin)['count'])")"
done