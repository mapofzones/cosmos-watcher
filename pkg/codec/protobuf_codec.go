package watcher

import (
	"github.com/gogo/protobuf/proto"

	cosmoscodectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmoscryptoed "github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cosmoscryptomultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	cosmoscryptosecp "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cosmoscryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	cosmossimapp "github.com/cosmos/cosmos-sdk/simapp"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	cosmosibcexported "github.com/cosmos/cosmos-sdk/x/ibc/core/exported"
	cosmosibcclients "github.com/cosmos/cosmos-sdk/x/ibc/light-clients/07-tendermint/types"

	//cyberapp "github.com/cybercongress/go-cyber/app"
	cybercrontypes "github.com/cybercongress/go-cyber/x/cron/types"
	cyberenergytypes "github.com/cybercongress/go-cyber/x/energy/types"
	cybergraphtypes "github.com/cybercongress/go-cyber/x/graph/types"
	cyberresourcestypes "github.com/cybercongress/go-cyber/x/resources/types"

	irissimapp "github.com/irisnet/irishub/simapp"

	akashapp "github.com/ovrclk/akash/app"

	sentinelapp "github.com/sentinel-official/hub"

	persistenceapp "github.com/persistenceOne/persistenceCore/application"

	cosmosapp "github.com/cosmos/gaia/v4/app"
)

func RegisterInterfacesAndImpls(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	impls := getMessageImplementations()
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), impls...)
	cosmosRegisterInterfaces(interfaceRegistry)
	irisRegisterInterfaces(interfaceRegistry)
	cyberRegisterInterfaces(interfaceRegistry)
	akashRegisterInterfaces(interfaceRegistry)
	sentinelRegisterInterfaces(interfaceRegistry)
	persistenceRegisterInterfaces(interfaceRegistry)
	registerTypes(interfaceRegistry)
}

func cosmosRegisterInterfaces(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	cosmossimapp.ModuleBasics.RegisterInterfaces(interfaceRegistry)
	cosmosapp.ModuleBasics.RegisterInterfaces(interfaceRegistry)
}

func irisRegisterInterfaces(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	irissimapp.ModuleBasics.RegisterInterfaces(interfaceRegistry)
}

func cyberRegisterInterfaces(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	//cyberapp.ModuleBasics.RegisterInterfaces(interfaceRegistry) // todo: need to fix docker build wasm error!
	cyberresourcestypes.RegisterInterfaces(interfaceRegistry)
	cybergraphtypes.RegisterInterfaces(interfaceRegistry)
	cybercrontypes.RegisterInterfaces(interfaceRegistry)
	cyberenergytypes.RegisterInterfaces(interfaceRegistry)
}

func akashRegisterInterfaces(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	akashapp.ModuleBasics().RegisterInterfaces(interfaceRegistry)
}

func sentinelRegisterInterfaces(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	sentinelapp.ModuleBasics.RegisterInterfaces(interfaceRegistry)
}

func persistenceRegisterInterfaces(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	persistenceapp.ModuleBasics.RegisterInterfaces(interfaceRegistry)
}

func registerTypes(interfaceRegistry cosmoscodectypes.InterfaceRegistry) { // todo: need to nest. Maybe we can remove it. Old code
	interfaceRegistry.RegisterInterface("cosmos.crypto.PubKey", (*cosmoscryptotypes.PubKey)(nil))
	interfaceRegistry.RegisterImplementations((*cosmoscryptotypes.PubKey)(nil), &cosmoscryptoed.PubKey{})
	interfaceRegistry.RegisterImplementations((*cosmoscryptotypes.PubKey)(nil), &cosmoscryptosecp.PubKey{})
	interfaceRegistry.RegisterImplementations((*cosmoscryptotypes.PubKey)(nil), &cosmoscryptomultisig.LegacyAminoPubKey{})

	interfaceRegistry.RegisterImplementations((*cosmosibcexported.ClientState)(nil), &cosmosibcclients.ClientState{})
	interfaceRegistry.RegisterImplementations((*cosmosibcexported.ConsensusState)(nil), &cosmosibcclients.ConsensusState{})
	interfaceRegistry.RegisterImplementations((*cosmosibcexported.Header)(nil), &cosmosibcclients.Header{})
	interfaceRegistry.RegisterImplementations((*cosmosibcexported.Misbehaviour)(nil), &cosmosibcclients.Misbehaviour{})
}

func getMessageImplementations() []proto.Message {
	var impls []proto.Message
	cosmosMessages := getCosmosMessages()
	impls = append(impls, cosmosMessages...)
	return impls
}

func getCosmosMessages() []proto.Message {
	cosmosMessages := []proto.Message{
		&cosmostypes.ServiceMsg{}, // do i need it? cosmostypes.RegisterInterfaces don't exist ServiceMsg
	}
	return cosmosMessages
}
