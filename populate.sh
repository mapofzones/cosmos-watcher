#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
rm $DIR/configs/*

TEMPLATE='{"NodeAddr":"ws://REPLACEME/websocket"}'
PROCFILETEMPLATE=": bin/watcher --config ./configs/"

cd /tmp

if [ -d "/tmp/relayer" ];then
    rm -rf "/tmp/relayer"
fi
git clone https://github.com/iqlusioninc/relayer.git

if [ -f $DIR/Procfile ];then
rm $DIR/Procfile
touch $DIR/Procfile
fi

for filename in /tmp/relayer/testnets/relayer-alpha-2/*.json; do
    echo $TEMPLATE > $DIR/configs/$(basename $filename)
    ADDR=$(jq <$filename '.["rpc-addr"]' -r | sed 's|.*://\(.*\)|\1|')
    sed -i 's/REPLACEME/'$ADDR'/g' $DIR/configs/$(basename $filename)
    echo "$(basename $filename .json)$PROCFILETEMPLATE$(basename $filename)" >> $DIR/Procfile
done

echo "all: ./run_all.sh" >> $DIR/Procfile