module github.com/mapofzones/cosmos-watcher

go 1.15

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
replace google.golang.org/grpc => google.golang.org/grpc v1.33.2
replace github.com/cosmos/cosmos-sdk => github.com/mapofzones/cosmos-sdk v0.42.9-unmarshal-fix
replace github.com/microtick/mtzone => github.com/joe-bowman/mtzone v1.0.1-0.20210926153514-274d10ece5a7

require (
	github.com/cosmos/cosmos-sdk v0.42.9
	github.com/gogo/protobuf v1.5.2
	github.com/microtick/mtzone v1.0.1-0.20210919121619-c26026bcc3cd
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/tendermint v0.34.13
)
