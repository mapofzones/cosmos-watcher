package watcher

import (
	watcher "github.com/mapofzones/cosmos-watcher/pkg/types"
	"github.com/tendermint/go-amino"
)

func RegisterTypes(codec *amino.Codec) {
	codec.RegisterInterface((*watcher.StateTransition)(nil), nil)
	codec.RegisterInterface((*watcher.Block)(nil), nil)
	codec.RegisterConcrete(&watcher.Transaction{}, "watcher/transaction", nil)
	codec.RegisterConcrete(&watcher.Transfer{}, "watcher/transfer", nil)
	codec.RegisterConcrete(&watcher.CreateClient{}, "watcher/create_client", nil)
	codec.RegisterConcrete(&watcher.CreateConnection{}, "watcher/create_connection", nil)
	codec.RegisterConcrete(&watcher.CreateChannel{}, "watcher/create_channel", nil)
	codec.RegisterConcrete(&watcher.OpenChannel{}, "watcher/open_channel", nil)
	codec.RegisterConcrete(&watcher.CloseChannel{}, "watcher/close_channel", nil)
	codec.RegisterConcrete(&watcher.IBCTransfer{}, "watcher/ibc_transfer", nil)
}
