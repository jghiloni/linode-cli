#!/usr/bin/env bash

set -euo pipefail

HERE=$(dirname "$(realpath "${BASH_SOURCE[0]}")")

now=$(date -z UTC +"%Y%m%d-%H%M%S")
head_branch="feat/update-config-${now}"
git switch -c "${head_branch}"

"${HERE}/update-config.sh"

if git status --porcelain | grep -E -e '^\s*M\s*\.config\.json' ; then
  git add "${HERE}/../.config.json"
  git commit -S -m "Update generator config"
  git push origin "${head_branch}"
  gh pr create -a "@me" -B main -f -H "${head_branch}"
else 
  echo "There were no changes, exiting..."
  exit 0
fi