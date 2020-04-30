#!/bin/bash
watcher --tmRPC "ws://95.217.135.90:26657/websocket" --rabbitMQ "$RABBITMQ" --zone ptpchain &
watcher --tmRPC "ws://shitcoincasinos.com:26657/websocket" --rabbitMQ "$RABBITMQ" --zone pylonchain &
watcher --tmRPC "ws://tnet-csg.c9ret.xyz:26657/websocket" --rabbitMQ "$RABBITMQ" --zone retz80chain &
watcher --tmRPC "ws://setan.ml:26657/websocket" --rabbitMQ "$RABBITMQ" --zone setanchain &
watcher --tmRPC "ws://80.64.211.64:26657/websocket" --rabbitMQ "$RABBITMQ" --zone simplystaking &
watcher --tmRPC "ws://ibc.staking.fund:26657/websocket" --rabbitMQ "$RABBITMQ" --zone stakingfund &
watcher --tmRPC "ws://supernova.commonwealth.im:26657/websocket" --rabbitMQ "$RABBITMQ" --zone supernova &
watcher --tmRPC "ws://ibc.umbrellavalidator.com:26657/websocket" --rabbitMQ "$RABBITMQ" --zone umbrellachain &
watcher --tmRPC "ws://173.249.12.108:26657/websocket" --rabbitMQ "$RABBITMQ" --zone vipnamai &
watcher --tmRPC "ws://ibc1.vitwit.in:26657/websocket" --rabbitMQ "$RABBITMQ" --zone vitwitchain-1 &
wait