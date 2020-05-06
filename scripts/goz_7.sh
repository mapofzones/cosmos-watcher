#!/bin/bash
watcher --tmRPC "tcp://goz.everstake.one:26657" --rabbitMQ "$RABBITMQ" --zone everstakechain &
watcher --tmRPC "tcp://52.231.28.219:26657" --rabbitMQ "$RABBITMQ" --zone dunhillchain &
watcher --tmRPC "tcp://goz.dos.network:26657" --rabbitMQ "$RABBITMQ" --zone dos-ibc &
watcher --tmRPC "tcp://3.21.156.79:26657" --rabbitMQ "$RABBITMQ" --zone disraptorchain &
watcher --tmRPC "http://goz.desmos.network:80" --rabbitMQ "$RABBITMQ" --zone morpheus-goz-1a &
watcher --tmRPC "tcp://ibc.defending.network:26657" --rabbitMQ "$RABBITMQ" --zone defending-network &
watcher --tmRPC "tcp://goz.dev.datachain.jp:26657" --rabbitMQ "$RABBITMQ" --zone dcz &
watcher --tmRPC "tcp://testnet.dawns.world:26657" --rabbitMQ "$RABBITMQ" --zone dawnsworld &
watcher --tmRPC "http://cyberdevs.cyberd.ai:26657" --rabbitMQ "$RABBITMQ" --zone cyberdevs &
watcher --tmRPC "tcp://47.240.29.196:26657" --rabbitMQ "$RABBITMQ" --zone crazyzoo &
wait