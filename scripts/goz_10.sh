#!/bin/bash
watcher --tmRPC "tcp://goz.aneka.io:26657" --rabbitMQ "$RABBITMQ" --zone aneka &
watcher --tmRPC "tcp://goz.cosmos.alphavirtual.com:26657" --rabbitMQ "$RABBITMQ" --zone avnet &
watcher --tmRPC "tcp://159.89.183.144:26657" --rabbitMQ "$RABBITMQ" --zone akashian &
watcher --tmRPC "tcp://138.197.157.152:26657" --rabbitMQ "$RABBITMQ" --zone aib-goz-1 &
watcher --tmRPC "tcp://178.128.254.143:26657" --rabbitMQ "$RABBITMQ" --zone agoric-goz-1.0.0 &
watcher --tmRPC "tcp://goz.wetez.io:26657" --rabbitMQ "$RABBITMQ" --zone Wetez &
watcher --tmRPC "tcp://47.245.35.172:26657" --rabbitMQ "$RABBITMQ" --zone tom &
watcher --tmRPC "tcp://34.66.86.162:26657" --rabbitMQ "$RABBITMQ" --zone dking &
watcher --tmRPC "tcp://13.250.207.24:26657" --rabbitMQ "$RABBITMQ" --zone TaidiHub &
watcher --tmRPC "tcp://ibc.staking.fund:26657" --rabbitMQ "$RABBITMQ" --zone stakingfund &
wait