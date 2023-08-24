package watcher

import (
	"github.com/gogo/protobuf/proto"

	cosmoscodectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmoscryptoed "github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cosmoscryptomultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	cosmoscryptosecp "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cosmoscryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	ibcexported "github.com/cosmos/ibc-go/v6/modules/core/exported"
	ibcclients "github.com/cosmos/ibc-go/v6/modules/light-clients/07-tendermint/types"
	functionxapp "github.com/functionx/fx-core/v5/app"

	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	ethcryptocodec "github.com/evmos/ethermint/crypto/codec"
	ethermint "github.com/evmos/ethermint/types"

	crosschaintypes "github.com/functionx/fx-core/v5/x/crosschain/types"
	gravitytypes "github.com/functionx/fx-core/v5/x/gravity/types"
)

func RegisterInterfacesAndImpls(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	impls := getMessageImplementations()
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), impls...)
	functionxRegisterInterfaces(interfaceRegistry)
	registerTypes(interfaceRegistry)
}

func functionxRegisterInterfaces(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	functionxapp.ModuleBasics.RegisterInterfaces(interfaceRegistry)

	//source to upgrade codec: https://github.com/FunctionX/fx-core/blob/v4.2.1/app/encoding.go#L32-L56
	//note, need to upgrade version
	ethermint.RegisterInterfaces(interfaceRegistry)
	authtypes.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
	ethcryptocodec.RegisterInterfaces(interfaceRegistry)
	crosschaintypes.RegisterInterfaces(interfaceRegistry)
	gravitytypes.RegisterInterfaces(interfaceRegistry)
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
