#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

for filename in $DIR/configs/*.json; do
    watcher --tmRPC $(jq <$filename .NodeAddr -r) --rabbitMQ $RABBITMQ &
done

wait