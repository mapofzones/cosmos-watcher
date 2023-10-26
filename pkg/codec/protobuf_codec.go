package watcher

import (
	cosmoscodectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"

	ibcclients "github.com/cosmos/ibc-go/v6/modules/light-clients/07-tendermint/types"
	theapp "github.com/haqq-network/haqq/app"
	cryptocodecs "github.com/haqq-network/haqq/crypto/codec"
	chaintypes "github.com/haqq-network/haqq/types"
)

func RegisterInterfacesAndImpls(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil))
	theapp.ModuleBasics.RegisterInterfaces(interfaceRegistry)
	chaintypes.RegisterInterfaces(interfaceRegistry)
	ibcclients.RegisterInterfaces(interfaceRegistry)
	cryptocodecs.RegisterInterfaces(interfaceRegistry)
}
