module github.com/mapofzones/cosmos-watcher

go 1.17

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)

require (
	github.com/cosmos/cosmos-sdk v0.44.3
	github.com/cosmos/ibc-go/v2 v2.0.0
	github.com/gogo/protobuf v1.3.3
	github.com/osmosis-labs/osmosis v1.0.3-0.20210807021107-13916d1e10bc
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/tendermint v0.34.11
)
