#!/bin/bash
watcher --tmRPC "tcp://goz.dos.network:26657" --rabbitMQ "$RABBITMQ" --zone dos-ibc &
watcher --tmRPC "tcp://3.21.156.79:26657" --rabbitMQ "$RABBITMQ" --zone disraptorchain &
watcher --tmRPC "http://goz.desmos.network:80" --rabbitMQ "$RABBITMQ" --zone morpheus-goz &
watcher --tmRPC "tcp://ibc.defending.network:26657" --rabbitMQ "$RABBITMQ" --zone defending-network &
watcher --tmRPC "tcp://goz.dev.datachain.jp:26657" --rabbitMQ "$RABBITMQ" --zone dcz &
watcher --tmRPC "tcp://testnet.dawns.world:26657" --rabbitMQ "$RABBITMQ" --zone dawnsworld &
watcher --tmRPC "tcp://47.240.29.196:26657" --rabbitMQ "$RABBITMQ" --zone crazyzoo &
watcher --tmRPC "tcp://ibc.cosmoon.org:26657" --rabbitMQ "$RABBITMQ" --zone cosmoon-testnet &
watcher --tmRPC "tcp://ibc.cosmiccompass.io:26657" --rabbitMQ "$RABBITMQ" --zone coco-post-chain &
watcher --tmRPC "tcp://fedzone.chorus.one:26657" --rabbitMQ "$RABBITMQ" --zone fedzone-1 &
wait