module github.com/mapofzones/cosmos-watcher

go 1.14

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/cosmos/cosmos-sdk => github.com/mapofzones/cosmos-sdk v0.42.6-unmarshal-fix

require (
	github.com/cosmos/cosmos-sdk v0.42.5
	github.com/gogo/protobuf v1.3.3
	github.com/sentinel-official/hub v0.7.0
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/tendermint v0.34.11
)
