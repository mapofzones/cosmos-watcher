package watcher

import (
	cosmoscodectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	cosmosbanktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	cosmosibctypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	cosmosstakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	cybergraphtypes "github.com/cybercongress/go-cyber/x/graph/types"
	cyberrecourcestypes "github.com/cybercongress/go-cyber/x/resources/types"
	"github.com/gogo/protobuf/proto"
	akashaudittypes "github.com/ovrclk/akash/x/audit/types"
	akashcerttypes "github.com/ovrclk/akash/x/cert/types"
	akashdeploymenttypes "github.com/ovrclk/akash/x/deployment/types"
	akashmarkettypes "github.com/ovrclk/akash/x/market/types"
	akashprovidertypes "github.com/ovrclk/akash/x/provider/types"
)

func RegisterMessagesImplementations(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	impls := getMessageImplementations()
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), impls...)
}

func getMessageImplementations() []proto.Message {
	impls := []proto.Message{}

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
		&cosmosbanktypes.MsgSend{},
		&cosmosibctypes.MsgCreateClient{},
		&cosmosstakingtypes.MsgCreateValidator{},
	}
	return cosmosMessages
}

func getCyberMessages() []proto.Message {
	cyberMessages := []proto.Message{
		&cyberrecourcestypes.MsgConvert{},
		&cybergraphtypes.MsgCyberlink{},
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
	}
	return akashMessages
}
