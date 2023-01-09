package types

// DONTCOVER

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*AttachmentContent)(nil), nil)
	cdc.RegisterConcrete(&Poll{}, "mage/Poll", nil)
	cdc.RegisterConcrete(&Media{}, "mage/Media", nil)

	cdc.RegisterConcrete(&MsgCreatePost{}, "mage/MsgCreatePost", nil)
	cdc.RegisterConcrete(&MsgEditPost{}, "mage/MsgEditPost", nil)
	cdc.RegisterConcrete(&MsgAddPostAttachment{}, "mage/MsgAddPostAttachment", nil)
	cdc.RegisterConcrete(&MsgRemovePostAttachment{}, "mage/MsgRemovePostAttachment", nil)
	cdc.RegisterConcrete(&MsgDeletePost{}, "mage/MsgDeletePost", nil)
	cdc.RegisterConcrete(&MsgAnswerPoll{}, "mage/MsgAnswerPoll", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterInterface(
		"mage.posts.v2.AttachmentContent",
		(*AttachmentContent)(nil),
		&Poll{},
		&Media{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePost{},
		&MsgEditPost{},
		&MsgAddPostAttachment{},
		&MsgRemovePostAttachment{},
		&MsgDeletePost{},
		&MsgAnswerPoll{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// AminoCodec references the global x/posts module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/posts and
	// defined at the application level.
	AminoCodec = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)

	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
}
