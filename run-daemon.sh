#!/usr/bin/env bash

# Runs Privateness in server daemon configuration

set -x

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
echo "Ness binary dir:" "$DIR"
pushd "$DIR" >/dev/null

COMMIT=$(git rev-parse HEAD)
BRANCH=$(git rev-parse --abbrev-ref HEAD)
GOLDFLAGS="${GOLDFLAGS} -X main.Commit=${COMMIT} -X main.Branch=${BRANCH}"

GORUNFLAGS=${GORUNFLAGS:-}
export USER_BURN_FACTOR=5
go run -ldflags "${GOLDFLAGS}" $GORUNFLAGS cmd/ness3/ness3.go \
    -gui-dir="${DIR}/src/gui/static/" \
    -max-default-peer-outgoing-connections=7 \
    -launch-browser=false \
<<<<<<< HEAD
    -enable-all-api-sets=true \
    -enable-gui=false \
=======
>>>>>>> release/0.27.1
    -log-level=debug \
    $@

popd >/dev/null
