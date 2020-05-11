package watcher

import (
	"github.com/mapofzones/cosmos-watcher/x/tendermint-rabbit/block"
)

// alias for convenient use
type TxStatus = block.TxStatus

// WType is used to choose Watcher implementation
type WType = uint32

const (
	// TmRabbit is tendermint and rabbitMQ watcher interface implementation
	TmRabbit WType = iota
)
