#!/usr/bin/env bash
set -euo pipefail

BACKUP_FILE="${1:-pdb-backup.json}"

if [[ ! -f "$BACKUP_FILE" ]]; then
  echo "Backup file '$BACKUP_FILE' not found." >&2
  exit 1
fi

echo "Restoring PodDisruptionBudgets from $BACKUP_FILE ..."

jq -c '.items[]' "$BACKUP_FILE" | while read -r item; do
  ns=$(jq -r '.metadata.namespace' <<<"$item")
  name=$(jq -r '.metadata.name' <<<"$item")
  minAvail=$(jq -r '.spec.minAvailable // "null"' <<<"$item")
  maxUnavail=$(jq -r '.spec.maxUnavailable // "null"' <<<"$item")

  # Build a merge patch that restores whichever field was originally set
  # and explicitly nulls out the other, since only one can be present at a time.
  if [[ "$minAvail" != "null" ]]; then
    patch=$(jq -n --argjson v "$(jq '.spec.minAvailable' <<<"$item")" \
      '{spec: {minAvailable: $v, maxUnavailable: null}}')
  elif [[ "$maxUnavail" != "null" ]]; then
    patch=$(jq -n --argjson v "$(jq '.spec.maxUnavailable' <<<"$item")" \
      '{spec: {maxUnavailable: $v, minAvailable: null}}')
  else
    echo "  [$ns/$name] no minAvailable/maxUnavailable found in backup — skipping"
    continue
  fi

  echo "  [$ns/$name] restoring: $patch"
  if oc patch pdb "$name" -n "$ns" --type=merge -p "$patch" >/dev/null 2>&1; then
    echo "    OK"
  else
    echo "    FAILED — patch may need manual review (immutable field conflict or PDB deleted since backup)"
  fi
done

echo "Done. Verify with: oc get pdb -A"