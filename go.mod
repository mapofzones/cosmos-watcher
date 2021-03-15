module github.com/mapofzones/cosmos-watcher

go 1.14

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4

replace github.com/cosmos/cosmos-sdk v0.40.0-rc3 => github.com/mapofzones/cosmos-sdk v0.40.0-rc3-fix

require (
	github.com/cosmos/cosmos-sdk v0.42.0
	github.com/cybercongress/go-cyber v0.2.0-alpha1
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/spf13/cobra v1.1.3 // indirect
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/tendermint v0.34.8
	golang.org/x/net v0.0.0-20201209123823-ac852fbbde11 // indirect
	golang.org/x/sys v0.0.0-20201211090839-8ad439b19e0f // indirect
	golang.org/x/text v0.3.4 // indirect
	google.golang.org/grpc v1.36.0 // indirect
)
