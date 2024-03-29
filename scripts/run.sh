#!/bin/bash

new_height=""

if [ "$graphql" != "" ]; then
    height=$(curl -H 'Content-Type: application/json' \
     -X POST -ss -H "x-hasura-admin-secret: $hasura_secret" \
     --data '{"query":"{blocks_log(where: {zone: {_eq: \"'"$chain_id"'\"}}) {last_processed_block}}"}' $graphql \
      | jq .data.blocks_log[].last_processed_block)
fi

if [ "$new_height" == "" ]; then
    new_height="$height"
    if [ "$new_height" == "" ]; then
        new_height=0
    fi
fi

height="$new_height"
# increment height since we need to start getting blocks from last_processed_height +1
let "height=height+1"
export height


if [ "$chain_id" != "" ]; then
    setTimeOffset() {
      case "$OSTYPE" in
        solaris*)
          echo "SOLARIS" ;;
        darwin*)
          echo "OSX"
          timeOffset=$(date -v-15M "+%Y-%m-%dT%H:%M:%S") #add %3N as we want millisecond too
          ;;
        linux*)
          echo "LINUX"
          timeOffset=$(date -d '0 days 15 mins ago' +%Y-%m-%dT%H:%M:%S)
          ;;
        bsd*)     echo "BSD" ;;
        msys*)    echo "WINDOWS" ;;
        cygwin*)  echo "ALSO WINDOWS" ;;
        *)        echo "unknown: $OSTYPE" ;;
      esac
    }
    setTimeOffset

    resp=$(curl -H 'Content-Type: application/json' \
     -X POST -ss -H "x-hasura-admin-secret: $hasura_secret" \
     --data '{"query":"{zone_nodes(where: {zone: {_eq: \"'"$chain_id"'\"}, is_alive: {_eq: true}, tx_index: {_eq: \"on\"}, _or: [{last_checked_at: {_lt: \"'"$timeOffset"'\"}}, {last_block_height: {_gt: '"$height"'}}], earliest_block_height: {_lte: '"$height"'}}, order_by: {last_checked_at: desc}) {rpc_addr}}"}' $graphql \
      | jq .data.zone_nodes[].rpc_addr)

    if [ ${#resp[@]} -eq 0 ] || [ ${#resp[@]} -eq 1 ] && [ ${#resp[0]} -eq 0 ] ; then
      resp=$(curl -H 'Content-Type: application/json' \
           -X POST -ss -H "x-hasura-admin-secret: $hasura_secret" \
           --data '{"query":"{zone_nodes(where: {zone: {_eq: \"'"$chain_id"'\"}, is_alive: {_eq: true}}, order_by: {last_checked_at: desc}) {rpc_addr}}"}' $graphql \
            | jq .data.zone_nodes[].rpc_addr)
    fi

    x=0
    declare -a rpcs
    while IFS=$'\n' read -ra ADDR; do
        for i in "${ADDR[@]}"; do
            rpcs=(${rpcs[@]} "$i")
        done
        x=$(( $x + 1 ))
    done <<< "$resp"

    rand=$[$RANDOM % ${#rpcs[@]}]
    rpc="${rpcs[$rand]}"
    rpc="${rpc%\"}" # remove the suffix "
    rpc="${rpc#\"}" # remove the prefix "
    export rpc
fi

if [ "$rpc" == "" ]; then
    echo "Unable to fetch rpc for chain_id: $chain_id"
    sleep 600
    exit 1
fi

echo "starting watcher for $chain_id on $rpc at height: $height"
/app/watcher ; sleep 60