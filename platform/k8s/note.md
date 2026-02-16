oc exec -n openshift-monitoring alertmanager-main-0 -c alertmanager -- \
  amtool silence add alertname=~".*" \
  --alertmanager.url=http://localhost:9093 \
  --author="joe" \
  --duration="2h" \
  --comment="Cluster upgrade maintenance"