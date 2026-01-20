# Elasticsearch Memory Investigation Commands

## 1. Check Node Stats and Memory Usage
```bash
curl -X GET "localhost:9200/_nodes/stats?pretty" | grep -A 20 "jvm"
```
Shows JVM heap usage, memory pools, and garbage collection stats for all nodes.

## 2. Check Cluster Health and Memory Pressure
```bash
curl -X GET "localhost:9200/_cat/nodes?v&h=name,heap.percent,heap.current,heap.max,ram.percent,ram.current,ram.max"
```
Displays memory usage in a concise table format for quick assessment.

## 3. Clear Field Data Cache
```bash
curl -X POST "localhost:9200/_cache/clear?fielddata=true&pretty"
```
Clears field data cache which can consume significant heap memory, especially with high cardinality fields.

## 4. Check and Reduce Index Segments
```bash
# Check segment count
curl -X GET "localhost:9200/_cat/segments?v"

# Force merge to reduce segments (use during low-traffic periods)
curl -X POST "localhost:9200/<index-name>/_forcemerge?max_num_segments=1"
```
Too many segments consume memory; force merging reduces segment count and memory overhead.
