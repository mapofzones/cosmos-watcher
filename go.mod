module github.com/mapofzones/cosmos-watcher

go 1.16

require (
	github.com/Workiva/go-datastructures v1.0.53 // indirect
	github.com/btcsuite/btcd v0.22.1 // indirect
	github.com/cosmos/go-bip39 v1.0.0 // indirect
	github.com/danieljoos/wincred v1.1.0 // indirect
	github.com/felixge/httpsnoop v1.0.2 // indirect
	github.com/go-kit/kit v0.12.0 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/gogo/protobuf v1.3.3 // indirect
	github.com/google/btree v1.0.1 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/libp2p/go-buffer-pool v0.1.0 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/okex/exchain v1.6.3
	github.com/onsi/gomega v1.19.0 // indirect
	github.com/pelletier/go-toml/v2 v2.0.2 // indirect
	github.com/prometheus/client_golang v1.12.2 // indirect
	github.com/prometheus/common v0.34.0 // indirect
	github.com/rakyll/statik v0.1.7 // indirect
	github.com/rogpeppe/go-internal v1.8.1 // indirect
	github.com/rs/cors v1.8.2 // indirect
	github.com/spf13/cobra v1.5.0 // indirect
	github.com/spf13/viper v1.12.0 // indirect
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.8.0 // indirect
	github.com/subosito/gotenv v1.4.0 // indirect
	github.com/tendermint/go-amino v0.16.0
	github.com/tklauser/go-sysconf v0.3.10 // indirect
	github.com/tyler-smith/go-bip39 v1.0.2 // indirect
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e // indirect
	golang.org/x/net v0.0.0-20220726230323-06994584191e // indirect
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4 // indirect
	golang.org/x/sys v0.0.0-20220727055044-e65921a090b8 // indirect
	golang.org/x/term v0.0.0-20220722155259-a9ba230a4035 // indirect
	google.golang.org/genproto v0.0.0-20220725144611-272f38e5d71b // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/ini.v1 v1.66.6 // indirect
)

replace (
	github.com/buger/jsonparser => github.com/buger/jsonparser v1.0.0 // imported by nacos-go-sdk, upgraded to v1.0.0 in case of a known vulnerable bug
	github.com/ethereum/go-ethereum => github.com/okex/go-ethereum v1.10.8-okc1
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	github.com/tendermint/go-amino => github.com/okex/go-amino v0.15.1-okc4
)
