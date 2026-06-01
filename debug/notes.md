# Catch short-lived CPU burners
pidstat 1 20
# if not available:
top -bn10 -d1 | grep -v "^ " | awk '$9 > 10 {print}'

# Watch kernel worker CPU consumption
ps -eo pid,stat,pcpu,comm --sort=-pcpu | head -20 | grep -E "kworker|overlay|crio"

# What processes are actually sitting in the runqueue right now
ps -eo pid,stat,pcpu,comm --sort=-pcpu | head -30