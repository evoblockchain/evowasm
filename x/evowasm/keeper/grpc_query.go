package keeper

import (
	"evowasm/x/evowasm/types"
)

var _ types.QueryServer = Keeper{}
