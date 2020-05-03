#!/bin/bash
watcher --tmRPC "ws://ibc.westaking.io:26657/websocket" --rabbitMQ "$RABBITMQ" --zone ibc.westaking.io &
watcher --tmRPC "ws://goz.irisnet.org:26657/websocket" --rabbitMQ "$RABBITMQ" --zone goz.irisnet.org &
watcher --tmRPC "ws://achain.nodeateam.com:26657/websocket" --rabbitMQ "$RABBITMQ" --zone achain.nodeateam.com &
watcher --tmRPC "ws://n01.gozlira.com:26657/websocket" --rabbitMQ "$RABBITMQ" --zone n01.gozlira.com &
watcher --tmRPC "ws://ibc01.dokia.cloud:26657/websocket" --rabbitMQ "$RABBITMQ" --zone ibc01.dokia.cloud &
watcher --tmRPC "ws://goz.dos.network:26657/websocket" --rabbitMQ "$RABBITMQ" --zone goz.dos.network &
watcher --tmRPC "ws://152.32.135.74:26657/websocket" --rabbitMQ "$RABBITMQ" --zone 152.32.135.74 &
watcher --tmRPC "ws://fridayco.in:26657/websocket" --rabbitMQ "$RABBITMQ" --zone fridayco.in &
watcher --tmRPC "ws://95.217.135.90:26657/websocket" --rabbitMQ "$RABBITMQ" --zone 95.217.135.90 &
watcher --tmRPC "ws://ibc-testnet1.bandchain.org:26657/websocket" --rabbitMQ "$RABBITMQ" --zone ibc-testnet1.bandchain.org &
wait