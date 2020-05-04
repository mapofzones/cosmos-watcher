#!/bin/bash
watcher --tmRPC "tcp://muzamint.com:26657/" --rabbitMQ "$RABBITMQ" --zone muzamint &
watcher --tmRPC "tcp://goz.modulus.network:26657" --rabbitMQ "$RABBITMQ" --zone modulus-goz-1 &
watcher --tmRPC "tcp://mmmh.sytes.net:26657" --rabbitMQ "$RABBITMQ" --zone mmmh-lazy &
watcher --tmRPC "tcp://45.77.91.232:26657" --rabbitMQ "$RABBITMQ" --zone mintonium &
watcher --tmRPC "tcp://melea.xyz:26657" --rabbitMQ "$RABBITMQ" --zone melea-111 &
watcher --tmRPC "tcp://validating.for.co.ke:26657" --rabbitMQ "$RABBITMQ" --zone meeseeks &
watcher --tmRPC "tcp://goz.labeleet.com:26657" --rabbitMQ "$RABBITMQ" --zone kugs-030 &
watcher --tmRPC "tcp://goz.konstellation.tech:26657" --rabbitMQ "$RABBITMQ" --zone konstellation &
watcher --tmRPC "tcp://goz.kiraex.com:10001" --rabbitMQ "$RABBITMQ" --zone kira-1 &
watcher --tmRPC "tcp://213.32.70.133:26657" --rabbitMQ "$RABBITMQ" --zone jtbchain &
wait