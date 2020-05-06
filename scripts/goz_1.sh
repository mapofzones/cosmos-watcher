#!/bin/bash
watcher --tmRPC "tcp:stratus.mycryptobets.com:26657" --rabbitMQ "$RABBITMQ" --zone stardust-1111 &
watcher --tmRPC "tcp://goz.starcluster.tech:26657" --rabbitMQ "$RABBITMQ" --zone starcluster-1337 &
watcher --tmRPC "tcp://161.35.45.178:26657" --rabbitMQ "$RABBITMQ" --zone stakin &
watcher --tmRPC "tcp://goz2.stake.zone:26657" --rabbitMQ "$RABBITMQ" --zone szchain &
watcher --tmRPC "tcp://stakewolf.com:26657" --rabbitMQ "$RABBITMQ" --zone stakewolf &
watcher --tmRPC "tcp://35.198.125.128:26657" --rabbitMQ "$RABBITMQ" --zone stakematic &
watcher --tmRPC "tcp://goz.cosmos.fish:26657" --rabbitMQ "$RABBITMQ" --zone jellyfish &
watcher --tmRPC "tcp://ibc.staked.cloud:26657" --rabbitMQ "$RABBITMQ" --zone staked-ibc &
watcher --tmRPC "tcp://goz.stakebird.com:26657" --rabbitMQ "$RABBITMQ" --zone stakebird-1 &
watcher --tmRPC "tcp://ibc.stake.sh:26657" --rabbitMQ "$RABBITMQ" --zone nuube-goz &
wait