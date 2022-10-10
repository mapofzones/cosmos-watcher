package cosmos

import (
	"errors"
	//sign "github.com/cosmos/cosmos-sdk/x/auth/signing"
	sign "github.com/okex/exchain/libs/cosmos-sdk/x/auth/signing"
	"log"

	//"github.com/cosmos/cosmos-sdk/codec"
	"github.com/okex/exchain/libs/cosmos-sdk/codec"
	//"github.com/tendermint/tendermint/types"
	"github.com/okex/exchain/libs/tendermint/types"

	//sdk "github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	//auth "github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	auth "github.com/okex/exchain/libs/cosmos-sdk/x/auth/legacy/legacytx"
	//auth2 "github.com/cosmos/cosmos-sdk/x/auth/tx"
	auth2 "github.com/okex/exchain/libs/cosmos-sdk/x/auth/tx"
)

var DecodeErr = errors.New("could not decode tx")

func decodeTx(codec *codec.ProtoCodec, tx types.Tx) (sdk.Tx, error) {
	txInterface, err := auth2.DefaultTxDecoder(codec)(tx)
	if err != nil {
		log.Println(err)
		return auth.StdTx{}, DecodeErr
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
