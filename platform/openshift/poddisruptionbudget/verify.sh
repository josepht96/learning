oc get pdb -A
diff <(oc get pdb -A -o json | jq -S '[.items[] | {ns:.metadata.namespace, name:.metadata.name, minAvailable:.spec.minAvailable, maxUnavailable:.spec.maxUnavailable}]') \
     <(jq -S '[.items[] | {ns:.metadata.namespace, name:.metadata.name, minAvailable:.spec.minAvailable, maxUnavailable:.spec.maxUnavailable}]' pdb-backup.json)