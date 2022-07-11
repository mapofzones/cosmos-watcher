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
	log.Println("pkg.cosmos_sdk.block.parsing.block_parsing.go - 1")
	Tx := watcher.Transaction{
		Hash:     hash,
		Accepted: errCode == 0,
	}
	log.Println("pkg.cosmos_sdk.block.parsing.block_parsing.go - 2")
	for _, msg := range tx.GetMsgs() {
		log.Println("pkg.cosmos_sdk.block.parsing.block_parsing.go - 3")
		msgs, err := parseMsg(msg, txResult, errCode)
		if err != nil {
			log.Println("pkg.cosmos_sdk.block.parsing.block_parsing.go - 4")
			return Tx, err
		}
		log.Println("pkg.cosmos_sdk.block.parsing.block_parsing.go - 5")
		Tx.Messages = append(Tx.Messages, msgs...)
	}
	log.Println("pkg.cosmos_sdk.block.parsing.block_parsing.go - 6")
	return Tx, nil
}

func txErrCode(b types.Block, hash []byte) uint32 {
	log.Println("pkg.cosmos_sdk.block.parsing.block_parsing.go - 7")
	for _, res := range b.Results {
		log.Println("pkg.cosmos_sdk.block.parsing.block_parsing.go - 8")
		if bytes.Equal(res.Hash, hash) {
			log.Println("pkg.cosmos_sdk.block.parsing.block_parsing.go - 9")
			return res.ResultCode
		}
	}

	panic("could not find tx status for given tx hash")
}

func DecodeBlock(cdc *codec.ProtoCodec, b types.Block) (types.ProcessedBlock, error) {
	log.Println("pkg.cosmos_sdk.block.parsing.block_parsing.go - 10")
	block := types.ProcessedBlock{
		Height_:          b.Height,
		ChainID_:         b.ChainID,
		BeginBlockEvents: nil,
		EndBlockEvents:   nil,
		T:                b.T,
	}

	log.Println("height:", b.Height, " txs:", len(b.Txs))
	block.Txs = make([]watcher.Message, 0, len(b.Txs))
	log.Println("pkg.cosmos_sdk.block.parsing.block_parsing.go - 11")
	for i, tx := range b.Txs {
		log.Println("pkg.cosmos_sdk.block.parsing.block_parsing.go - 12")
		decoded, err := decodeTx(cdc, tx)
		log.Println("pkg.cosmos_sdk.block.parsing.block_parsing.go - 13")
		if err != nil {
			log.Println("pkg.cosmos_sdk.block.parsing.block_parsing.go - 14")
			return block, err
		}

		log.Println("pkg.cosmos_sdk.block.parsing.block_parsing.go - 15")
		txMessage, err := txToMessage(decoded, hex.EncodeToString(tx.Hash()), txErrCode(b, tx.Hash()), b.TxsResults[i])
		log.Println("pkg.cosmos_sdk.block.parsing.block_parsing.go - 16")
		if err != nil {
			log.Println("pkg.cosmos_sdk.block.parsing.block_parsing.go - 17")
			return block, err
		}
		log.Println("pkg.cosmos_sdk.block.parsing.block_parsing.go - 18")
		block.Txs = append(block.Txs, txMessage)
		log.Println("pkg.cosmos_sdk.block.parsing.block_parsing.go - 19")
	}

	log.Println("pkg.cosmos_sdk.block.parsing.block_parsing.go - 20")
	return block, nil
}
