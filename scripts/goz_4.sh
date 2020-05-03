#!/bin/bash
watcher --tmRPC "ws://supernova.commonwealth.im:26657/websocket" --rabbitMQ "$RABBITMQ" --zone supernova.commonwealth.im &
watcher --tmRPC "ws://173.249.12.108:26657/websocket" --rabbitMQ "$RABBITMQ" --zone 173.249.12.108 &
watcher --tmRPC "ws://chain.exchange-fees.com:26657/websocket" --rabbitMQ "$RABBITMQ" --zone chain.exchange-fees.com &
watcher --tmRPC "ws://ibc.izo.ro:26657/websocket" --rabbitMQ "$RABBITMQ" --zone ibc.izo.ro &
watcher --tmRPC "ws://goz.kalpatech.co:26657/websocket" --rabbitMQ "$RABBITMQ" --zone goz.kalpatech.co &
watcher --tmRPC "ws://goz.01node.com:26657/websocket" --rabbitMQ "$RABBITMQ" --zone goz.01node.com &
watcher --tmRPC "ws://ibc.cosmoon.org:26657/websocket" --rabbitMQ "$RABBITMQ" --zone ibc.cosmoon.org &
watcher --tmRPC "ws://mmmh.sytes.net:26657/websocket" --rabbitMQ "$RABBITMQ" --zone mmmh.sytes.net &
wait