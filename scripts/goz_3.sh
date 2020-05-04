#!/bin/bash
watcher --tmRPC "tcp://ibc.j96.me:26657" --rabbitMQ "$RABBITMQ" --zone plex &
watcher --tmRPC "tcp://goz.cyphercore.io:26657" --rabbitMQ "$RABBITMQ" --zone petomhub &
watcher --tmRPC "tcp://relayer.persistence.one:26657" --rabbitMQ "$RABBITMQ" --zone persistence &
watcher --tmRPC "tcp://p2p-org-1.goz.p2p.org:26657" --rabbitMQ "$RABBITMQ" --zone p2p-org-1 &
watcher --tmRPC "tcp://goz.ozonechain.xyz:26657" --rabbitMQ "$RABBITMQ" --zone ozone &
watcher --tmRPC "tcp://goz.kysenpool.io:26657" --rabbitMQ "$RABBITMQ" --zone outpost &
watcher --tmRPC "tcp://3.112.29.150:26657" --rabbitMQ "$RABBITMQ" --zone okchain &
watcher --tmRPC "tcp://144.76.118.133:26657" --rabbitMQ "$RABBITMQ" --zone nodeasy &
watcher --tmRPC "tcp://193.30.121.61:26657" --rabbitMQ "$RABBITMQ" --zone Node123 &
watcher --tmRPC "tcp://goz.nibiru.network:26657" --rabbitMQ "$RABBITMQ" --zone nibiru-ibc &
wait