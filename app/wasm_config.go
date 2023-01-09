package app

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/cosmos/cosmos-sdk/codec"

	wasmmage "github.com/mage-war/mage/cosmwasm"
	postskeeper "github.com/mage-war/mage/x/posts/keeper"
	postswasm "github.com/mage-war/mage/x/posts/wasm"
	profileskeeper "github.com/mage-war/mage/x/profiles/keeper"
	profileswasm "github.com/mage-war/mage/x/profiles/wasm"
	reactionskeeper "github.com/mage-war/mage/x/reactions/keeper"
	reactionswasm "github.com/mage-war/mage/x/reactions/wasm"
	relationshipskeeper "github.com/mage-war/mage/x/relationships/keeper"
	relationshipswasm "github.com/mage-war/mage/x/relationships/wasm"
	reportskeeper "github.com/mage-war/mage/x/reports/keeper"
	reportswasm "github.com/mage-war/mage/x/reports/wasm"
	subspaceskeeper "github.com/mage-war/mage/x/subspaces/keeper"
	subspaceswasm "github.com/mage-war/mage/x/subspaces/wasm"
)

const (
	// DefaultMageInstanceCost is how much SDK gas we charge each time we load a WASM instance
	DefaultMageInstanceCost uint64 = 60_000
	// DefaultMageCompileCost is how much SDK gas is charged *per byte* for compiling WASM code
	DefaultMageCompileCost uint64 = 2
)

// MageWasmGasRegister is defaults plus a custom compile amount
func MageWasmGasRegister() wasmkeeper.WasmGasRegisterConfig {
	gasConfig := wasmkeeper.DefaultGasRegisterConfig()
	gasConfig.InstanceCost = DefaultMageInstanceCost
	gasConfig.CompileCost = DefaultMageCompileCost

	return gasConfig
}

func NewMageWasmGasRegister() wasmkeeper.WasmGasRegister {
	return wasmkeeper.NewWasmGasRegister(MageWasmGasRegister())
}

// NewMageCustomQueryPlugin initialize the custom querier to handle mage queries for contracts
func NewMageCustomQueryPlugin(
	cdc codec.Codec,
	profilesKeeper profileskeeper.Keeper,
	subspacesKeeper subspaceskeeper.Keeper,
	relationshipsKeeper relationshipskeeper.Keeper,
	postsKeeper postskeeper.Keeper,
	reportsKeeper reportskeeper.Keeper,
	reactionsKeeper reactionskeeper.Keeper,
) wasm.QueryPlugins {
	queriers := map[string]wasmmage.Querier{
		wasmmage.QueryRouteProfiles:      profileswasm.NewProfilesWasmQuerier(profilesKeeper, cdc),
		wasmmage.QueryRouteSubspaces:     subspaceswasm.NewSubspacesWasmQuerier(subspacesKeeper, cdc),
		wasmmage.QueryRouteRelationships: relationshipswasm.NewRelationshipsWasmQuerier(relationshipsKeeper, cdc),
		wasmmage.QueryRoutePosts:         postswasm.NewPostsWasmQuerier(postsKeeper, cdc),
		wasmmage.QueryRouteReports:       reportswasm.NewReportsWasmQuerier(reportsKeeper, cdc),
		wasmmage.QueryRouteReactions:     reactionswasm.NewReactionsWasmQuerier(reactionsKeeper, cdc),
		// add other modules querier here
	}

	querier := wasmmage.NewQuerier(queriers)

	return wasm.QueryPlugins{
		Custom: querier.QueryCustom,
	}
}

// NewMageCustomMessageEncoder initialize the custom message encoder to mage app for contracts
func NewMageCustomMessageEncoder(cdc codec.Codec) wasm.MessageEncoders {
	// Initialization of custom Mage messages for contracts
	parserRouter := wasmmage.NewParserRouter()
	parsers := map[string]wasmmage.MsgParserInterface{
		wasmmage.WasmMsgParserRouteProfiles:      profileswasm.NewWasmMsgParser(cdc),
		wasmmage.WasmMsgParserRouteSubspaces:     subspaceswasm.NewWasmMsgParser(cdc),
		wasmmage.WasmMsgParserRouteRelationships: relationshipswasm.NewWasmMsgParser(cdc),
		wasmmage.WasmMsgParserRoutePosts:         postswasm.NewWasmMsgParser(cdc),
		wasmmage.WasmMsgParserRouteReports:       reportswasm.NewWasmMsgParser(cdc),
		wasmmage.WasmMsgParserRouteReactions:     reactionswasm.NewWasmMsgParser(cdc),
		// add other modules parsers here
	}

	parserRouter.Parsers = parsers
	return wasm.MessageEncoders{
		Custom: parserRouter.ParseCustom,
	}
}
