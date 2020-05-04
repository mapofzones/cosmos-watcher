package watcher

import (
	"encoding/json"
	"time"

	"github.com/mapofzones/cosmos-watcher/x/tendermint-rabbit/block"
	"github.com/tendermint/tendermint/types"
)

// Block is a unit of data being sent over in order to be processed
type Block struct {
	ChainID   string `json:"chain_id"`
	Height    int64  `json:"height"`
	types.Txs `json:"txs"`
	Results   []block.TxStatus `json:"tx_results"`
	T         time.Time        `json:"block_time"`
}

// alias for convenient use
type TxStatus = block.TxStatus

// Normalize takes block with transactions and transforms it Block structure that is being send over rabbitmq
func Normalize(w block.WithTxs) Block {
	return Block{
		ChainID: w.B.ChainID,
		Height:  w.B.Height,
		Txs:     w.B.Txs,
		Results: w.Txs,
		T:       w.B.Time,
	}
}

// JSON returns byte slice which represents block in json from
func (b Block) JSON() []byte {
	data, _ := json.Marshal(b)
	return data
}

// WType is used to choose Watcher implementation
type WType = uint32

const (
	// TmRabbit is tendermint and rabbitMQ watcher interface implementation
	TmRabbit WType = iota
)
