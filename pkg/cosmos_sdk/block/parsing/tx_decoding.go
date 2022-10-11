package cosmos

import (
	"errors"
	"log"

	"github.com/okex/exchain/libs/tendermint/types"

	"github.com/okex/exchain/app"
	oeccodec "github.com/okex/exchain/app/codec"
	//auth "github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	//auth2 "github.com/okex/exchain/libs/cosmos-sdk/x/auth"
	auth "github.com/okex/exchain/libs/cosmos-sdk/x/auth/types"
	evmtypes "github.com/okex/exchain/x/evm/types"
)

var DecodeErr = errors.New("could not decode tx")
var cdc = oeccodec.MakeCodec(app.ModuleBasics)

func decodeTx(tx types.Tx) (sdk.Tx, error) {
	txInterface, err := evmtypes.TxDecoder(cdc)(tx, evmtypes.IGNORE_HEIGHT_CHECKING)
	if err != nil {
		log.Println(err)
		return &auth.StdTx{}, DecodeErr
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
