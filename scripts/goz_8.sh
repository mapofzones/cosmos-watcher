#!/bin/bash
watcher --tmRPC "tcp://ibc.cosmoon.org:26657" --rabbitMQ "$RABBITMQ" --zone cosmoon-testnet &
watcher --tmRPC "tcp://supernova.commonwealth.im:26657" --rabbitMQ "$RABBITMQ" --zone supernova &
watcher --tmRPC "tcp://ibc.cosmiccompass.io:26657" --rabbitMQ "$RABBITMQ" --zone coco-post-chain &
watcher --tmRPC "tcp://fedzone.chorus.one:26657" --rabbitMQ "$RABBITMQ" --zone fedzone-1 &
watcher --tmRPC "tcp://176.9.238.157:26657" --rabbitMQ "$RABBITMQ" --zone chainlayer &
watcher --tmRPC "tcp://goz.chainflow.io:26657" --rabbitMQ "$RABBITMQ" --zone Chainflow &
watcher --tmRPC "http://goz.chainapsis.com:80" --rabbitMQ "$RABBITMQ" --zone chainapsis-1a &
watcher --tmRPC "tcp://wasntus.goz.certus.one:26657" --rabbitMQ "$RABBITMQ" --zone it-wasnt-us &
watcher --tmRPC "tcp://3.130.208.130:26657" --rabbitMQ "$RABBITMQ" --zone cat &
watcher --tmRPC "tcp://capychain.com:26657" --rabbitMQ "$RABBITMQ" --zone capychain &
wait