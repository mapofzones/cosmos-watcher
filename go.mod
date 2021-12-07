module github.com/mapofzones/cosmos-watcher

go 1.17

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/cosmos/ibc-go => github.com/mapofzones/ibc-go v1.1.0-unmarshal-fix

require (
	github.com/cosmos/cosmos-sdk v0.44.0
	github.com/cosmos/ibc-go v1.1.0
	github.com/cosmos/cosmos-sdk v0.42.9
	github.com/crypto-org-chain/chain-main/v2 v2.1.2
	github.com/gogo/protobuf v1.3.3
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/tendermint v0.34.11
)
