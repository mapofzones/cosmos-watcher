package watcher

import (
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/gogo/protobuf/proto"

	wasmx "github.com/InjectiveLabs/sdk-go/chain/wasmx/types"
	cosmoscodectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmoscryptoed "github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cosmoscryptomultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	cosmoscryptosecp "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cosmoscryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	ibcexported "github.com/cosmos/ibc-go/v3/modules/core/exported"
	ibcclients "github.com/cosmos/ibc-go/v3/modules/light-clients/07-tendermint/types"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	auction "github.com/InjectiveLabs/sdk-go/chain/auction/types"
	keyscodec "github.com/InjectiveLabs/sdk-go/chain/crypto/codec"
	evm "github.com/InjectiveLabs/sdk-go/chain/evm/types"
	exchange "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	insurance "github.com/InjectiveLabs/sdk-go/chain/insurance/types"
	ocr "github.com/InjectiveLabs/sdk-go/chain/ocr/types"
	oracle "github.com/InjectiveLabs/sdk-go/chain/oracle/types"
	peggy "github.com/InjectiveLabs/sdk-go/chain/peggy/types"
	chaintypes "github.com/InjectiveLabs/sdk-go/chain/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramproposaltypes "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	ibcapplicationtypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	ibccoretypes "github.com/cosmos/ibc-go/v3/modules/core/types"
)

func RegisterInterfacesAndImpls(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	impls := getMessageImplementations()
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), impls...)
	injectiveRegisterInterfaces(interfaceRegistry)
	registerTypes(interfaceRegistry)
}

func injectiveRegisterInterfaces(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	//injectiveapp.ModuleBasics.RegisterInterfaces(interfaceRegistry)

	//source to upgrade codec: https://github.com/InjectiveLabs/sdk-go/blob/962c8f3e7ea5e72ae351b05951cf0d8eb0886115/client/chain/context.go#84
	//source to upgrade codec: https://github.com/InjectiveLabs/sdk-go/blob/9f6b7221a84f87c49c80ec6223284c2240bc4d00/chain/client/context.go#L82

	//interfaceRegistry := types.NewInterfaceRegistry()
	keyscodec.RegisterInterfaces(interfaceRegistry)
	std.RegisterInterfaces(interfaceRegistry)
	exchange.RegisterInterfaces(interfaceRegistry)
	insurance.RegisterInterfaces(interfaceRegistry)
	auction.RegisterInterfaces(interfaceRegistry)
	oracle.RegisterInterfaces(interfaceRegistry)
	evm.RegisterInterfaces(interfaceRegistry)
	peggy.RegisterInterfaces(interfaceRegistry)
	ocr.RegisterInterfaces(interfaceRegistry)
	chaintypes.RegisterInterfaces(interfaceRegistry)
	wasmx.RegisterInterfaces(interfaceRegistry)

	// more cosmos types
	authtypes.RegisterInterfaces(interfaceRegistry)
	authztypes.RegisterInterfaces(interfaceRegistry)
	vestingtypes.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)
	crisistypes.RegisterInterfaces(interfaceRegistry)
	distributiontypes.RegisterInterfaces(interfaceRegistry)
	evidencetypes.RegisterInterfaces(interfaceRegistry)
	govtypes.RegisterInterfaces(interfaceRegistry)
	paramproposaltypes.RegisterInterfaces(interfaceRegistry)
	ibcapplicationtypes.RegisterInterfaces(interfaceRegistry)
	ibccoretypes.RegisterInterfaces(interfaceRegistry)
	slashingtypes.RegisterInterfaces(interfaceRegistry)
	stakingtypes.RegisterInterfaces(interfaceRegistry)
	upgradetypes.RegisterInterfaces(interfaceRegistry)
	feegranttypes.RegisterInterfaces(interfaceRegistry)
	wasmtypes.RegisterInterfaces(interfaceRegistry)
	icatypes.RegisterInterfaces(interfaceRegistry)
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
