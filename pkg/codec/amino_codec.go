package watcher

import (
	cosmos "github.com/mapofzones/cosmos-watcher/pkg/cosmos_sdk/block/types"
	watcher "github.com/mapofzones/cosmos-watcher/pkg/types"
	"github.com/tendermint/go-amino"
	"log"
)

func RegisterTypes(codec *amino.Codec) {
	log.Println("pkg.codec.amino_codec.go - 1")
	registerBlocks(codec)
	log.Println("pkg.codec.amino_codec.go - 2")
	registerMessages(codec)
	log.Println("pkg.codec.amino_codec.go - 3")
}

func registerBlocks(codec *amino.Codec) {
	codec.RegisterInterface((*watcher.Block)(nil), nil)
	codec.RegisterConcrete(&cosmos.ProcessedBlock{}, "watcher/cosmos_block", nil)
}

func registerMessages(codec *amino.Codec) {
	codec.RegisterInterface((*watcher.Message)(nil), nil)
	codec.RegisterConcrete(watcher.Transaction{}, "watcher/transaction", nil)
	codec.RegisterConcrete(watcher.Transfer{}, "watcher/transfer", nil)
	codec.RegisterConcrete(watcher.CreateClient{}, "watcher/create_client", nil)
	codec.RegisterConcrete(watcher.CreateConnection{}, "watcher/create_connection", nil)
	codec.RegisterConcrete(watcher.CreateChannel{}, "watcher/create_channel", nil)
	codec.RegisterConcrete(watcher.OpenChannel{}, "watcher/open_channel", nil)
	codec.RegisterConcrete(watcher.CloseChannel{}, "watcher/close_channel", nil)
	codec.RegisterConcrete(watcher.IBCTransfer{}, "watcher/ibc_transfer", nil)
}
