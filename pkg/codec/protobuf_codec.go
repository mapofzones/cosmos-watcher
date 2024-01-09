package watcher

import (
	etherapp "github.com/evmos/ethermint/app"
	ethercodec "github.com/evmos/ethermint/crypto/codec"
	ethertypes "github.com/evmos/ethermint/types"

	cosmoscodectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmoscryptoed "github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cosmoscryptomultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	cosmoscryptosecp "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cosmoscryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"
	ibcclients "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"
	irisapp "github.com/irisnet/irishub/v2/app"
)

func RegisterInterfacesAndImpls(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	//SetConfig()
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil))
	irisRegisterInterfaces(interfaceRegistry)
	registerTypes(interfaceRegistry)
}

func irisRegisterInterfaces(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	irisapp.ModuleBasics.RegisterInterfaces(interfaceRegistry)
	ethercodec.RegisterInterfaces(interfaceRegistry)
	etherapp.ModuleBasics.RegisterInterfaces(interfaceRegistry)
	ethertypes.RegisterInterfaces(interfaceRegistry)
}

func registerTypes(interfaceRegistry cosmoscodectypes.InterfaceRegistry) { // todo: need to nest. Maybe we can remove it. Old code
	interfaceRegistry.RegisterInterface("cosmos.crypto.PubKey", (*cosmoscryptotypes.PubKey)(nil))
	interfaceRegistry.RegisterImplementations((*cosmoscryptotypes.PubKey)(nil), &cosmoscryptoed.PubKey{})
	interfaceRegistry.RegisterImplementations((*cosmoscryptotypes.PubKey)(nil), &cosmoscryptosecp.PubKey{})
	interfaceRegistry.RegisterImplementations((*cosmoscryptotypes.PubKey)(nil), &cosmoscryptomultisig.LegacyAminoPubKey{})

	interfaceRegistry.RegisterImplementations((*ibcexported.ClientState)(nil), &ibcclients.ClientState{})
	interfaceRegistry.RegisterImplementations((*ibcexported.ConsensusState)(nil), &ibcclients.ConsensusState{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &authz.MsgGrant{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &authz.MsgExec{})
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), &authz.MsgRevoke{})
	interfaceRegistry.RegisterImplementations((*authz.Authorization)(nil), &authz.GenericAuthorization{})

}
