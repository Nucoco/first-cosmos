package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "first-cosmos/testutil/keeper"
	"first-cosmos/x/firstcosmos/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.FirstcosmosKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
