#!/bin/bash
watcher --tmRPC "ws://goz.01node.com:26657/websocket" --rabbitMQ "$RABBITMQ" --zone 01node &
watcher --tmRPC "ws://achain.nodeateam.com:26657/websocket" --rabbitMQ "$RABBITMQ" --zone ateam &
watcher --tmRPC "ws://capychain.com:26657/websocket" --rabbitMQ "$RABBITMQ" --zone capychain &
watcher --tmRPC "ws://n01.gozlira.com:26657/websocket" --rabbitMQ "$RABBITMQ" --zone chainylira &
watcher --tmRPC "ws://51.178.119.163:26657/websocket" --rabbitMQ "$RABBITMQ" --zone coco-post-chain &
watcher --tmRPC "ws://goz.val.network:26657/websocket" --rabbitMQ "$RABBITMQ" --zone Compass &
watcher --tmRPC "ws://testnet.dawns.world:26657/websocket" --rabbitMQ "$RABBITMQ" --zone dawnsworld &
watcher --tmRPC "ws://ibc01.dokia.cloud:26657/websocket" --rabbitMQ "$RABBITMQ" --zone dokia &
watcher --tmRPC "ws://goz.dos.network:26657/websocket" --rabbitMQ "$RABBITMQ" --zone dos-ibc &
watcher --tmRPC "ws://dropschain.com:26657/websocket" --rabbitMQ "$RABBITMQ" --zone dropschain &
wait