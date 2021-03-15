# cosmos-watcher

Status of Last Deployment:<br>
<img src="https://github.com/starway-monster/cosmos-watcher/workflows/Docker%20Image%20CI/badge.svg"><br>

# General
The cosmos-watcher is a standalone process that takes 2 input arguments: 
* a zone RPC address, 
* a starting block number, 

and listens to the given zone starting from the given block number.

## Usage

Running in a container:
* `docker build -t cosmos-watcher:v1 .`
* `docker run --env height=1 --env rpc=http://<ip>:<default_port=26657> --env rabbitmq=amqp://<login>:<pass>@<ip>:<default_port=5672> -it --network="host" cosmos-watcher:v1`

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

