#!/usr/bin/env sh

GOPATH=$GOPATH:$(pwd)
export GOPATH

GIT_COMMIT=$(git rev-list -1 HEAD)
export GIT_COMMIT
CANONICAL_VERSION=$(cat ./VERSION)
export CANONICAL_VERSION
VERSION_STRING="$CANONICAL_VERSION-$GIT_COMMIT"
export VERSION_STRING

go mod vendor
go mod tidy

go build -o build/gostations -ldflags "-X main.version=$VERSION_STRING"
