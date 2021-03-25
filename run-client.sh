#!/usr/bin/env bash

# Runs Privateness in desktop client configuration

set -x

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
echo "Ness binary dir:" "$DIR"
pushd "$DIR" >/dev/null

COMMIT=$(git rev-parse HEAD)
BRANCH=$(git rev-parse --abbrev-ref HEAD)
GOLDFLAGS="${GOLDFLAGS} -X main.Commit=${COMMIT} -X main.Branch=${BRANCH}"

GORUNFLAGS=${GORUNFLAGS:-}

go run -ldflags "${GOLDFLAGS}" $GORUNFLAGS cmd/ness3/ness3.go \
    -gui-dir="${DIR}/src/gui/static/" \
    -max-default-peer-outgoing-connections=7 \
    -launch-browser=true \
    -enable-all-api-sets=true \
    -enable-gui=true \
    -log-level=debug \
    $@

popd >/dev/null
