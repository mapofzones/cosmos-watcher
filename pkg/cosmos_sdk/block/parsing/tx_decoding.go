package cosmos

import (
	"errors"
	"log"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/tendermint/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	auth2 "github.com/cosmos/cosmos-sdk/x/auth/tx"
)

var DecodeErr = errors.New("could not decode tx")

func decodeTx(codec *codec.ProtoCodec, tx types.Tx) (sdk.Tx, error) {
	log.Println("pkg.cosmos_sdk.block.parsing.tx_decoding.go - 1")
	txInterface, err := auth2.DefaultTxDecoder(codec)(tx)
	if err != nil {
		log.Println("pkg.cosmos_sdk.block.parsing.tx_decoding.go - 2")
		log.Println(err)
		return auth.StdTx{}, DecodeErr
	}
	log.Println("pkg.cosmos_sdk.block.parsing.tx_decoding.go - 3")
	return toStdTx(txInterface)
}

// Decode accept tx bytes and transforms them to cosmos std tx
func toStdTx(tx sdk.Tx) (sdk.Tx, error) {
	log.Println("pkg.cosmos_sdk.block.parsing.tx_decoding.go - 4")
	stdTx, ok := tx.(sdk.Tx)
	if !ok {
		log.Println("pkg.cosmos_sdk.block.parsing.tx_decoding.go - 5")
		return nil, DecodeErr
	}
	log.Println("pkg.cosmos_sdk.block.parsing.tx_decoding.go - 6")
	return stdTx, nil
}
