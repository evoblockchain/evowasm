package keeper_test

import (
	"testing"

	testkeeper "evowasm/testutil/keeper"
	"evowasm/x/evowasm/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.EvowasmKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
