package watcher

import (
	"time"

	"github.com/cosmos/cosmos-sdk/std"

	wasmx "github.com/InjectiveLabs/sdk-go/chain/wasmx/types"
	cosmoscodectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmoscryptoed "github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cosmoscryptomultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	cosmoscryptosecp "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cosmoscryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"
	solomachine "github.com/cosmos/ibc-go/v7/modules/light-clients/06-solomachine"
	ibcclients "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	auction "github.com/InjectiveLabs/sdk-go/chain/auction/types"
	keyscodec "github.com/InjectiveLabs/sdk-go/chain/crypto/codec"
	exchange "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	insurance "github.com/InjectiveLabs/sdk-go/chain/insurance/types"
	ocr "github.com/InjectiveLabs/sdk-go/chain/ocr/types"
	oracle "github.com/InjectiveLabs/sdk-go/chain/oracle/types"
	peggy "github.com/InjectiveLabs/sdk-go/chain/peggy/types"
	tokenfactory "github.com/InjectiveLabs/sdk-go/chain/tokenfactory/types"
	chaintypes "github.com/InjectiveLabs/sdk-go/chain/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govtypesv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	paramproposaltypes "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	icatypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v7/modules/apps/29-fee/types"
	ibcapplicationtypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibccoretypes "github.com/cosmos/ibc-go/v7/modules/core/types"
)

const (
	AccountAddressPrefix = "inj"
)

var (
	AccountPubKeyPrefix    = AccountAddressPrefix + "pub"
	ValidatorAddressPrefix = AccountAddressPrefix + "valoper"
	ValidatorPubKeyPrefix  = AccountAddressPrefix + "valoperpub"
	ConsNodeAddressPrefix  = AccountAddressPrefix + "valcons"
	ConsNodePubKeyPrefix   = AccountAddressPrefix + "valconspub"
)

type Header interface {
	GetTime() time.Time
	GetLastCommitHash() []byte
}

func RegisterInterfacesAndImpls(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	SetConfig()
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil))
	injectiveRegisterInterfaces(interfaceRegistry)
	registerTypes(interfaceRegistry)
}

func SetConfig() {
	config := cosmostypes.GetConfig()
	config.SetBech32PrefixForAccount(AccountAddressPrefix, AccountPubKeyPrefix)
	config.SetBech32PrefixForValidator(ValidatorAddressPrefix, ValidatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(ConsNodeAddressPrefix, ConsNodePubKeyPrefix)
	config.Seal()
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
	peggy.RegisterInterfaces(interfaceRegistry)
	ocr.RegisterInterfaces(interfaceRegistry)
	chaintypes.RegisterInterfaces(interfaceRegistry)
	wasmx.RegisterInterfaces(interfaceRegistry)
	tokenfactory.RegisterInterfaces(interfaceRegistry)

	// more cosmos types
	authtypes.RegisterInterfaces(interfaceRegistry)
	authztypes.RegisterInterfaces(interfaceRegistry)
	vestingtypes.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)
	crisistypes.RegisterInterfaces(interfaceRegistry)
	distributiontypes.RegisterInterfaces(interfaceRegistry)
	evidencetypes.RegisterInterfaces(interfaceRegistry)
	paramproposaltypes.RegisterInterfaces(interfaceRegistry)
	ibcapplicationtypes.RegisterInterfaces(interfaceRegistry)
	ibccoretypes.RegisterInterfaces(interfaceRegistry)
	slashingtypes.RegisterInterfaces(interfaceRegistry)
	stakingtypes.RegisterInterfaces(interfaceRegistry)
	upgradetypes.RegisterInterfaces(interfaceRegistry)
	feegranttypes.RegisterInterfaces(interfaceRegistry)
	govtypesv1beta1.RegisterInterfaces(interfaceRegistry)
	govtypesv1.RegisterInterfaces(interfaceRegistry)
	wasmtypes.RegisterInterfaces(interfaceRegistry)
	icatypes.RegisterInterfaces(interfaceRegistry)
	ibcfeetypes.RegisterInterfaces(interfaceRegistry)
}

func registerTypes(interfaceRegistry cosmoscodectypes.InterfaceRegistry) { // todo: need to nest. Maybe we can remove it. Old code
	interfaceRegistry.RegisterInterface("cosmos.crypto.PubKey", (*cosmoscryptotypes.PubKey)(nil))
	interfaceRegistry.RegisterImplementations((*cosmoscryptotypes.PubKey)(nil), &cosmoscryptoed.PubKey{})
	interfaceRegistry.RegisterImplementations((*cosmoscryptotypes.PubKey)(nil), &cosmoscryptosecp.PubKey{})
	interfaceRegistry.RegisterImplementations((*cosmoscryptotypes.PubKey)(nil), &cosmoscryptomultisig.LegacyAminoPubKey{})

	interfaceRegistry.RegisterImplementations((*ibcexported.ClientState)(nil), &ibcclients.ClientState{})
	interfaceRegistry.RegisterImplementations((*ibcexported.ConsensusState)(nil), &ibcclients.ConsensusState{})
	interfaceRegistry.RegisterImplementations(
		(*ibcexported.ClientMessage)(nil),
		&ibcclients.Header{},
		&solomachine.Header{},
		&solomachine.Misbehaviour{},
	)

}
