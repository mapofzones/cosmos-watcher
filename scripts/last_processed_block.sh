#!/bin/bash

# curl wants http 
rpc="${rpc/tcp:\/\//http:\/\/}"
#get chain id
echo $chain_id
chain_id=$(curl -ss $rpc/status | jq .result.node_info.network -r)

# get last processed block
block=$(curl -H 'Content-Type: application/json' \
 -X POST -ss --data  '{"query":"{blocks_log(where: {chain_id: {_eq: \"'"$chain_id"'\"}}) {last_processed_block}}"}' $GRAPHQL | jq .data.blocks_log[].last_processed_block)

if [ "$block" == "" ]; then
    echo 1
    exit 0
fi

echo $block