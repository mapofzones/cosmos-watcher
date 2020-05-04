#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
rm $DIR/../configs/*

TEMPLATE='{"NodeAddr":"ws://REPLACEME/websocket"}'
PROCFILETEMPLATE=": bin/watcher --rabbitMQ "$RABBITMQ" --zone "

cd /tmp

if [ -d "/tmp/relayer" ];then
    rm -rf "/tmp/relayer"
fi
git clone https://github.com/iqlusioninc/relayer.git

if [ -f $DIR/../Procfile ];then
rm $DIR/../Procfile
touch $DIR/../Procfile
fi

for filename in /tmp/relayer/testnets/relayer-alpha-2/*.json; do
    echo $TEMPLATE > $DIR/../configs/$(basename $filename)
    ADDR=$(jq <$filename '.["rpc-addr"]' -r | sed 's|.*://\(.*\)|\1|')
    sed -i 's/REPLACEME/'$ADDR'/g' $DIR/../configs/$(basename $filename)
    echo "$(basename $filename .json)$PROCFILETEMPLATE$(basename $filename .json)" >> $DIR/../Procfile
done

printf "all: ./scripts/run_all.sh\n" >> $DIR/../Procfile

python3 $DIR/goz_configs.py

#populate run_all.sh
echo "#!/bin/bash" > $DIR/run_all.sh
for filename in $DIR/../configs/*.json; do
    name=$(basename -- "$filename" .json)
    echo "watcher --tmRPC \"$(jq <$filename .NodeAddr -r)\" --rabbitMQ" '"$RABBITMQ"'" --zone $name &" >> $DIR/run_all.sh
done

echo "watcher --tmRPC \"ws://35.233.155.199:26657/websocket\" --rabbitMQ" '"$RABBITMQ"'" --zone gameofzoneshub-1a &" >> $DIR/run_all.sh
echo "wait" >> $DIR/run_all.sh

python3 $DIR/split.py