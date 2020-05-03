#!/bin/bash
watcher --tmRPC "ws://167.179.104.210:26657/websocket" --rabbitMQ "$RABBITMQ" --zone nibiru-ibc &
watcher --tmRPC "ws://ibc.blockscape.network:26657/websocket" --rabbitMQ "$RABBITMQ" --zone nochainnogain &
watcher --tmRPC "ws://13.231.12.191:26657/websocket" --rabbitMQ "$RABBITMQ" --zone okchain &
watcher --tmRPC "ws://ibc.ping.pub:26657/websocket" --rabbitMQ "$RABBITMQ" --zone ping-ibc &
watcher --tmRPC "ws://95.217.135.90:26657/websocket" --rabbitMQ "$RABBITMQ" --zone ptpchain &
watcher --tmRPC "ws://shitcoincasinos.com:26657/websocket" --rabbitMQ "$RABBITMQ" --zone pylonchain &
watcher --tmRPC "ws://tnet-csg.c9ret.xyz:26657/websocket" --rabbitMQ "$RABBITMQ" --zone retz80chain &
watcher --tmRPC "ws://setan.ml:26657/websocket" --rabbitMQ "$RABBITMQ" --zone setanchain &
watcher --tmRPC "ws://80.64.211.64:26657/websocket" --rabbitMQ "$RABBITMQ" --zone simplystaking &
watcher --tmRPC "ws://ibc.staking.fund:26657/websocket" --rabbitMQ "$RABBITMQ" --zone stakingfund &
wait