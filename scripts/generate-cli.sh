#!/usr/bin/env bash

set -euo pipefail

HERE=$(dirname "$(realpath "${BASH_SOURCE[0]}")")
CONFIG_JSON=$(cat "${HERE}/../.config.json")
ONLYCLI="onlycli"

GEN_TOOL_REPO=$(jq -r '.gen_tool_repository' <<< "${CONFIG_JSON}")
GEN_TOOL_VERSION=$(jq -r '.gen_tool_version' <<< "${CONFIG_JSON}")
API_REPO=$(jq -r '.api_repository' <<< "${CONFIG_JSON}")
API_VERSION=$(jq -r '.api_version' <<< "${CONFIG_JSON}")

INSTALL_DIR="${INSTALL_DIR:-}"
if [[ -z "${INSTALL_DIR}" ]]; then
  INSTALL_DIR="${HOME}/bin"
elif [[ -w "/usr/local/bin" ]]; then
  INSTALL_DIR="/usr/local/bin"
else
  INSTALL_DIR="${PWD}"
  ONLYCLI="./onlycli"
fi

case "$(uname -s)" in
Linux) OS=linux ;;
Darwin) OS=darwin ;;
*)
	echo "install.sh: unsupported OS (need Linux or Darwin)" >&2
	exit 1
	;;
esac

ARCH_RAW=$(uname -m)
case "$ARCH_RAW" in
x86_64 | amd64) ARCH=amd64 ;;
arm64 | aarch64) ARCH=arm64 ;;
*)
	echo "install.sh: unsupported architecture: $ARCH_RAW (need amd64 or arm64)" >&2
	exit 1
	;;
esac

TMP_DIR=$(mktemp -d)
TMPSPEC=$(mktemp -p "${TMP_DIR}")
echo $TMPSPEC

pushd "${HERE}/../helpers/normalize-schema"
curl -fsSL "https://raw.githubusercontent.com/${API_REPO}/refs/tags/${API_VERSION}/openapi.json" | go run main.go > "${TMPSPEC}"
popd

# shellcheck disable=SC2064
trap "rm -fr ${TMP_DIR}" EXIT

gh release download "${GEN_TOOL_VERSION}" -O - -p "*_${OS}_${ARCH}.tar.gz" -R "${GEN_TOOL_REPO}"| tar xzf - -C "${TMP_DIR}" onlycli
install -m 0755 "${TMP_DIR}/onlycli" "${INSTALL_DIR}"
${ONLYCLI} version
set -x
${ONLYCLI} generate --module "github.com/jghiloni/linodectl/cli" --name "linodectl" --out "${HERE}/../cli" --spec "${TMPSPEC}"
