package cosmos

import (
	"errors"
	"github.com/okex/exchain/app"
	evmtypes "github.com/okex/exchain/x/evm/types"

	//sign "github.com/cosmos/cosmos-sdk/x/auth/signing"
	//sign "github.com/okex/exchain/libs/cosmos-sdk/x/auth/signing"
	"log"

	//"github.com/cosmos/cosmos-sdk/codec"
	"github.com/okex/exchain/libs/cosmos-sdk/codec"
	//"github.com/tendermint/tendermint/types"
	"github.com/okex/exchain/libs/tendermint/types"

	//sdk "github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	//auth "github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	//auth "github.com/okex/exchain/libs/cosmos-sdk/x/auth/legacy/legacytx"
	auth "github.com/okex/exchain/libs/cosmos-sdk/x/auth/types"
	//auth2 "github.com/cosmos/cosmos-sdk/x/auth/tx"
	//auth2 "github.com/okex/exchain/libs/cosmos-sdk/x/auth/tx"

	okexchaincodec "github.com/okex/exchain/app/codec"
)

var DecodeErr = errors.New("could not decode tx")
var codecProxy, interfaceReg = okexchaincodec.MakeCodecSuit(app.ModuleBasics)

func decodeTx(codec *codec.ProtoCodec, tx types.Tx) (sdk.Tx, error) {
	txInterface, err := evmtypes.TxDecoder(codecProxy)(tx, evmtypes.IGNORE_HEIGHT_CHECKING)
	if err != nil {
		log.Println(err)
		return &auth.StdTx{}, DecodeErr
	}
	return txInterface, nil
}

// Decode accept tx bytes and transforms them to cosmos std tx
func toStdTx(tx sdk.Tx) (sdk.Tx, error) {
	stdTx, ok := tx.(sdk.Tx)

	if !ok {
		return nil, DecodeErr
	}
	return stdTx, nil
}
