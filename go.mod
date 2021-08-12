module github.com/mapofzones/cosmos-watcher

go 1.14

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/tharsis/ethermint => github.com/crypto-org-chain/ethermint v0.5.0-rc0

require (
	github.com/cosmos/cosmos-sdk v0.43.0
	github.com/gogo/protobuf v1.3.3
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/tendermint v0.34.11
	github.com/tharsis/ethermint v0.4.1
	github.com/cosmos/ibc-go v1.0.0
)
