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

var ErrDecode = errors.New("could not decode tx")

func decodeTx(codec *codec.ProtoCodec, tx types.Tx) (sdk.Tx, error) {
	txInterface, err := auth2.DefaultTxDecoder(codec)(tx)
	if err != nil {
		log.Println(err)
		return nil, ErrDecode
	}
	return txInterface, nil
}

// Decode accept tx bytes and transforms them to cosmos std tx
func toStdTx(tx sdk.Tx) (sdk.Tx, error) {
	return tx, nil
}

// Decode accept tx bytes and transforms them to cosmos sign tx
func toSignTx(tx sdk.Tx) (sign.Tx, error) {
	stdTx, ok := tx.(sign.Tx)

	if !ok {
		return nil, ErrDecode
	}
	return stdTx, nil
}
