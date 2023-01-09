package provider

import (
	"github.com/mage-war/mage/app/mage/cmd/chainlink/builder"
	multibuilder "github.com/mage-war/mage/app/mage/cmd/chainlink/builder/multi"
	singlebuilder "github.com/mage-war/mage/app/mage/cmd/chainlink/builder/single"
	multigetter "github.com/mage-war/mage/app/mage/cmd/chainlink/getter/multi"
	singlegetter "github.com/mage-war/mage/app/mage/cmd/chainlink/getter/single"
)

// DefaultChainLinkJSONBuilderProvider returns the default ChainLinkJSONBuilder provider implementation
func DefaultChainLinkJSONBuilderProvider(owner string, isSingleAccount bool) builder.ChainLinkJSONBuilder {
	if isSingleAccount {
		return singlebuilder.NewAccountChainLinkJSONBuilder(owner, singlegetter.NewChainLinkJSONInfoGetter())
	}
	return multibuilder.NewAccountChainLinkJSONBuilder(multigetter.NewChainLinkJSONInfoGetter())
}
