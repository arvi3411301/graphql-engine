#!/usr/bin/env bash

set -euo pipefail
IFS=$'\n\t'
CONSOLE_ROOT="${BASH_SOURCE[0]%/*}/../console"

# start graphql-engine
/build/_server_output/graphql-engine \
    --database-url postgres://gql_test@localhost:5432/gql_test serve &

# start cli
cd /tmp
/build/_server_output/hasura-linux-amd64 init --directory gql-test && cd gql-test
/build/_server_output/hasura-linux-amd64 console --no-browser &

cd "$CONSOLE_ROOT" 
# test console
make test