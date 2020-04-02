#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
rm $DIR/configs/*

AMQP=amqp://ggmjxdkq:HJQ4N7gABKrLDWoneYwr0M-qZZDconkO@clam.rmq.cloudamqp.com/ggmjxdkq

TEMPLATE='{"NodeAddr":"ws://REPLACEME/websocket","RabbitMQAddr":"$AMQP","BatchSize":1,"Precision":0}'
PROCFILETEMPLATE="worker: exec go run cmd/main.go --config ./configs/"

cd /tmp

# if [ -d "/tmp/relayer" ];then
#     rm -rf "/tmp/relayer"
# fi
# git clone https://github.com/cosmos/relayer.git

if [ -f $DIR/Procfile ];then
rm $DIR/Procfile
touch $DIR/Procfile
fi

for filename in /tmp/relayer/testnets/relayer-alpha/*.json; do
    echo $TEMPLATE > $DIR/configs/$(basename $filename)
    ADDR=$(jq <$filename '.["rpc-addr"]' -r | sed 's|.*://\(.*\)|\1|')
    sed -i 's/REPLACEME/'$ADDR'/g' $DIR/configs/$(basename $filename)
    echo $PROCFILETEMPLATE$(basename $filename) >> $DIR/Procfile
done
