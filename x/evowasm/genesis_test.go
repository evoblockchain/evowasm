package evowasm_test

import (
	"testing"

	keepertest "evowasm/testutil/keeper"
	"evowasm/testutil/nullify"
	"evowasm/x/evowasm"
	"evowasm/x/evowasm/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EvowasmKeeper(t)
	evowasm.InitGenesis(ctx, *k, genesisState)
	got := evowasm.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
