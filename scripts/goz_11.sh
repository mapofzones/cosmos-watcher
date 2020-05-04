#!/bin/bash
watcher --tmRPC "tcp://5.181.51.80:26657" --rabbitMQ "$RABBITMQ" --zone stakesstonechain &
watcher --tmRPC "tcp://18.136.225.71:26657" --rabbitMQ "$RABBITMQ" --zone SNZPoolHub &
watcher --tmRPC "tcp://rvc.novy.pw:26657" --rabbitMQ "$RABBITMQ" --zone rvc-1 &
watcher --tmRPC "tcp://goz.newroad.network:26657" --rabbitMQ "$RABBITMQ" --zone newroadchain &
watcher --tmRPC "tcp://35.193.176.142:26657" --rabbitMQ "$RABBITMQ" --zone universe &
watcher --tmRPC "tcp://45.79.207.112:26657" --rabbitMQ "$RABBITMQ" --zone microtick-ibc &
watcher --tmRPC "tcp://35.236.168.104:26657" --rabbitMQ "$RABBITMQ" --zone irishub-goz &
watcher --tmRPC "tcp://35.222.132.154:26657" --rabbitMQ "$RABBITMQ" --zone seen &
watcher --tmRPC "tcp://val1.goz.enigma.co:26657" --rabbitMQ "$RABBITMQ" --zone enigma-goz &
watcher --tmRPC "tcp://dropschain.com:26657" --rabbitMQ "$RABBITMQ" --zone dropschain &
wait