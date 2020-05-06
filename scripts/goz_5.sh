#!/bin/bash
watcher --tmRPC "tcp://goz.kiraex.com:10001" --rabbitMQ "$RABBITMQ" --zone kira-1 &
watcher --tmRPC "tcp://213.32.70.133:26657" --rabbitMQ "$RABBITMQ" --zone jtbchain &
watcher --tmRPC "tcp://54.211.26.151:26657" --rabbitMQ "$RABBITMQ" --zone js &
watcher --tmRPC "tcp://joon-chain-goz.cosmostation.io:26657" --rabbitMQ "$RABBITMQ" --zone joon-chain-goz &
watcher --tmRPC "tcp://49.12.106.6:26657" --rabbitMQ "$RABBITMQ" --zone isillienchain &
watcher --tmRPC "tcp://interstation.cosmostation.io:26657" --rabbitMQ "$RABBITMQ" --zone interstation &
watcher --tmRPC "tcp://18.178.211.15:26657" --rabbitMQ "$RABBITMQ" --zone hongo-3 &
watcher --tmRPC "tcp://goz.hashquark.io:26657" --rabbitMQ "$RABBITMQ" --zone hashquarkchain &
watcher --tmRPC "tcp://15.165.120.204:26657" --rabbitMQ "$RABBITMQ" --zone hada &
watcher --tmRPC "tcp://goz.gunray.xyz:26657" --rabbitMQ "$RABBITMQ" --zone gunchain &
wait