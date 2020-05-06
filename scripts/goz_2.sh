#!/bin/bash
watcher --tmRPC "tcp://goz.stakedao.org:26657" --rabbitMQ "$RABBITMQ" --zone stake-capital &
watcher --tmRPC "tcp://goz.source.network:26657" --rabbitMQ "$RABBITMQ" --zone source-xgoz &
watcher --tmRPC "tcp://ananas.alpe1.net:26657" --rabbitMQ "$RABBITMQ" --zone snakey &
watcher --tmRPC "tcp://80.64.211.64:26657" --rabbitMQ "$RABBITMQ" --zone simplystaking &
watcher --tmRPC "tcp://setan.ml:26657" --rabbitMQ "$RABBITMQ" --zone setanchain &
watcher --tmRPC "tcp://one.goz.sentinel.co:26657" --rabbitMQ "$RABBITMQ" --zone sentinel-goz &
watcher --tmRPC "tcp://tnet-csg.c9ret.xyz:26657" --rabbitMQ "$RABBITMQ" --zone retz80chain &
watcher --tmRPC "tcp://regengoz.vaasl.io:26657" --rabbitMQ "$RABBITMQ" --zone regengoz &
watcher --tmRPC "tcp://node.pylons.tech:26657" --rabbitMQ "$RABBITMQ" --zone pylonschain &
watcher --tmRPC "tcp://35.230.42.221:26657" --rabbitMQ "$RABBITMQ" --zone pupu &
wait