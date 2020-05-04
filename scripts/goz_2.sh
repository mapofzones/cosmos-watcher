#!/bin/bash
watcher --tmRPC "ws://ibc.blockscape.network:26657/websocket" --rabbitMQ "$RABBITMQ" --zone ibc.blockscape.network &
watcher --tmRPC "ws://shitcoincasinos.com:26657/websocket" --rabbitMQ "$RABBITMQ" --zone shitcoincasinos.com &
watcher --tmRPC "ws://ibc.umbrellavalidator.com:26657/websocket" --rabbitMQ "$RABBITMQ" --zone ibc.umbrellavalidator.com &
watcher --tmRPC "ws://ibc1.vitwit.in:26657/websocket" --rabbitMQ "$RABBITMQ" --zone ibc1.vitwit.in &
watcher --tmRPC "ws://ibc.vitwit.in:26657/websocket" --rabbitMQ "$RABBITMQ" --zone ibc.vitwit.in &
watcher --tmRPC "ws://80.64.211.64:26657/websocket" --rabbitMQ "$RABBITMQ" --zone 80.64.211.64 &
watcher --tmRPC "ws://goz.val.network:26657/websocket" --rabbitMQ "$RABBITMQ" --zone goz.val.network &
watcher --tmRPC "ws://testnet.dawns.world:26657/websocket" --rabbitMQ "$RABBITMQ" --zone testnet.dawns.world &
watcher --tmRPC "ws://goz.konstellation.tech:26657/websocket" --rabbitMQ "$RABBITMQ" --zone goz.konstellation.tech &
watcher --tmRPC "ws://167.179.104.210:26657/websocket" --rabbitMQ "$RABBITMQ" --zone 167.179.104.210 &
wait