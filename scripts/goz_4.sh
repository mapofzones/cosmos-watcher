#!/bin/bash
watcher --tmRPC "tcp://193.30.121.61:26657" --rabbitMQ "$RABBITMQ" --zone Node123 &
watcher --tmRPC "tcp://goz.nibiru.network:26657" --rabbitMQ "$RABBITMQ" --zone nibiru-ibc &
watcher --tmRPC "tcp://muzamint.com:26657/" --rabbitMQ "$RABBITMQ" --zone muzamint &
watcher --tmRPC "tcp://goz.modulus.network:26657" --rabbitMQ "$RABBITMQ" --zone modulus-goz-1 &
watcher --tmRPC "tcp://mmmh.sytes.net:26657" --rabbitMQ "$RABBITMQ" --zone mmmh-lazy &
watcher --tmRPC "tcp://45.77.91.232:26657" --rabbitMQ "$RABBITMQ" --zone mintonium &
watcher --tmRPC "tcp://melea.xyz:26657" --rabbitMQ "$RABBITMQ" --zone melea-1111 &
watcher --tmRPC "tcp://validating.for.co.ke:26657" --rabbitMQ "$RABBITMQ" --zone meeseeks &
watcher --tmRPC "tcp://goz.labeleet.com:26657" --rabbitMQ "$RABBITMQ" --zone kugs-030 &
watcher --tmRPC "tcp://goz.konstellation.tech:26657" --rabbitMQ "$RABBITMQ" --zone konstellation &
wait