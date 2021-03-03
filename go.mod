module github.com/mapofzones/cosmos-watcher

go 1.14

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
replace github.com/cosmos/cosmos-sdk v0.40.0-rc3 => github.com/mapofzones/cosmos-sdk v0.40.0-rc3-fix

require (
	github.com/cosmos/cosmos-sdk v0.41.3
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/tendermint v0.34.7
)
