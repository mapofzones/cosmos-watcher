#!/bin/bash
watcher --tmRPC "ws://ibc.vitwit.in:26657/websocket" --rabbitMQ "$RABBITMQ" --zone vitwitchain &
watcher --tmRPC "ws://chain.exchange-fees.com:26657/websocket" --rabbitMQ "$RABBITMQ" --zone vostok-1 &
watcher --tmRPC "ws://ibc.westaking.io:26657/websocket" --rabbitMQ "$RABBITMQ" --zone westaking &
wait