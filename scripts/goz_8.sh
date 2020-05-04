#!/bin/bash
watcher --tmRPC "tcp://176.9.238.157:26657" --rabbitMQ "$RABBITMQ" --zone chainlayer &
watcher --tmRPC "tcp://goz.chainflow.io:26657" --rabbitMQ "$RABBITMQ" --zone Chainflow &
watcher --tmRPC "http://goz.chainapsis.com:80" --rabbitMQ "$RABBITMQ" --zone chainapsis-1 &
watcher --tmRPC "tcp://wasntus.goz.certus.one:26657" --rabbitMQ "$RABBITMQ" --zone it-wasnt-us &
watcher --tmRPC "tcp://3.130.208.130:26657" --rabbitMQ "$RABBITMQ" --zone cat &
watcher --tmRPC "tcp://capychain.com:26657" --rabbitMQ "$RABBITMQ" --zone capychain &
watcher --tmRPC "tcp://178.128.186.74:26657" --rabbitMQ "$RABBITMQ" --zone byzantine-gos &
watcher --tmRPC "tcp://95.216.198.111:26657" --rabbitMQ "$RABBITMQ" --zone burek &
watcher --tmRPC "tcp://149.248.61.193:26657" --rabbitMQ "$RABBITMQ" --zone brankochain &
watcher --tmRPC "tcp://ibc.blockscape.network:26657" --rabbitMQ "$RABBITMQ" --zone NoChainNoGain-1000 &
wait