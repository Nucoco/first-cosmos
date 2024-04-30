package firstcosmos_test

import (
	"testing"

	keepertest "first-cosmos/testutil/keeper"
	"first-cosmos/testutil/nullify"
	firstcosmos "first-cosmos/x/firstcosmos/module"
	"first-cosmos/x/firstcosmos/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FirstcosmosKeeper(t)
	firstcosmos.InitGenesis(ctx, k, genesisState)
	got := firstcosmos.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
