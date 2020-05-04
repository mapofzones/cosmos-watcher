#!/bin/bash
watcher --tmRPC "ws://ibc-testnet1.bandchain.org:26657/websocket" --rabbitMQ "$RABBITMQ" --zone ibc-testnet1.bandchain.org &
watcher --tmRPC "ws://157.230.255.202:26657/websocket" --rabbitMQ "$RABBITMQ" --zone 157.230.255.202 &
watcher --tmRPC "ws://ibc-alpha.kava.io:26657/websocket" --rabbitMQ "$RABBITMQ" --zone ibc-alpha.kava.io &
watcher --tmRPC "ws://ibc-alpha.desmos.network:26657/websocket" --rabbitMQ "$RABBITMQ" --zone ibc-alpha.desmos.network &
watcher --tmRPC "ws://ibc.ping.pub:26657/websocket" --rabbitMQ "$RABBITMQ" --zone ibc.ping.pub &
watcher --tmRPC "ws://3.211.57.24:26657/websocket" --rabbitMQ "$RABBITMQ" --zone 3.211.57.24 &
watcher --tmRPC "ws://ibc.staking.fund:26657/websocket" --rabbitMQ "$RABBITMQ" --zone ibc.staking.fund &
watcher --tmRPC "ws://ibct01.newroad.network:26657/websocket" --rabbitMQ "$RABBITMQ" --zone ibct01.newroad.network &
watcher --tmRPC "ws://13.231.12.191:26657/websocket" --rabbitMQ "$RABBITMQ" --zone 13.231.12.191 &
watcher --tmRPC "ws://95.217.135.90:26657/websocket" --rabbitMQ "$RABBITMQ" --zone 95.217.135.90 &
wait