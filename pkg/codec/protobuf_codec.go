package watcher

import (
	"github.com/gogo/protobuf/proto"

	//cosmoscodectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmoscodectypes "github.com/okex/exchain/libs/cosmos-sdk/codec/types"
	//cosmoscryptoed "github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cosmoscryptoed "github.com/okex/exchain/libs/cosmos-sdk/crypto/keys/ed25519"
	//cosmoscryptomultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	cosmoscryptomultisig "github.com/okex/exchain/libs/cosmos-sdk/crypto/keys/multisig"
	//cosmoscryptosecp "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cosmoscryptosecp "github.com/okex/exchain/libs/cosmos-sdk/crypto/keys/secp256k1"
	//cosmoscryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	cosmoscryptotypes "github.com/okex/exchain/libs/cosmos-sdk/crypto/types"
	//cosmostypes "github.com/cosmos/cosmos-sdk/types"
	cosmostypes "github.com/okex/exchain/libs/cosmos-sdk/types"
	//ibcexported "github.com/cosmos/ibc-go/v3/modules/core/exported"
	ibcexported "github.com/okex/exchain/libs/ibc-go/modules/core/exported"
	//ibcclients "github.com/cosmos/ibc-go/v3/modules/light-clients/07-tendermint/types"
	ibcclients "github.com/okex/exchain/libs/ibc-go/modules/light-clients/07-tendermint/types"

	//okexsapp "github.com/osmosis-labs/osmosis/v9/app"
	okexsapp "github.com/okex/exchain/app"
)

func RegisterInterfacesAndImpls(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	impls := getMessageImplementations()
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), impls...)
	okexRegisterInterfaces(interfaceRegistry)
	registerTypes(interfaceRegistry)
}

func okexRegisterInterfaces(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	okexsapp.ModuleBasics.RegisterInterfaces(interfaceRegistry)
}

func registerTypes(interfaceRegistry cosmoscodectypes.InterfaceRegistry) { // todo: need to nest. Maybe we can remove it. Old code
	interfaceRegistry.RegisterInterface("cosmos.crypto.PubKey", (*cosmoscryptotypes.PubKey)(nil))
	interfaceRegistry.RegisterImplementations((*cosmoscryptotypes.PubKey)(nil), &cosmoscryptoed.PubKey{})
	interfaceRegistry.RegisterImplementations((*cosmoscryptotypes.PubKey)(nil), &cosmoscryptosecp.PubKey{})
	interfaceRegistry.RegisterImplementations((*cosmoscryptotypes.PubKey)(nil), &cosmoscryptomultisig.LegacyAminoPubKey{})

	interfaceRegistry.RegisterImplementations((*ibcexported.ClientState)(nil), &ibcclients.ClientState{})
	interfaceRegistry.RegisterImplementations((*ibcexported.ConsensusState)(nil), &ibcclients.ConsensusState{})
	interfaceRegistry.RegisterImplementations((*ibcexported.Header)(nil), &ibcclients.Header{})
	interfaceRegistry.RegisterImplementations((*ibcexported.Misbehaviour)(nil), &ibcclients.Misbehaviour{})
}

func getMessageImplementations() []proto.Message {
	var impls []proto.Message
	cosmosMessages := getCosmosMessages()
	impls = append(impls, cosmosMessages...)
	return impls
}

func getCosmosMessages() []proto.Message {
	cosmosMessages := []proto.Message{
		//&cosmostypes.ServiceMsg{}, // do i need it? cosmostypes.RegisterInterfaces don't exist ServiceMsg
	}
	return cosmosMessages
}
