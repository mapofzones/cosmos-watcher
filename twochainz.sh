#!/bin/bash
watcher --tmRPC "http://localhost:26657" --rabbitMQ "$RABBIT" &
watcher --tmRPC "http://localhost:26557" --rabbitMQ "$RABBIT" &
wait