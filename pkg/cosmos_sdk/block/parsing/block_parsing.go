package cosmos

import (
	"bytes"
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk/codec"
	types3 "github.com/tendermint/tendermint/abci/types"
	"log"

	types2 "github.com/cosmos/cosmos-sdk/types"
	types "github.com/mapofzones/cosmos-watcher/pkg/cosmos_sdk/block/types"
	watcher "github.com/mapofzones/cosmos-watcher/pkg/types"
)

func txToMessage(tx types2.Tx, hash string, errCode uint32, txResult *types3.ResponseDeliverTx) (watcher.Message, error) {
	Tx := watcher.Transaction{
		Hash:     hash,
		Accepted: errCode == 0,
	}
	for _, msg := range tx.GetMsgs() {
		msgs, err := parseMsg(msg, txResult, errCode)
		if err != nil {
			return Tx, err
		}
		Tx.Messages = append(Tx.Messages, msgs...)
	}
	return Tx, nil
}

func txErrCode(b types.Block, hash []byte) uint32 {
	for _, res := range b.Results {
		if bytes.Equal(res.Hash, hash) {
			return res.ResultCode
		}
	}

	panic("could not find tx status for given tx hash")
}

func DecodeBlock(cdc *codec.ProtoCodec, b types.Block) (types.ProcessedBlock, error) {
	block := types.ProcessedBlock{
		Height_:          b.Height,
		ChainID_:         b.ChainID,
		BeginBlockEvents: nil,
		EndBlockEvents:   nil,
		T:                b.T,
	}

	log.Println("height:", b.Height, " txs:", len(b.Txs))
	block.Txs = make([]watcher.Message, 0, len(b.Txs))
	txResults := b.BlockResults.TxsResults
	for i, tx := range b.Txs {
		decoded, err := decodeTx(cdc, tx)
		if err != nil {
			return block, err
		}

		txMessage, err := txToMessage(decoded, hex.EncodeToString(tx.Hash()), txErrCode(b, tx.Hash()), txResults[i])
		if err != nil {
			return block, err
		}
		block.Txs = append(block.Txs, txMessage)
	}

	return block, nil
}
