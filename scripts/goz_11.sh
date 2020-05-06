#!/bin/bash
watcher --tmRPC "tcp://goz.wetez.io:26657" --rabbitMQ "$RABBITMQ" --zone Wetez &
watcher --tmRPC "tcp://47.245.35.172:26657" --rabbitMQ "$RABBITMQ" --zone tom &
watcher --tmRPC "tcp://34.66.86.162:26657" --rabbitMQ "$RABBITMQ" --zone dking &
watcher --tmRPC "tcp://13.250.207.24:26657" --rabbitMQ "$RABBITMQ" --zone TaidiHub &
watcher --tmRPC "tcp://ibc.staking.fund:26657" --rabbitMQ "$RABBITMQ" --zone stakingfund &
watcher --tmRPC "tcp://5.181.51.80:26657" --rabbitMQ "$RABBITMQ" --zone stakesstonechain &
watcher --tmRPC "tcp://18.136.225.71:26657" --rabbitMQ "$RABBITMQ" --zone SNZPoolHub &
watcher --tmRPC "tcp://rvc.novy.pw:26657" --rabbitMQ "$RABBITMQ" --zone rvc-1 &
watcher --tmRPC "tcp://goz.newroad.network:26657" --rabbitMQ "$RABBITMQ" --zone newroadchain &
watcher --tmRPC "tcp://35.193.176.142:26657" --rabbitMQ "$RABBITMQ" --zone universe &
wait