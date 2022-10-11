package cosmos

import (
	"bytes"
	"encoding/json"
	"time"

	watcher "github.com/mapofzones/cosmos-watcher/pkg/types"
	abci "github.com/okex/exchain/libs/tendermint/abci/types"
	ctypes "github.com/okex/exchain/libs/tendermint/rpc/core/types"
	"github.com/okex/exchain/libs/tendermint/types"
)

// Block is a unit of data being sent over in order to be processed
type Block struct {
	ChainID      string `json:"chain_id"`
	Height       int64  `json:"height"`
	types.Txs    `json:"txs"`
	Results      []TxStatus `json:"tx_results"`
	T            time.Time  `json:"block_time"`
	BlockResults *ctypes.ResultBlockResults
	TxsResults   []*abci.ResponseDeliverTx `json:"txs_results"`
}

// JSON returns byte slice which represents block in json from
func (b Block) JSON() []byte {
	data, _ := json.Marshal(b)
	return data
}

// TmBlock is like tendermint block but without unnecessary mutex and unused data
type TmBlock struct {
	types.Header
	types.Txs
}

// BlockFromTmResultBlock parses tm websocket response to fit our block structure
func BlockFromTmResultBlock(b types.EventDataNewBlock) TmBlock {
	return TmBlock{
		Header: b.Block.Header,
		Txs:    b.Block.Txs,
	}
}

// TxStatus only stores transaction result code and it's hash
type TxStatus struct {
	ResultCode uint32
	Hash       []byte
	Height     int64
}

// WithTxs is used to couple blocks with txs data together for our map
type WithTxs struct {
	B   *TmBlock
	Txs []TxStatus
}

// JSON marshall struct to json format
func (w WithTxs) JSON() []byte {
	bytes, _ := json.Marshal(w)
	return bytes
}

// Normalize takes block with transactions and transforms it Block structure that is being send over rabbitmq
func Normalize(w WithTxs) Block {
	return Block{
		ChainID: w.B.ChainID,
		Height:  w.B.Height,
		Txs:     w.B.Txs,
		Results: w.Txs,
		T:       w.B.Time,
	}
}

// Full returns true if all txs are matched inside this block
func (w WithTxs) Full() bool {
	// if we don't have yet received the block itself, exit here so we don't panic
	if w.B == nil {
		return false
	}

	// helper function telling that hash is present in our txStatus slice
	present := func(hash []byte, txs []TxStatus) bool {
		for _, tx := range txs {
			if bytes.Equal(tx.Hash, hash) {
				return true
			}
		}
		return false
	}

	for _, tx := range w.B.Txs {
		if !present(tx.Hash(w.B.Height), w.Txs) {
			return false
		}
	}

	return true
}

// TxStatusFromTmResultTx parses ResultTx to give us tx hash, height and error code
// actual tx body is stored in the block
func TxStatusFromTmResultTx(t types.EventDataTx) TxStatus {
	return TxStatus{
		ResultCode: t.Result.Code,
		Hash:       t.Tx,
		Height:     t.Height,
	}
}

type ProcessedBlock struct {
	Height_          int64
	ChainID_         string
	T                time.Time
	BeginBlockEvents []watcher.Message
	Txs              []watcher.Message
	EndBlockEvents   []watcher.Message
}

func (b ProcessedBlock) Height() int64 {
	return b.Height_
}

func (b ProcessedBlock) ChainID() string {
	return b.ChainID_
}

func (b ProcessedBlock) Time() time.Time {
	return b.T
}

func (b ProcessedBlock) Messages() []watcher.Message {
	res := make([]watcher.Message, 0, len(b.BeginBlockEvents)+
		len(b.Txs)+len(b.EndBlockEvents))

	res = append(res, b.BeginBlockEvents...)
	res = append(res, b.Txs...)
	res = append(res, b.EndBlockEvents...)
	return res
}
