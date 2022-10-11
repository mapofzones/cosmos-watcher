package cosmos

import (
	"bytes"
	"encoding/hex"
	"fmt"
	types3 "github.com/okex/exchain/libs/tendermint/abci/types"
	evmtypes "github.com/okex/exchain/x/evm/types"
	"log"

	types "github.com/mapofzones/cosmos-watcher/pkg/cosmos_sdk/block/types"
	watcher "github.com/mapofzones/cosmos-watcher/pkg/types"

	_ "github.com/okex/exchain/app/codec"
	ethermint "github.com/okex/exchain/app/types"
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	_ "github.com/okex/exchain/x/evm/types"
)

func txToMessage(tx sdk.Tx, hash string, errCode uint32, txResult *types3.ResponseDeliverTx, height int64) (watcher.Message, error) {
	switch tx.GetType() {
	case sdk.EvmTxType:
		msgEthTx, ok := tx.(*evmtypes.MsgEthereumTx)
		if !ok {

		}
		chainId, err := ethermint.ParseChainID("exchain-66")
		if err != nil {
			return watcher.Transaction{}, fmt.Errorf("txToMessage ParseChainID error: %v", err)
		}
		err = msgEthTx.VerifySig(chainId, height)
		if err != nil {
			return watcher.Transaction{}, fmt.Errorf("txToMessage VerifySig error: %v", err)
		}
	case sdk.StdTxType:
	default:
		return watcher.Transaction{}, fmt.Errorf("invalid transaction type: %T", tx)
	}
	Tx := watcher.Transaction{
		Hash:     hash,
		Accepted: errCode == 0,
		Sender:   tx.GetSigners()[0].String(),
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

func DecodeBlock(b types.Block) (types.ProcessedBlock, error) {
	block := types.ProcessedBlock{
		Height_:          b.Height,
		ChainID_:         b.ChainID,
		BeginBlockEvents: nil,
		EndBlockEvents:   nil,
		T:                b.T,
	}

	log.Println("height:", b.Height, " txs:", len(b.Txs))
	block.Txs = make([]watcher.Message, 0, len(b.Txs))
	for i, tx := range b.Txs {
		decoded, err := decodeTx(tx)
		if err != nil {
			return block, err
		}
		stdTx, err := toStdTx(decoded)
		if err != nil {
			return block, err
		}

		txMessage, err := txToMessage(stdTx, hex.EncodeToString(tx.Hash(b.Height)), txErrCode(b, tx.Hash(b.Height)), b.TxsResults[i], b.Height)
		if err != nil {
			return block, err
		}
		block.Txs = append(block.Txs, txMessage)
	}

	return block, nil
}
