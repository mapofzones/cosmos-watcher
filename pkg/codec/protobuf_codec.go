package watcher

import (
	"github.com/gogo/protobuf/proto"

	cosmoscodectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmoscryptoed "github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cosmoscryptomultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	cosmoscryptosecp "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cosmoscryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	ibcexported "github.com/cosmos/ibc-go/v4/modules/core/exported"
	ibcclients "github.com/cosmos/ibc-go/v4/modules/light-clients/07-tendermint/types"

	akashMarket "github.com/akash-network/akash-api/go/node/market/v1beta4"
	akashapp "github.com/akash-network/node/app"
)

const (
	AccountAddressPrefix = "akash"
)

var (
	AccountPubKeyPrefix    = AccountAddressPrefix + "pub"
	ValidatorAddressPrefix = AccountAddressPrefix + "valoper"
	ValidatorPubKeyPrefix  = AccountAddressPrefix + "valoperpub"
	ConsNodeAddressPrefix  = AccountAddressPrefix + "valcons"
	ConsNodePubKeyPrefix   = AccountAddressPrefix + "valconspub"
)

func RegisterInterfacesAndImpls(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	//addressConfig()
	impls := getMessageImplementations()
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), impls...)
	akashRegisterInterfaces(interfaceRegistry)
	registerTypes(interfaceRegistry)
}

func addressConfig() {
	config := cosmostypes.GetConfig()
	config.SetBech32PrefixForAccount(AccountAddressPrefix, AccountPubKeyPrefix)
	config.SetBech32PrefixForValidator(ValidatorAddressPrefix, ValidatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(ConsNodeAddressPrefix, ConsNodePubKeyPrefix)
	config.Seal()
}

func akashRegisterInterfaces(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	akashapp.ModuleBasics().RegisterInterfaces(interfaceRegistry)
	akashMarket.RegisterInterfaces(interfaceRegistry)
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
