package watcher

import (
	"github.com/gogo/protobuf/proto"
	jackalapp "github.com/jackalLabs/canine-chain/app"

	cosmoscodectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmoscryptoed "github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cosmoscryptomultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	cosmoscryptosecp "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cosmoscryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	ibcexported "github.com/cosmos/ibc-go/v3/modules/core/exported"
	ibcclients "github.com/cosmos/ibc-go/v3/modules/light-clients/07-tendermint/types"
	jackaltypes "github.com/jackalLabs/canine-chain/x/storage/types"
)

const (
	AccountAddressPrefix = "jkl"
)

var (
	AccountPubKeyPrefix    = AccountAddressPrefix + "pub"
	ValidatorAddressPrefix = AccountAddressPrefix + "valoper"
	ValidatorPubKeyPrefix  = AccountAddressPrefix + "valoperpub"
	ConsNodeAddressPrefix  = AccountAddressPrefix + "valcons"
	ConsNodePubKeyPrefix   = AccountAddressPrefix + "valconspub"
)

func RegisterInterfacesAndImpls(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	SetConfig()
	impls := getMessageImplementations()
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), impls...)
	jackalRegisterInterfaces(interfaceRegistry)
	registerTypes(interfaceRegistry)
}

func jackalRegisterInterfaces(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	jackalapp.ModuleBasics.RegisterInterfaces(interfaceRegistry)
}

func SetConfig() {
	config := cosmostypes.GetConfig()
	config.SetBech32PrefixForAccount(AccountAddressPrefix, AccountPubKeyPrefix)
	config.SetBech32PrefixForValidator(ValidatorAddressPrefix, ValidatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(ConsNodeAddressPrefix, ConsNodePubKeyPrefix)
	config.Seal()
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

	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgCreateActiveDeals{})
	//interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgCancelContractResponse{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgCreateContracts{})
	//interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgCreateContractsResponse{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgUpdateContracts{})
	//interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgUpdateContractsResponse{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgDeleteContracts{})
	//interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgDeleteContractsResponse{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgCreateProofs{})
	//interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgCreateProofsResponse{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgUpdateProofs{})
	//interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgUpdateProofsResponse{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgDeleteProofs{})
	//interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgDeleteProofsResponse{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgItem{})
	//interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgItemResponse{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgPostproof{})
	//interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgPostproofResponse{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgCreateActiveDeals{})
	//interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgCreateActiveDealsResponse{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgUpdateActiveDeals{})
	//interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgUpdateActiveDealsResponse{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgDeleteActiveDeals{})
	//interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgDeleteActiveDealsResponse{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgSignContract{})
	//interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgSignContractResponse{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgCreateProviders{})
	//interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgCreateProvidersResponse{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgUpdateProviders{})
	//interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgUpdateProvidersResponse{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgDeleteProviders{})
	//interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgDeleteProvidersResponse{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgSetProviderIP{})
	//interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgSetProviderIPResponse{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgSetProviderTotalspace{})
	//interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgSetProviderTotalspaceResponse{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgInitProvider{})
	//interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgInitProviderResponse{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgCancelContract{})
	//interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgCancelContractResponse{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgPostContract{})
	//interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgPostContractResponse{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgBuyStorage{})
	//interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgBuyStorageResponse{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgClaimStray{})
	//interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgClaimStrayResponse{})
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
