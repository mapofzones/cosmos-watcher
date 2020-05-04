#!/bin/bash
watcher --tmRPC "tcp://goz.val.network:26657" --rabbitMQ "$RABBITMQ" --zone Compass &
watcher --tmRPC "tcp://achain.nodeateam.com:26657" --rabbitMQ "$RABBITMQ" --zone achain &
watcher --tmRPC "ws://35.233.155.199:26657/websocket" --rabbitMQ "$RABBITMQ" --zone gameofzoneshub-1a &
wait