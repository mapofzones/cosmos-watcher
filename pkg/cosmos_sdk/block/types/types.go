package cosmos

import (
	"bytes"
	"encoding/json"
	"log"
	"time"

	watcher "github.com/mapofzones/cosmos-watcher/pkg/types"
	abci "github.com/tendermint/tendermint/abci/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
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
	log.Println("pkg.cosmos_sdk.block.types.types.go - 1")
	data, _ := json.Marshal(b)
	log.Println("pkg.cosmos_sdk.block.types.types.go - 2")
	return data
}

// TmBlock is like tendermint block but without unnecessary mutex and unused data
type TmBlock struct {
	types.Header
	types.Txs
}

// BlockFromTmResultBlock parses tm websocket response to fit our block structure
func BlockFromTmResultBlock(b types.EventDataNewBlock) TmBlock {
	log.Println("pkg.cosmos_sdk.block.types.types.go - 3")
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
	log.Println("pkg.cosmos_sdk.block.types.types.go - 4")
	bytes, _ := json.Marshal(w)
	log.Println("pkg.cosmos_sdk.block.types.types.go - 5")
	return bytes
}

// Normalize takes block with transactions and transforms it Block structure that is being send over rabbitmq
func Normalize(w WithTxs) Block {
	log.Println("pkg.cosmos_sdk.block.types.types.go - 6")
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
	log.Println("pkg.cosmos_sdk.block.types.types.go - 7")
	// if we don't have yet received the block itself, exit here so we don't panic
	if w.B == nil {
		log.Println("pkg.cosmos_sdk.block.types.types.go - 8")
		return false
	}

	log.Println("pkg.cosmos_sdk.block.types.types.go - 9")
	// helper function telling that hash is present in our txStatus slice
	present := func(hash []byte, txs []TxStatus) bool {
		log.Println("pkg.cosmos_sdk.block.types.types.go - 10")
		for _, tx := range txs {
			log.Println("pkg.cosmos_sdk.block.types.types.go - 11")
			if bytes.Equal(tx.Hash, hash) {
				return true
			}
		}
		log.Println("pkg.cosmos_sdk.block.types.types.go - 12")
		return false
	}

	log.Println("pkg.cosmos_sdk.block.types.types.go - 13")
	for _, tx := range w.B.Txs {
		log.Println("pkg.cosmos_sdk.block.types.types.go - 14")
		if !present(tx.Hash(), w.Txs) {
			return false
		}
	}

	log.Println("pkg.cosmos_sdk.block.types.types.go - 15")
	return true
}

// TxStatusFromTmResultTx parses ResultTx to give us tx hash, height and error code
// actual tx body is stored in the block
func TxStatusFromTmResultTx(t types.EventDataTx) TxStatus {
	log.Println("pkg.cosmos_sdk.block.types.types.go - 16")
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
	log.Println("pkg.cosmos_sdk.block.types.types.go - 17")
	return b.Height_
}

func (b ProcessedBlock) ChainID() string {
	log.Println("pkg.cosmos_sdk.block.types.types.go - 18")
	return b.ChainID_
}

func (b ProcessedBlock) Time() time.Time {
	log.Println("pkg.cosmos_sdk.block.types.types.go - 19")
	return b.T
}

func (b ProcessedBlock) Messages() []watcher.Message {
	log.Println("pkg.cosmos_sdk.block.types.types.go - 20")
	res := make([]watcher.Message, 0, len(b.BeginBlockEvents)+
		len(b.Txs)+len(b.EndBlockEvents))

	log.Println("pkg.cosmos_sdk.block.types.types.go - 21")
	res = append(res, b.BeginBlockEvents...)
	log.Println("pkg.cosmos_sdk.block.types.types.go - 22")
	res = append(res, b.Txs...)
	log.Println("pkg.cosmos_sdk.block.types.types.go - 23")
	res = append(res, b.EndBlockEvents...)
	log.Println("pkg.cosmos_sdk.block.types.types.go - 24")
	return res
}
