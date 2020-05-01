#!/bin/bash
watcher --tmRPC "ws://supernova.commonwealth.im:26657/websocket" --rabbitMQ "$RABBITMQ" --zone supernova &
watcher --tmRPC "ws://ibc.umbrellavalidator.com:26657/websocket" --rabbitMQ "$RABBITMQ" --zone umbrellachain &
watcher --tmRPC "ws://173.249.12.108:26657/websocket" --rabbitMQ "$RABBITMQ" --zone vipnamai &
watcher --tmRPC "ws://ibc1.vitwit.in:26657/websocket" --rabbitMQ "$RABBITMQ" --zone vitwitchain-1 &
watcher --tmRPC "ws://ibc.vitwit.in:26657/websocket" --rabbitMQ "$RABBITMQ" --zone vitwitchain &
watcher --tmRPC "ws://chain.exchange-fees.com:26657/websocket" --rabbitMQ "$RABBITMQ" --zone vostok-1 &
watcher --tmRPC "ws://ibc.westaking.io:26657/websocket" --rabbitMQ "$RABBITMQ" --zone westaking &
wait