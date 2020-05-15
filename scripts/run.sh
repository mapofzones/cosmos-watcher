#!/bin/sh


# get last processed block

rpcQuery="${rpc/tcp:\/\//http:\/\/}"

export chain_id=$(curl --connect-timeout 3 -ss $rpcQuery/status | jq .result.node_info.network -r)

if [ "$chain_id" == "" ]; then
    sleep 600
    exit 1
fi

export height=$(curl -H 'Content-Type: application/json' \
 -X POST -ss --data  '{"query":"{blocks_log(where: {chain_id: {_eq: \"'"$chain_id"'\"}}) {last_processed_block}}"}' $GRAPHQL | jq .data.blocks_log[].last_processed_block)

if [ "$height" == "" ]; then
    export height=1
fi

echo "Try connect to $rpc on $chain_id with $height"

/app/watcher ; sleep 600