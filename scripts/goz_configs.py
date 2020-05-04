import json
import subprocess
import os
from pathlib import Path
from urllib.parse import urlsplit
from os import listdir
from os.path import isfile, join
files = [f for f in listdir(os.path.dirname(os.path.realpath(__file__)) + "/../configs")
         if isfile(join(os.path.dirname(os.path.realpath(__file__)) + "/../configs", f))]

# here we store list of our chains
chains = {}

for file in files:
    data = json.load(open(os.path.dirname(
        os.path.realpath(__file__)) + "/../configs/" + file, "r"))
    for key, value in data.items():
        chains[file[:-5]] = urlsplit(value).hostname


subprocess.run(
    ["rm", "-rf", "/tmp/goz"]
)

subprocess.run(
    ["git", "clone", "https://github.com/cosmosdevs/GameOfZones.git", "/tmp/goz"])

ibc_configs = [f for f in listdir(
    "/tmp/goz/contestant_info") if isfile(join("/tmp/goz/contestant_info", f))]

# add goz configs to a separate file
# watcher --tmRPC "ws://ibc.staking.fund:26657/websocket" --rabbitMQ "$RABBITMQ" --zone stakingfund &
f = open(os.path.dirname(os.path.realpath(__file__)) + "/goz.sh", "w+")
os.chmod(f.name, 0o755)
f.write("#!/bin/bash\n")

for file in ibc_configs:
    try:
        data = json.load(open("/tmp/goz/contestant_info/"+file))
    except:
        print("could not read: ", file)
        continue
    try:
        if 'node_addr' in data:
            f.write("watcher --tmRPC \"" + data["node_addr"] + "\"" +
                    " --rabbitMQ \"$RABBITMQ\" --zone " + data["chain_id"] + " &\n")
        if 'node addr' in data:
            f.write("watcher --tmRPC \"" + data["node addr"] + "\"" +
                    " --rabbitMQ \"$RABBITMQ\" --zone " + data["chain_id"] + " &\n")
    except:
        print("could not read: ", file)
f.write('watcher --tmRPC "ws://35.233.155.199:26657/websocket" --rabbitMQ "$RABBITMQ" --zone gameofzoneshub-1a &\n')
f.write("wait")
