package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"

	"github.com/magewar/mage/types/crypto/ethsecp256k1"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	// Register custom key types
	cdc.RegisterConcrete(&ethsecp256k1.PubKey{}, ethsecp256k1.PubKeyName, nil)
	cdc.RegisterConcrete(&ethsecp256k1.PrivKey{}, ethsecp256k1.PrivKeyName, nil)

	cdc.RegisterConcrete(&MsgSaveProfile{}, "mage/MsgSaveProfile", nil)
	cdc.RegisterConcrete(&MsgDeleteProfile{}, "mage/MsgDeleteProfile", nil)
	cdc.RegisterConcrete(&MsgRequestDTagTransfer{}, "mage/MsgRequestDTagTransfer", nil)
	cdc.RegisterConcrete(&MsgCancelDTagTransferRequest{}, "mage/MsgCancelDTagTransferRequest", nil)
	cdc.RegisterConcrete(&MsgAcceptDTagTransferRequest{}, "mage/MsgAcceptDTagTransferRequest", nil)
	cdc.RegisterConcrete(&MsgRefuseDTagTransferRequest{}, "mage/MsgRefuseDTagTransferRequest", nil)
	cdc.RegisterConcrete(&MsgLinkChainAccount{}, "mage/MsgLinkChainAccount", nil)
	cdc.RegisterConcrete(&MsgUnlinkChainAccount{}, "mage/MsgUnlinkChainAccount", nil)
	cdc.RegisterConcrete(&MsgSetDefaultExternalAddress{}, "mage/MsgSetDefaultExternalAddress", nil)
	cdc.RegisterConcrete(&MsgLinkApplication{}, "mage/MsgLinkApplication", nil)
	cdc.RegisterConcrete(&MsgUnlinkApplication{}, "mage/MsgUnlinkApplication", nil)

	cdc.RegisterInterface((*AddressData)(nil), nil)
	cdc.RegisterConcrete(&Bech32Address{}, "mage/Bech32Address", nil)
	cdc.RegisterConcrete(&Base58Address{}, "mage/Base58Address", nil)
	cdc.RegisterConcrete(&HexAddress{}, "mage/HexAddress", nil)

	cdc.RegisterInterface((*Signature)(nil), nil)
	cdc.RegisterConcrete(&SingleSignature{}, "mage/SingleSignature", nil)
	cdc.RegisterConcrete(&CosmosMultiSignature{}, "mage/CosmosMultiSignature", nil)

	cdc.RegisterConcrete(&Profile{}, "mage/Profile", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*authtypes.AccountI)(nil), &Profile{})
	registry.RegisterImplementations((*exported.VestingAccount)(nil), &Profile{})
	registry.RegisterImplementations((*authtypes.GenesisAccount)(nil), &Profile{})
	registry.RegisterInterface(
		"mage.profiles.v3.AddressData",
		(*AddressData)(nil),
		&Bech32Address{},
		&Base58Address{},
		&HexAddress{},
	)
	registry.RegisterInterface(
		"mage.profiles.v3.Signature",
		(*Signature)(nil),
		&SingleSignature{},
		&CosmosMultiSignature{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSaveProfile{},
		&MsgDeleteProfile{},
		&MsgRequestDTagTransfer{},
		&MsgCancelDTagTransferRequest{},
		&MsgAcceptDTagTransferRequest{},
		&MsgRefuseDTagTransferRequest{},
		&MsgLinkChainAccount{},
		&MsgUnlinkChainAccount{},
		&MsgLinkApplication{},
		&MsgUnlinkApplication{},
		&MsgSetDefaultExternalAddress{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// AminoCdc references the global x/relationships module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/relationships and
	// defined at the application level.
	AminoCdc = codec.NewAminoCodec(amino)

	ModuleCdc = codec.NewProtoCodec(types.NewInterfaceRegistry())
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)

	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
}
