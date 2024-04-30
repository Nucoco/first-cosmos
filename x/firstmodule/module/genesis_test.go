package firstmodule_test

import (
	"testing"

	keepertest "first-cosmos/testutil/keeper"
	"first-cosmos/testutil/nullify"
	firstmodule "first-cosmos/x/firstmodule/module"
	"first-cosmos/x/firstmodule/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FirstmoduleKeeper(t)
	firstmodule.InitGenesis(ctx, k, genesisState)
	got := firstmodule.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
