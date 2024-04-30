package keeper

import (
	"first-cosmos/x/firstmodule/types"
)

var _ types.QueryServer = Keeper{}
