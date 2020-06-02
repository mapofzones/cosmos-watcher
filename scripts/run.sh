#!/bin/sh


# get last processed block
rpcQuery="${rpc/tcp:\/\//http:\/\/}"

export chain_id=$(curl --connect-timeout 3 -ss $rpcQuery/status | jq .result.node_info.network -r)

if [ "$chain_id" == "" ]; then
    echo "unable to fetch chain_id for $rpcQuery"
    sleep 600
    exit 1
fi

height=$(curl -H 'Content-Type: application/json' \
 -X POST -ss -H "x-hasura-admin-secret: $hasura_secret" \
 --data '{"query":"{blocks_log(where: {zone: {_eq: \"'"$chain_id"'\"}}) {last_processed_block}}"}' $graphql \
  | jq .data.blocks_log[].last_processed_block)

if [ "$height" == "" ]; then
    height=0
fi

# increment height since we need to start getting blocks from last_processed_height +1
let "height=height+1"
export height

echo "starting watcher for $chain_id on $rpc at height: $height"
/app/watcher ; sleep 600