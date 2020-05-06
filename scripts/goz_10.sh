#!/bin/bash
watcher --tmRPC "tcp://bombers.dokia.cloud:26657" --rabbitMQ "$RABBITMQ" --zone atomicbombers &
watcher --tmRPC "tcp://bombers.dokia.cloud:26657" --rabbitMQ "$RABBITMQ" --zone atomic-bombers &
watcher --tmRPC "tcp://goz.armyids.com:26657" --rabbitMQ "$RABBITMQ" --zone armyids &
watcher --tmRPC "tcp://116.203.208.175:26657" --rabbitMQ "$RABBITMQ" --zone ape_smash &
watcher --tmRPC "tcp://116.203.245.68:26657" --rabbitMQ "$RABBITMQ" --zone anon-chain &
watcher --tmRPC "tcp://goz.aneka.io:26657" --rabbitMQ "$RABBITMQ" --zone aneka &
watcher --tmRPC "tcp://goz.cosmos.alphavirtual.com:26657" --rabbitMQ "$RABBITMQ" --zone avnet &
watcher --tmRPC "tcp://159.89.183.144:26657" --rabbitMQ "$RABBITMQ" --zone akashian &
watcher --tmRPC "tcp://138.197.157.152:26657" --rabbitMQ "$RABBITMQ" --zone aib-goz-1 &
watcher --tmRPC "tcp://178.128.254.143:26657" --rabbitMQ "$RABBITMQ" --zone agoric-goz-1.0.0 &
wait