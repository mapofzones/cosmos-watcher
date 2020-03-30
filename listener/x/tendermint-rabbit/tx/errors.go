package tx

import "errors"

var (
	ErrInvalidTx = errors.New("could not unmarshal data into tendermint tx")
)
