# cosmos-watcher

Status of Last Deployment:<br>
<img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=akash"><br>

# General
The cosmos-watcher is a standalone process that takes 2 input arguments: 
* a zone RPC address, 
* a starting block number, 

and listens to the given zone starting from the given block number.

| Repository Branch | Supported zone | Blockchain version |
| ---:   |                    :---:    |                                       :--- |
| master | `cosmoshub-4 (cosmoshub)`   | `gaia v4.2.1`                              |
| master | `irishub-1 (irishub)`       | `irishub v1.1.1`                           |
| master | `akashnet-2 (akash.network)`| `akash v0.12.1`                            |
| master | `sentinel (sentinel)`       | `sentinelhub v0.6.2`                       |
| master | `core-1 (persistence)`      | `persistenceCore v0.1.3`                   |
| wasm   | `bostromdev-1(cyber)`       | `go-cyber v0.2.0-alpha1`                   |
| wasm   | `musslenet-4(wasm)`         | `wasmd v0.16.0-alpha1`                     |
| regen  | `regen-1(regen-network)`    | `regen-ledger v1.0.0`                      |

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

