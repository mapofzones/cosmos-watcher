import json
import subprocess
import os
from pathlib import Path
from urllib.parse import urlsplit
from os import listdir
from os.path import isfile, join
files = [f for f in listdir(os.path.dirname(os.path.realpath(__file__)) + "/configs")
         if isfile(join(os.path.dirname(os.path.realpath(__file__)) + "/configs", f))]

# here we store list of our chains
chains = {}

for file in files:
    data = json.load(open(os.path.dirname(
        os.path.realpath(__file__)) + "/configs/" + file, "r"))
    for key, value in data.items():
        chains[file[:-5]] = urlsplit(value).hostname


subprocess.run(
    ["rm", "-rf", "/tmp/ibc-viz-server"]
)

subprocess.run(
    ["git", "clone", "https://github.com/allinbits/ibc-viz-server.git", "/tmp/ibc-viz-server"])

ibc_config = open("/tmp/ibc-viz-server/src/config.json")

data = json.load(ibc_config)

Procfile = open("Procfile", "a+")

for blockchain in data["blockchains"]:
    if blockchain not in chains.values():
        open(os.path.dirname(os.path.realpath(__file__)) + "/configs/" + blockchain + ".json",
             "w+").write(json.dumps({"NodeAddr": "ws://"+blockchain+":26657/websocket"}))
