package watcher

import (
	"github.com/gogo/protobuf/proto"
	jackalapp "github.com/jackalLabs/canine-chain/v3/app"

	cosmoscodectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmoscryptoed "github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cosmoscryptomultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	cosmoscryptosecp "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cosmoscryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	ibcexported "github.com/cosmos/ibc-go/v4/modules/core/exported"
	ibcclients "github.com/cosmos/ibc-go/v4/modules/light-clients/07-tendermint/types"
	jackaltypes "github.com/jackalLabs/canine-chain/v3/x/storage/types"
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

	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgPostproof{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgSignContract{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgSetProviderIP{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgSetProviderTotalspace{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgInitProvider{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgCancelContract{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgPostContract{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgBuyStorage{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &jackaltypes.MsgClaimStray{})
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
