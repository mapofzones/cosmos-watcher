# cosmos-watcher

Status of Last Deployment:<br>
<img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=cosmoshub"><br>

# General
The cosmos-watcher is a standalone process that takes 2 input arguments: 
* a zone RPC address, 
* a starting block number, 

and listens to the given zone starting from the given block number.

| Repository Branch | Supported zone                            | Workflow status |
| ---:              |                    :---:                  |                                       :--- |
| master, cosmoshub | `cosmoshub-4 (cosmoshub)`                 | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=cosmoshub">   |
| irishub           | `irishub-1 (irishub)`                     | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=irishub">     |
| akash             | `akashnet-2 (akash.network)`              | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=akash">       |
| sentinelhub       | `sentinelhub-2 (sentinel)`                | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=sentinelhub"> |
| persistence       | `core-1 (persistence)`                    | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=persistence"> |
| regen             | `regen-1 (regen-network)`                 | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=regen">       |
| osmosis           | `osmosis-1 (osmosis)`                     | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=osmosis">     |
| crypto-org        | `crypto-org-chain-mainnet-1 (crypto.org)` | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=crypto-org">  |
| starname          | `iov-mainnet-ibc (starname)`              | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=starname">    |
| sifchain          | `sifchain-1 (sifchain)`                   | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=sifchain">    |
| microtick         | `microtick-1 (microtick)`                 | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=microtick">   |
| emoney            | `emoney-3 (emoney)`                       | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=emoney">      |
| wasm              | `bostromdev-1 (cyber)`                    |  |
| wasm              | `musslenet-4 (wasm)`                      |  |

## Usage

Running in a container:
* `docker build -t cosmos-watcher:v1 .`
* `docker run --env chain_id=<network like cosmoshub-4> --env graphql=<graphql endpoint like https://ip:port/v1/graphql> --env rabbitmq=amqp://<login>:<pass>@<ip>:<default_port=5672> -it --network="host" cosmos-watcher:v1`

# Responsibilies
The watcher listens to the new blocks, parses them, and assembly the information into the zone-neutral data structures.
```
block {
   chain_id: <string>, the zone chain id
   block_time: <timestamp> 
   block_num: <number>
   txs: array [transaction]
}

transaction {
   hash: <string>
   msgs: array [message]
}

message {
   transfer_info: {
     sender: <address>
     recipient: <address>
     quantity: <int>
     precision: <smallint>
     token: <code>
   }
   type: (send | receive | open_channel | open_connection | open_client | unknown)
   ibc: true | false
   ibc_channel_id: <string>
   ibc_connection_id: <string>
   ibc_client_id: <string>
}
```

the newly created object of the ```block``` type is sent to the queue.

