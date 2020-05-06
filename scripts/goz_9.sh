#!/bin/bash
watcher --tmRPC "tcp://178.128.186.74:26657" --rabbitMQ "$RABBITMQ" --zone byzantine-gos &
watcher --tmRPC "tcp://95.216.198.111:26657" --rabbitMQ "$RABBITMQ" --zone burek &
watcher --tmRPC "tcp://149.248.61.193:26657" --rabbitMQ "$RABBITMQ" --zone brankochain &
watcher --tmRPC "tcp://ibc.blockscape.network:26657" --rabbitMQ "$RABBITMQ" --zone NoChainNoGain-1000 &
watcher --tmRPC "tcp://rpc.blockngine.io:26657" --rabbitMQ "$RABBITMQ" --zone blockngine-ibc &
watcher --tmRPC "tcp://goz-1.bitsong.network:26657" --rabbitMQ "$RABBITMQ" --zone bitsong-goz-1 &
watcher --tmRPC "tcp://ibc.bharvest.io:26657" --rabbitMQ "$RABBITMQ" --zone B-Harvest &
watcher --tmRPC "tcp://blogchain.xyz:26657" --rabbitMQ "$RABBITMQ" --zone BlogChain &
watcher --tmRPC "tcp://52.247.127.40:26657" --rabbitMQ "$RABBITMQ" --zone aynchain &
watcher --tmRPC "tcp://goz.audit.one:26657" --rabbitMQ "$RABBITMQ" --zone audit.one &
wait