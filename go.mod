module github.com/mapofzones/cosmos-watcher

go 1.20

replace (
	// use notional's wasmd fork with token factory
	github.com/CosmWasm/wasmd => github.com/notional-labs/wasmd v0.40.0-tf.rc2
	// Fix upstream GHSA-h395-qcrw-5vmq vulnerability.
	// TODO Remove it: https://github.com/cosmos/cosmos-sdk/issues/10409
	github.com/gin-gonic/gin => github.com/gin-gonic/gin v1.8.1

	// use notional's packet-forward-middleware fork with version 7 using cometbft
	github.com/strangelove-ventures/packet-forward-middleware/v7 v7.0.0 => github.com/notional-labs/packet-forward-middleware/v7 v7.0.0

	// use a patched alliance
	github.com/terra-money/alliance => github.com/faddat/alliance v0.0.1-beta3.0.20230407115440-caf3f6aa780a
)

require (
	github.com/White-Whale-Defi-Platform/migaloo-chain/v3 v3.0.0-20230412014946-3a3a25d9cf21
	github.com/cosmos/cosmos-sdk v0.47.1
	github.com/cosmos/ibc-go/v7 v7.0.0
	github.com/gogo/protobuf v1.3.2
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.8.2
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/tendermint v0.35.9
)
