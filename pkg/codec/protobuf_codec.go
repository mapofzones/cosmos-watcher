package watcher

import (
	"github.com/gogo/protobuf/proto"

	cosmoscodectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmoscryptoed "github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cosmoscryptomultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	cosmoscryptosecp "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cosmoscryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	ibcexported "github.com/cosmos/ibc-go/modules/core/exported"
	ibcclients "github.com/cosmos/ibc-go/modules/light-clients/07-tendermint/types"
	irisapp "github.com/irisnet/irishub/simapp"
)

const (
	// Bech32ChainPrefix defines the prefix of this chain
	Bech32ChainPrefix = "i"

	// PrefixAcc is the prefix for account
	PrefixAcc = "a"

	// PrefixValidator is the prefix for validator keys
	PrefixValidator = "v"

	// PrefixConsensus is the prefix for consensus keys
	PrefixConsensus = "c"

	// PrefixPublic is the prefix for public
	PrefixPublic = "p"

	// PrefixAddress is the prefix for address
	PrefixAddress = "a"

	// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address
	Bech32PrefixAccAddr = Bech32ChainPrefix + PrefixAcc + PrefixAddress
	// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key
	Bech32PrefixAccPub = Bech32ChainPrefix + PrefixAcc + PrefixPublic
	// Bech32PrefixValAddr defines the Bech32 prefix of a validator's operator address
	Bech32PrefixValAddr = Bech32ChainPrefix + PrefixValidator + PrefixAddress
	// Bech32PrefixValPub defines the Bech32 prefix of a validator's operator public key
	Bech32PrefixValPub = Bech32ChainPrefix + PrefixValidator + PrefixPublic
	// Bech32PrefixConsAddr defines the Bech32 prefix of a consensus node address
	Bech32PrefixConsAddr = Bech32ChainPrefix + PrefixConsensus + PrefixAddress
	// Bech32PrefixConsPub defines the Bech32 prefix of a consensus node public key
	Bech32PrefixConsPub = Bech32ChainPrefix + PrefixConsensus + PrefixPublic
)

func RegisterInterfacesAndImpls(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	SetConfig()
	impls := getMessageImplementations()
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), impls...)
	irisRegisterInterfaces(interfaceRegistry)
	registerTypes(interfaceRegistry)
}

func SetConfig() {
	config := cosmostypes.GetConfig()
	config.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)
	config.Seal()
}

func irisRegisterInterfaces(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	irisapp.ModuleBasics.RegisterInterfaces(interfaceRegistry)
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
