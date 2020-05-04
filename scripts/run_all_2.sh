#!/bin/bash
watcher --tmRPC "ws://ibc-alpha.kava.io:26657/websocket" --rabbitMQ "$RABBITMQ" --zone kava-ibc &
watcher --tmRPC "ws://goz.konstellation.tech:26657/websocket" --rabbitMQ "$RABBITMQ" --zone konstellation &
watcher --tmRPC "ws://3.211.57.24:26657/websocket" --rabbitMQ "$RABBITMQ" --zone mallowchain &
watcher --tmRPC "ws://gozmelea.mycryptobets.com:26657/websocket" --rabbitMQ "$RABBITMQ" --zone melea-11 &
watcher --tmRPC "ws://ibc-alpha.desmos.network:26657/websocket" --rabbitMQ "$RABBITMQ" --zone morpheus-ibc-3000 &
watcher --tmRPC "ws://ibct01.newroad.network:26657/websocket" --rabbitMQ "$RABBITMQ" --zone newroadchain &
watcher --tmRPC "ws://167.179.104.210:26657/websocket" --rabbitMQ "$RABBITMQ" --zone nibiru-ibc &
watcher --tmRPC "ws://ibc.blockscape.network:26657/websocket" --rabbitMQ "$RABBITMQ" --zone nochainnogain &
watcher --tmRPC "ws://13.231.12.191:26657/websocket" --rabbitMQ "$RABBITMQ" --zone okchain &
watcher --tmRPC "ws://ibc.ping.pub:26657/websocket" --rabbitMQ "$RABBITMQ" --zone ping-ibc &
wait