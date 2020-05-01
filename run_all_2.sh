#!/bin/bash
watcher --tmRPC "ws://goz.irisnet.org:26657/websocket" --rabbitMQ "$RABBITMQ" --zone irishub &
watcher --tmRPC "ws://49.12.106.6:26657/websocket" --rabbitMQ "$RABBITMQ" --zone isillienchain &
watcher --tmRPC "ws://157.230.255.202:26657/websocket" --rabbitMQ "$RABBITMQ" --zone kappa &
watcher --tmRPC "ws://ibc-alpha.kava.io:26657/websocket" --rabbitMQ "$RABBITMQ" --zone kava-ibc &
watcher --tmRPC "ws://goz.konstellation.tech:26657/websocket" --rabbitMQ "$RABBITMQ" --zone konstellation &
watcher --tmRPC "ws://3.211.57.24:26657/websocket" --rabbitMQ "$RABBITMQ" --zone mallowchain &
watcher --tmRPC "ws://gozmelea.mycryptobets.com:26657/websocket" --rabbitMQ "$RABBITMQ" --zone melea-11 &
watcher --tmRPC "ws://mmmh.sytes.net:26657/websocket" --rabbitMQ "$RABBITMQ" --zone mmmh.sytes.net &
watcher --tmRPC "ws://ibc-alpha.desmos.network:26657/websocket" --rabbitMQ "$RABBITMQ" --zone morpheus-ibc-3000 &
watcher --tmRPC "ws://ibct01.newroad.network:26657/websocket" --rabbitMQ "$RABBITMQ" --zone newroadchain &
wait