#!/usr/bin/env bash
set -o errexit -o nounset -o pipefail

CGO_ENABLED=0 go test -tags containers_image_openpgp "$@"
