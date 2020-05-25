package cosmos

import (
	"errors"

	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

var DecodeErr = errors.New("could not decode tx")

func decodeTx(codec *amino.Codec, tx types.Tx) (auth.StdTx, error) {
	txInterface, err := auth.DefaultTxDecoder(codec)(tx)
	if err != nil {
		return auth.StdTx{}, DecodeErr
	}
	return toStdTx(txInterface)
}

// Decode accept tx bytes and transforms them to cosmos std tx
func toStdTx(tx sdk.Tx) (auth.StdTx, error) {
	stdTx, ok := tx.(auth.StdTx)
	if !ok {
		return auth.StdTx{}, DecodeErr
	}
	return stdTx, nil
}
