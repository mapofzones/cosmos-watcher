package watcher

import (
	"github.com/gogo/protobuf/proto"

	cosmoscodectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	cosmosauthvestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	cosmosbanktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	cosmoscrisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	cosmosdistributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	cosmosevidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	cosmosgovtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	cosmosibcapptransfertypes "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/types"
	cosmosibcclienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	cosmosibcconnectiontypes "github.com/cosmos/cosmos-sdk/x/ibc/core/03-connection/types"
	cosmosibcchanneltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	cosmosslashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	cosmosstakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	cybercrontypes "github.com/cybercongress/go-cyber/x/cron/types"
	cyberenergytypes "github.com/cybercongress/go-cyber/x/energy/types"
	cybergraphtypes "github.com/cybercongress/go-cyber/x/graph/types"
	cyberresourcestypes "github.com/cybercongress/go-cyber/x/resources/types"

	akashaudittypes "github.com/ovrclk/akash/x/audit/types"
	akashcerttypes "github.com/ovrclk/akash/x/cert/types"
	akashdeploymenttypes "github.com/ovrclk/akash/x/deployment/types"
	akashmarkettypes "github.com/ovrclk/akash/x/market/types"
	akashprovidertypes "github.com/ovrclk/akash/x/provider/types"

	irisguardiantypes "github.com/irisnet/irishub/modules/guardian/types"
	irisnfttypes "github.com/irisnet/irismod/modules/nft/types"
)

func RegisterMessagesImplementations(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	impls := getMessageImplementations()
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), impls...)
	irisguardiantypes.RegisterInterfaces(interfaceRegistry)
	irisnfttypes.RegisterInterfaces(interfaceRegistry)
	registerCosmosInterfaces(interfaceRegistry)
}

func registerCosmosInterfaces(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	cosmosibcchanneltypes.RegisterInterfaces(interfaceRegistry)
	cosmosbanktypes.RegisterInterfaces(interfaceRegistry)
	cosmosibcclienttypes.RegisterInterfaces(interfaceRegistry)
	cosmosstakingtypes.RegisterInterfaces(interfaceRegistry)
	cosmosibcconnectiontypes.RegisterInterfaces(interfaceRegistry)
	cosmosauthvestingtypes.RegisterInterfaces(interfaceRegistry)
	cosmosgovtypes.RegisterInterfaces(interfaceRegistry)
	cosmosdistributiontypes.RegisterInterfaces(interfaceRegistry)
	cosmosevidencetypes.RegisterInterfaces(interfaceRegistry)
	cosmosibcapptransfertypes.RegisterInterfaces(interfaceRegistry)
	cosmosslashingtypes.RegisterInterfaces(interfaceRegistry)
	cosmoscrisistypes.RegisterInterfaces(interfaceRegistry)
	cosmostypes.RegisterInterfaces(interfaceRegistry)
}

func getMessageImplementations() []proto.Message {
	var impls []proto.Message

	cosmosMessages := getCosmosMessages()
	cyberMessages := getCyberMessages()
	akashMessages := getAkashMessages()
	impls = append(impls, cosmosMessages...)
	impls = append(impls, cyberMessages...)
	impls = append(impls, akashMessages...)

	return impls
}

func getCosmosMessages() []proto.Message {
	cosmosMessages := []proto.Message{
		&cosmostypes.ServiceMsg{}, // do i need it? cosmostypes.RegisterInterfaces don't exist ServiceMsg
	}
	return cosmosMessages
}

func getCyberMessages() []proto.Message {
	cyberMessages := []proto.Message{
		&cyberresourcestypes.MsgConvert{},
		&cybergraphtypes.MsgCyberlink{},
		&cyberresourcestypes.MsgCreateResource{},
		&cybercrontypes.MsgCronAddJob{},
		&cybercrontypes.MsgCronChangeJobBlock{},
		&cybercrontypes.MsgCronChangeJobCID{},
		&cybercrontypes.MsgCronChangeJobCallData{},
		&cybercrontypes.MsgCronChangeJobGasPrice{},
		&cybercrontypes.MsgCronChangeJobLabel{},
		&cybercrontypes.MsgCronRemoveJob{},
		&cyberenergytypes.MsgDeleteEnergyRoute{},
		&cyberenergytypes.MsgEditEnergyRoute{},
		&cyberenergytypes.MsgEditEnergyRouteAlias{},
		&cyberresourcestypes.MsgRedeemResource{},
	}
	return cyberMessages
}

func getAkashMessages() []proto.Message {
	akashMessages := []proto.Message{
		&akashprovidertypes.MsgCreateProvider{},
		&akashprovidertypes.MsgUpdateProvider{},
		&akashcerttypes.MsgCreateCertificate{},
		&akashcerttypes.MsgRevokeCertificate{},
		&akashdeploymenttypes.MsgCreateDeployment{},
		&akashdeploymenttypes.MsgCloseDeployment{},
		&akashdeploymenttypes.MsgUpdateDeployment{},
		&akashaudittypes.MsgSignProviderAttributes{},
		&akashmarkettypes.MsgCreateBid{},
		&akashmarkettypes.MsgCloseBid{},
		&akashmarkettypes.MsgCreateLease{},
		&akashdeploymenttypes.MsgCloseGroup{},
		&akashmarkettypes.MsgCloseLease{},
		&akashprovidertypes.MsgDeleteProvider{},
		&akashaudittypes.MsgDeleteProviderAttributes{},
		&akashdeploymenttypes.MsgDepositDeployment{},
		&akashdeploymenttypes.MsgPauseGroup{},
		&akashdeploymenttypes.MsgStartGroup{},
		&akashmarkettypes.MsgWithdrawLease{},
	}
	return akashMessages
}
