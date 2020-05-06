#!/bin/bash
watcher --tmRPC "tcp://molecule.adri.co:26657" --rabbitMQ "$RABBITMQ" --zone moleculechain &
watcher --tmRPC "tcp://45.79.207.112:26657" --rabbitMQ "$RABBITMQ" --zone microtick-ibc &
watcher --tmRPC "tcp://35.236.168.104:26657" --rabbitMQ "$RABBITMQ" --zone irishub-goz &
watcher --tmRPC "tcp://35.222.132.154:26657" --rabbitMQ "$RABBITMQ" --zone seen &
watcher --tmRPC "tcp://val1.goz.enigma.co:26657" --rabbitMQ "$RABBITMQ" --zone enigma-goz &
watcher --tmRPC "tcp://dropschain.com:26657" --rabbitMQ "$RABBITMQ" --zone dropschain &
watcher --tmRPC "tcp://goz.val.network:26657" --rabbitMQ "$RABBITMQ" --zone Compass &
watcher --tmRPC "tcp://achain.nodeateam.com:26657" --rabbitMQ "$RABBITMQ" --zone achain &
watcher --tmRPC "ws://35.233.155.199:26657/websocket" --rabbitMQ "$RABBITMQ" --zone gameofzoneshub-1a &
wait