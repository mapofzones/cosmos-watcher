#!/bin/bash
watcher --tmRPC "ws://setan.ml:26657/websocket" --rabbitMQ "$RABBITMQ" --zone setan.ml &
watcher --tmRPC "ws://goz.01node.com:26657/websocket" --rabbitMQ "$RABBITMQ" --zone goz.01node.com &
watcher --tmRPC "ws://capychain.com:26657/websocket" --rabbitMQ "$RABBITMQ" --zone capychain.com &
watcher --tmRPC "ws://51.178.119.163:26657/websocket" --rabbitMQ "$RABBITMQ" --zone 51.178.119.163 &
watcher --tmRPC "ws://dropschain.com:26657/websocket" --rabbitMQ "$RABBITMQ" --zone dropschain.com &
watcher --tmRPC "ws://goz.everstake.one:26657/websocket" --rabbitMQ "$RABBITMQ" --zone goz.everstake.one &
watcher --tmRPC "ws://51.178.119.162:26657/websocket" --rabbitMQ "$RABBITMQ" --zone 51.178.119.162 &
watcher --tmRPC "ws://goz.jptpool.com:26657/websocket" --rabbitMQ "$RABBITMQ" --zone goz.jptpool.com &
watcher --tmRPC "ws://goz.gunray.xyz:26657/websocket" --rabbitMQ "$RABBITMQ" --zone goz.gunray.xyz &
watcher --tmRPC "ws://49.12.106.6:26657/websocket" --rabbitMQ "$RABBITMQ" --zone 49.12.106.6 &
wait