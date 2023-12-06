package cosmos

import (
	"errors"
	"log"

	sign "github.com/cosmos/cosmos-sdk/x/auth/signing"

	"github.com/cometbft/cometbft/types"
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	auth2 "github.com/cosmos/cosmos-sdk/x/auth/tx"
)

var DecodeErr = errors.New("could not decode tx")

func decodeTx(codec *codec.ProtoCodec, tx types.Tx) (sdk.Tx, error) {
	txInterface, err := auth2.DefaultTxDecoder(codec)(tx)
	if err != nil {
		log.Println(err)
		return nil, DecodeErr
	}
	return txInterface, nil
}

// Decode accept tx bytes and transforms them to cosmos std tx
func toStdTx(tx sdk.Tx) (sdk.Tx, error) {
	stdTx, ok := tx.(sdk.Tx)

	//log.Println(stdTx)
	if !ok {
		return nil, DecodeErr
	}
	return stdTx, nil
}

// Decode accept tx bytes and transforms them to cosmos sign tx
func toSignTx(tx sdk.Tx) (sign.Tx, error) {
	stdTx, ok := tx.(sign.Tx)

	//log.Println(stdTx)
	if !ok {
		return nil, DecodeErr
	}
	return stdTx, nil
}
