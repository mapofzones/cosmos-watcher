module github.com/mapofzones/cosmos-watcher

go 1.14

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
replace google.golang.org/grpc => google.golang.org/grpc v1.33.2
replace github.com/cosmos/cosmos-sdk v0.42.6 => github.com/mapofzones/cosmos-sdk v0.42.6-unmarshal-fix

require (
	github.com/Sifchain/sifnode v0.0.0-20210824131358-35744c0c9631
	github.com/cosmos/cosmos-sdk v0.42.9
	github.com/gogo/protobuf v1.3.3
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/tendermint v0.34.11
)
