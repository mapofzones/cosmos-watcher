#!/bin/bash
watcher --tmRPC "ws://goz.everstake.one:26657/websocket" --rabbitMQ "$RABBITMQ" --zone everstakechain &
watcher --tmRPC "ws://51.178.119.162:26657/websocket" --rabbitMQ "$RABBITMQ" --zone freeflix-media-hub &
watcher --tmRPC "ws://goz.jptpool.com:26657/websocket" --rabbitMQ "$RABBITMQ" --zone gemstone &
watcher --tmRPC "ws://goz.kalpatech.co:26657/websocket" --rabbitMQ "$RABBITMQ" --zone goz.kalpatech.co &
watcher --tmRPC "ws://goz.gunray.xyz:26657/websocket" --rabbitMQ "$RABBITMQ" --zone gunchain &
watcher --tmRPC "ws://152.32.135.74:26657/websocket" --rabbitMQ "$RABBITMQ" --zone hashquarkchain &
watcher --tmRPC "ws://ibc-testnet1.bandchain.org:26657/websocket" --rabbitMQ "$RABBITMQ" --zone ibc-band-testnet1 &
watcher --tmRPC "ws://ibc.cosmoon.org:26657/websocket" --rabbitMQ "$RABBITMQ" --zone ibc.cosmoon.org &
watcher --tmRPC "ws://ibc.izo.ro:26657/websocket" --rabbitMQ "$RABBITMQ" --zone ibc.izo.ro &
watcher --tmRPC "ws://fridayco.in:26657/websocket" --rabbitMQ "$RABBITMQ" --zone iqlusionchain &
wait