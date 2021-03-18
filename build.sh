#!/usr/bin/env sh

# shellcheck disable=SC2155
export GIT_COMMIT=$(git rev-list -1 HEAD)
CANONICAL_VERSION=$(cat ./VERSION)
VERSION_STRING="$CANONICAL_VERSION ($GIT_COMMIT)"

go mod vendor
go mod tidy
go build -o build/gostations -ldflags "-X main.version=$VERSION_STRING"
