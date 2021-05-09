#!/bin/sh

if [ "$chain_id" != "" ]; then
    rpc=$(curl -H 'Content-Type: application/json' \
     -X POST -ss -H "x-hasura-admin-secret: $hasura_secret" \
     --data '{"query":"{zone_nodes(where: {zone: {_eq: \"'"$chain_id"'\"}, is_alive: {_eq: true}}, order_by: {last_checked_at: desc}, limit: 1) {rpc_addr}}"}' $graphql \
      | jq .data.zone_nodes[].rpc_addr)
    rpc="${rpc%\"}" # remove the suffix "
    rpc="${rpc#\"}" # remove the prefix "
    export rpc
fi

if [ "$rpc" == "" ]; then
    echo "Unable to fetch rpc for chain_id: $chain_id"
    sleep 600
    exit 1
fi

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

echo "starting watcher for $chain_id on $rpc at height: $height"
/app/watcher ; sleep 60