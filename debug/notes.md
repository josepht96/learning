# Catch short-lived CPU burners
pidstat 1 20
# if not available:
top -bn10 -d1 | grep -v "^ " | awk '$9 > 10 {print}'

# Watch kernel worker CPU consumption
ps -eo pid,stat,pcpu,comm --sort=-pcpu | head -20 | grep -E "kworker|overlay|crio"

# What processes are actually sitting in the runqueue right now
ps -eo pid,stat,pcpu,comm --sort=-pcpu | head -30

# How many elasticsearch processes and their CPU
ps -eo pid,stat,pcpu,pmem,comm --sort=-pcpu | grep elasticsearch | head -20

# Total CPU across all ES processes
ps -eo pcpu,comm | grep elasticsearch | awk '{sum += $1} END {print sum "%"}'

# Is it GC thrashing — JVM GC shows as short CPU spikes
# Check ES logs
journalctl -u elasticsearch --since "30 minutes ago" | grep -iE "gc|warn|error|timeout" | tail -30

# Or if running as a pod
oc logs <elasticsearch-pod> --since=30m | grep -iE "gc|warn|heap|timeout" | tail -30

for pid in $(ps aux | grep "elasticsearch-o" | grep -v grep | awk '{print $2}'); do
  echo -n "PID $pid: "
  cat /proc/$pid/cgroup | grep -o 'pod[a-z0-9-]*' | head -1
done

oc get pods -A -o json | jq -r '.items[] | select(.metadata.uid=="<POD_UID>") | "\(.metadata.namespace)/\(.metadata.name)"'