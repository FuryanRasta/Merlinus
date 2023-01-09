package types

// DONTCOVER

import (
	subspacestypes "github.com/mage-war/mage/x/subspaces/types"
)

var (
	PermissionsReact                    = subspacestypes.RegisterPermission("add reaction")
	PermissionManageRegisteredReactions = subspacestypes.RegisterPermission("manage registered reactions")
	PermissionManageReactionParams      = subspacestypes.RegisterPermission("manage reaction params")
)
