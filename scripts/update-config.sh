#!/usr/bin/env bash

set -euo pipefail

HERE=$(dirname "${BASH_SOURCE[0]}")
CONFIG_FILE="${HERE}/../.config.json"
CONFIG=$(cat "${CONFIG_FILE}")
for field in "api" "gen_tool"; do
  repo=$(jq -r ".${field}_repository" <<< "${CONFIG}")
  latest_version=$(gh --json tagName --jq '.[0].tagName' release list -R "${repo}")
  CONFIG=$(jq --arg field "${field}" --arg version "${latest_version}" '. + {"\($field)_version": $version}' <<< "${CONFIG}")
done

cp "${CONFIG_FILE}" "${CONFIG_FILE}".bak ||:
echo "${CONFIG}" > "${CONFIG_FILE}"