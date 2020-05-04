#!/bin/bash
watcher --tmRPC "tcp://3.21.169.1:26657" --rabbitMQ "$RABBITMQ" --zone zilotchain &
watcher --tmRPC "tcp://ibc.westaking.io:26657" --rabbitMQ "$RABBITMQ" --zone westaking &
watcher --tmRPC "tcp://chain.exchange-fees.com:26657" --rabbitMQ "$RABBITMQ" --zone vostok-1 &
watcher --tmRPC "tcp://173.249.12.108:26657" --rabbitMQ "$RABBITMQ" --zone vipnamai &
watcher --tmRPC "tcp://ibc.vgng.io:26657" --rabbitMQ "$RABBITMQ" --zone vgng-1 &
watcher --tmRPC "tcp://144.202.100.245:26657" --rabbitMQ "$RABBITMQ" --zone vbstreetz &
watcher --tmRPC "tcp://goz.umbrellavalidator.com:26657" --rabbitMQ "$RABBITMQ" --zone umbrella &
watcher --tmRPC "tcp://3.22.166.56:26657" --rabbitMQ "$RABBITMQ" --zone timetowinchain &
watcher --tmRPC "tcp://thetechtrap.com:26657" --rabbitMQ "$RABBITMQ" --zone thetechtrap-goz &
watcher --tmRPC "tcp:stratus.mycryptobets.com:26657" --rabbitMQ "$RABBITMQ" --zone stardust-1111 &
wait