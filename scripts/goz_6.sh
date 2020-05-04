#!/bin/bash
watcher --tmRPC "tcp://goz.jptpool.com:26657" --rabbitMQ "$RABBITMQ" --zone gemstone &
watcher --tmRPC "tcp://ibc.freeflix.media:26657" --rabbitMQ "$RABBITMQ" --zone freeflix-media-hub &
watcher --tmRPC "tcp://18.217.240.174:26657" --rabbitMQ "$RABBITMQ" --zone finalbattlechain &
watcher --tmRPC "tcp:/goz-ibc.figment.network:26657" --rabbitMQ "$RABBITMQ" --zone figment &
watcher --tmRPC "tcp://fetch-goz.fetch.ai:26657" --rabbitMQ "$RABBITMQ" --zone fetchBeacon &
watcher --tmRPC "tcp://15.236.69.21:26657" --rabbitMQ "$RABBITMQ" --zone ublochain &
watcher --tmRPC "tcp://35.209.174.13:26657" --rabbitMQ "$RABBITMQ" --zone stateset &
watcher --tmRPC "tcp://3.22.194.241:26657" --rabbitMQ "$RABBITMQ" --zone EVM.Protofire.io &
watcher --tmRPC "tcp://goz.everstake.one:26657" --rabbitMQ "$RABBITMQ" --zone everstakechain &
watcher --tmRPC "tcp://52.231.28.219:26657" --rabbitMQ "$RABBITMQ" --zone dunhillchain &
wait