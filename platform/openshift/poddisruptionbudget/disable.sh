# before upgrade: snapshot current values, then relax everything
oc get pdb -A -o json > pdb-backup.json
oc get pdb -A -o json | jq -r '.items[] | "\(.metadata.namespace) \(.metadata.name)"' | \
  while read ns name; do
    oc patch pdb "$name" -n "$ns" --type=merge -p '{"spec":{"minAvailable":0}}'
  done

# after upgrade: restore from backup
# (script to reapply original minAvailable/maxUnavailable per PDB from pdb-backup.json)