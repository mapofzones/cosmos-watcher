#!/bin/bash
watcher --tmRPC "ws://goz.everstake.one:26657/websocket" --rabbitMQ "$RABBITMQ" --zone everstakechain &
watcher --tmRPC "ws://51.178.119.162:26657/websocket" --rabbitMQ "$RABBITMQ" --zone freeflix-media-hub &
watcher --tmRPC "ws://goz.jptpool.com:26657/websocket" --rabbitMQ "$RABBITMQ" --zone gemstone &
watcher --tmRPC "ws://goz.gunray.xyz:26657/websocket" --rabbitMQ "$RABBITMQ" --zone gunchain &
watcher --tmRPC "ws://152.32.135.74:26657/websocket" --rabbitMQ "$RABBITMQ" --zone hashquarkchain &
watcher --tmRPC "ws://ibc-testnet1.bandchain.org:26657/websocket" --rabbitMQ "$RABBITMQ" --zone ibc-band-testnet1 &
watcher --tmRPC "ws://fridayco.in:26657/websocket" --rabbitMQ "$RABBITMQ" --zone iqlusionchain &
watcher --tmRPC "ws://goz.irisnet.org:26657/websocket" --rabbitMQ "$RABBITMQ" --zone irishub &
watcher --tmRPC "ws://49.12.106.6:26657/websocket" --rabbitMQ "$RABBITMQ" --zone isillienchain &
watcher --tmRPC "ws://157.230.255.202:26657/websocket" --rabbitMQ "$RABBITMQ" --zone kappa &
wait