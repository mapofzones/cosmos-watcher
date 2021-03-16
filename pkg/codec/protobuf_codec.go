package watcher

import (
	"github.com/gogo/protobuf/proto"

	cosmoscodectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmossimapp "github.com/cosmos/cosmos-sdk/simapp"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"

	//cyberapp "github.com/cybercongress/go-cyber/app"
	cybercrontypes "github.com/cybercongress/go-cyber/x/cron/types"
	cyberenergytypes "github.com/cybercongress/go-cyber/x/energy/types"
	cybergraphtypes "github.com/cybercongress/go-cyber/x/graph/types"
	cyberresourcestypes "github.com/cybercongress/go-cyber/x/resources/types"

	irisguardiantypes "github.com/irisnet/irishub/modules/guardian/types"
	irisnfttypes "github.com/irisnet/irismod/modules/nft/types"

	akashapp "github.com/ovrclk/akash/app"
)

func RegisterMessagesImplementations(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	impls := getMessageImplementations()
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), impls...)
	irisguardiantypes.RegisterInterfaces(interfaceRegistry)
	irisnfttypes.RegisterInterfaces(interfaceRegistry)
	registerCosmosInterfaces(interfaceRegistry)
	registerCyberInterfaces(interfaceRegistry)
	registerAkashInterfaces(interfaceRegistry)
}

func registerCosmosInterfaces(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	cosmossimapp.ModuleBasics.RegisterInterfaces(interfaceRegistry)
}

func registerCyberInterfaces(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	//cyberapp.ModuleBasics.RegisterInterfaces(interfaceRegistry) // todo: need to fix docker build wasm error!
	cyberresourcestypes.RegisterInterfaces(interfaceRegistry)
	cybergraphtypes.RegisterInterfaces(interfaceRegistry)
	cybercrontypes.RegisterInterfaces(interfaceRegistry)
	cyberenergytypes.RegisterInterfaces(interfaceRegistry)
}

func registerAkashInterfaces(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	akashapp.ModuleBasics().RegisterInterfaces(interfaceRegistry)
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
