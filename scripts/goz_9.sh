#!/bin/bash
watcher --tmRPC "tcp://rpc.blockngine.io:26657" --rabbitMQ "$RABBITMQ" --zone blockngine-ibc &
watcher --tmRPC "tcp://goz-1.bitsong.network:26657" --rabbitMQ "$RABBITMQ" --zone bitsong-goz-1 &
watcher --tmRPC "tcp://ibc.bharvest.io:26657" --rabbitMQ "$RABBITMQ" --zone B-Harvest &
watcher --tmRPC "tcp://blogchain.xyz:26657" --rabbitMQ "$RABBITMQ" --zone BlogChain &
watcher --tmRPC "tcp://52.247.127.40:26657" --rabbitMQ "$RABBITMQ" --zone aynchain &
watcher --tmRPC "tcp://goz.audit.one:26657" --rabbitMQ "$RABBITMQ" --zone audit.one &
watcher --tmRPC "tcp://bombers.dokia.cloud:26657" --rabbitMQ "$RABBITMQ" --zone atomic-bombers &
watcher --tmRPC "tcp://goz.armyids.com:26657" --rabbitMQ "$RABBITMQ" --zone armyids &
watcher --tmRPC "tcp://116.203.208.175:26657" --rabbitMQ "$RABBITMQ" --zone ape_smash &
watcher --tmRPC "tcp://n1.anonstake:26657" --rabbitMQ "$RABBITMQ" --zone anon-chain &
wait