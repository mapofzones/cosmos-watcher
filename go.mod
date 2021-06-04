module github.com/mapofzones/cosmos-watcher

go 1.14

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/cosmos/cosmos-sdk v0.42.5 => github.com/mapofzones/cosmos-sdk v0.42.3-unmarshal-fix

require (
	github.com/cosmos/cosmos-sdk v0.42.5
	github.com/cybercongress/go-cyber v0.2.0-alpha1
	github.com/gogo/protobuf v1.3.3
	github.com/irisnet/irishub v1.1.1
	github.com/ovrclk/akash v0.12.1
	github.com/sentinel-official/hub v0.6.2
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/tendermint v0.34.10
	google.golang.org/genproto v0.0.0-20210524171403-669157292da3 // indirect
	google.golang.org/grpc v1.38.0 // indirect
)
