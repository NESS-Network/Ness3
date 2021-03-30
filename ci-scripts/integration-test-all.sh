#!/bin/bash
# Runs "stable"-mode tests against a mdl node configured with a pinned database
# "stable" mode tests assume the blockchain data is static, in order to check API responses more precisely
# $TEST defines which test to run i.e, cli or gui; If empty both are run

#Set Script Name variable
SCRIPT=`basename ${BASH_SOURCE[0]}`
PORT="8320"
RPC_PORT="8330"
IP_ADDR="0.0.0.0"
HOST="http://$IP_ADDR:$PORT"
RPC_ADDR="$IP_ADDR:$RPC_PORT"
MODE="stable"
MODE_LIVE="live"

BINARY="$PWD/mdl-integration"
TEST=""
UPDATE=""
DISABLE_CSRF="-disable-csrf"

usage () {
  echo "Usage: $SCRIPT"
  echo "Optional command line arguments"
  echo "-t <string>  -- Test to run, gui or cli; empty runs both tests"
  echo "-u <boolean> -- Update stable testdata"
  exit 1
}

while getopts "h?t:u" args; do
case $args in
    h|\?)
        usage;
        exit;;
    t ) TEST=${OPTARG};;
    u ) UPDATE="--update";;
  esac
done

set -euxo pipefail

DATA_DIR=$(mktemp -d -t ness3-data-dir.XXXXXX)
WALLET_DIR="${DATA_DIR}/wallets"

if [[ ! "$DATA_DIR" ]]; then
  echo "Could not create temp dir"
  exit 1
fi

# Compile the ness3 node
# We can't use "go run" because this creates two processes which doesn't allow us to kill it at the end
echo "compiling mdl"
go build -o "$BINARY" $PWD/cmd/ness3/ness3.go

# Run ness3 node with pinned blockchain database
echo "starting mdl ($PWD/mdl-integration) node in background with http listener on $HOST"

$PWD/mdl-integration -disable-networking=true \
                      -genesis-signature 05d4045854103f8a8938bb701cc4101c38942a180ba02d328d6f880bf37b387c47e95813d061f94bdf5d894bfebf17f933c5fc92fc9d010480765257c3d19d9b00 \
                      -genesis-address 24GJTLPMoz61sV4J4qg1n14x5qqDwXqyJJy \
                      -blockchain-public-key 02933015bd2fa1e0a885c05fb08eb7c647bf8c3188ed5120b51d0d09ccaf525036 \
                      -db-path=./src/api/integration/testdata/blockchain-180.db \
                      -peerlist-url https://nodes.privateness.network/blockchain/peers.txt
                      -web-interface-addr=$IP_ADDR \
                      -web-interface-port=$PORT \
                      -download-peerlist=false \
                      -db-read-only=true \
                      -rpc-interface=true \
                      -launch-browser=false \
                      -data-dir="$DATA_DIR" \
                      -enable-wallet-api=true \
                      -wallet-dir="$WALLET_DIR" \
                       $DISABLE_CSRF &
MDL_PID=$!

echo "mdl node pid=$MDL_PID"

echo "sleeping for startup"
sleep 5
echo "done sleeping"

set +e

if [[ -z $TEST || $TEST = "api" ]]; then

MDL_INTEGRATION_TESTS=1 MDL_INTEGRATION_TEST_MODE=$MODE MDL_NODE_HOST=$HOST go test ./src/api/integration/... $UPDATE -timeout=60s -v

API_FAIL=$?

MDL_INTEGRATION_TESTS=1 MDL_INTEGRATION_TEST_MODE=$MODE_LIVE MDL_NODE_HOST=$HOST go test ./src/api/integration/... -timeout=60s -v

API_FAIL_LIVE=$?

fi

if [[ -z $TEST  || $TEST = "cli" ]]; then

MDL_INTEGRATION_TESTS=1 MDL_INTEGRATION_TEST_MODE=$MODE RPC_ADDR=$RPC_ADDR go test ./src/api/cli/integration/... $UPDATE -timeout=60s -v

CLI_FAIL=$?

MDL_INTEGRATION_TESTS=1 MDL_INTEGRATION_TEST_MODE=$MODE_LIVE RPC_ADDR=$RPC_ADDR go test ./src/api/cli/integration/... $UPDATE -timeout=60s -v

CLI_FAIL_LIVE=$?


fi


echo "shutting down ness3 node"

# Shutdown ness3 node
kill -9 $MDL_PID
wait $MDL_PID

rm "$BINARY"


if [[ (-z $TEST || $TEST = "api") && $API_FAIL -ne 0 ]]; then
  exit $API_FAIL
elif [[ (-z $TEST || $TEST = "api") && $API_FAIL_LIVE -ne 0 ]]; then
    exit $API_FAIL_LIVE
elif [[ (-z $TEST || $TEST = "cli") && $CLI_FAIL -ne 0 ]]; then
  exit $CLI_FAIL
elif [[ (-z $TEST || $TEST = "cli") && $CLI_FAIL_LIVE -ne 0 ]]; then
  exit $CLI_FAIL_LIVE
else
  exit 0
fi
