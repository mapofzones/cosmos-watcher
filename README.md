# cosmos-watcher

Status of Last Deployment:<br>
<img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=juno"><br>

# General
The cosmos-watcher is a standalone process that takes

2 env var running via docker:
* chain_id - blockchain network_id,
* graphql - connection string to graphql hasura api,
* rabbitmq - connection string to rabbitMQ message broker,
* queue - queue name for rabbitMQ message broker,

5 env var running directly:
* rpc - a zone RPC address,
* height - a starting block number,
* rabbitmq - connection string to rabbitMQ message broker,
* queue - queue name for rabbitMQ message broker,
* chain_id - blockchain network_id to validate rpc address,

and listens to the given zone starting from the given block number.

## Usage

Running in a container:
* `docker build -t cosmos-watcher:v1 .`
* `docker run --env chain_id=<network like cosmoshub-4> --env graphql=<graphql endpoint like https://ip:port/v1/graphql> --env rabbitmq=amqp://<login>:<pass>@<ip>:<default_port=5672> --env queue=<message_broker_queue_name> -it --network="host" cosmos-watcher:v1`

# Responsibilies
The watcher listens to the new blocks, parses them, and assembly the information into the zone-neutral data structures.
```
block {
   chain_id:    <string>, the zone chain id
   block_time:  <timestamp> 
   block_num:   <number>
   txs:         array [transaction]
}

transaction {
   hash:        <string>
   sender:      <string>
   accepted:    <boolean>
   msgs:        array [message]
}

message {
        type: (transfer | ibc_transfer | create_channel | create_connection | create_client | open_channel | close_channel)
        <data related to message type>
}
```

the newly created object of the ```block``` type is sent to the queue.

# Supported blockchains

| Repository Branch | Supported zone                            | Workflow status                                                                                                             |
| ---:              |                    :---:                  |                                                                                                                        :--- |
| master, cosmoshub | `cosmoshub-4 (cosmoshub)`                 | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=cosmoshub">      |
| irishub           | `irishub-1 (irishub)`                     | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=irishub">        |
| akash             | `akashnet-2 (akash.network)`              | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=akash">          |
| sentinelhub       | `sentinelhub-2 (sentinel)`                | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=sentinelhub">    |
| persistence       | `core-1 (persistence)`                    | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=persistence">    |
| regen             | `regen-1 (regen-network)`                 | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=regen">          |
| osmosis           | `osmosis-1 (osmosis)`                     | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=osmosis">        |
| crypto-org        | `crypto-org-chain-mainnet-1 (crypto.org)` | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=crypto-org">     |
| starname          | `iov-mainnet-ibc (starname)`              | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=starname">       |
| sifchain          | `sifchain-1 (sifchain)`                   | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=sifchain">       |
| konstellation     | `darchub (konstellation)`                 | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=konstellation">  |
| stargaze          | `stargaze-1 (stargaze)`                   | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=stargaze">       |
| cronos            | `cronosmainnet_25-1 (cronos)`             | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=cronos">         |
| injective         | `injective-1 (injective)`                 | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=injective">      |
| dig               | `dig-1 (dig)`                             | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=dig">            |
| bitsong           | `bitsong-2b (bitsong)`                    | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=bitsong">        |
| juno              | `juno-1 (juno)`                           | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=juno">           |
| chihuahua         | `chihuahua-1 (chihuahua)`                 | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=chihuahua">      |
| alteredcarbon     | `alteredcarbon (alteredcarbon)`           | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=alteredcarbon">  |
| cheqd             | `cheqd-mainnet-1 (cheqd)`                 | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=cheqd">          |
| bitcanna          | `bitcanna-1 (bitcanna)`                   | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=bitcanna">       |
| desmos            | `desmos-mainnet (desmos)`                 | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=desmos">         |
| likecoin          | `likecoin-mainnet-2 (likecoin)`           | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=likecoin">       |
| fetchai           | `fetchhub-3 (fetchhub)`                   | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=fetchai">        |
| gravity-bridge    | `gravity-bridge-3 (gravity-bridge)`       | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=gravity-bridge"> |
| sommelier         | `sommelier-3 (sommelier)`                 | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=sommelier">      |
| impacthub         | `impacthub-3 (impacthub)`                 | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=impacthub">      |
| vidulum           | `vidulum-1 (vidulum)`                     | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=vidulum">        |
| comdex            | `comdex-1 (comdex)`                       | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=comdex">         |
| panacea           | `panacea-3 (panacea)`                     | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=panacea">        |
| secret            | `secret-4 (secret)`                       | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=secret">         |
| kichain           | `kichain-2 (ki)`                          | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=kichain">        |
| certik            | `shentu-2.2 (shentu)`                     | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=certik">         |
| terra             | `columbus-5 (terra)`                      | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=terra">          |
| band              | `laozi-mainnet (band)`                    | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=band">           |
| kava              | `kava-9 (kava)`                           | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=kava">           |
| cyber             | `bostrom (bostrom)`                       | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=cyber">          |
| microtick         | `microtick-1 (microtick)`                 | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=microtick">      |
| axelar            | `axelar-dojo-1 (axelar)`                  | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=axelar">         |
| emoney            | `emoney-3 (emoney)`                       | <img src="https://github.com/mapofzones/cosmos-watcher/actions/workflows/docker-image.yml/badge.svg?branch=emoney">         |
