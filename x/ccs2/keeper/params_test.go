package keeper_test

import (
	"testing"

	testkeeper "ccs2/testutil/keeper"
	"ccs2/x/ccs2/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.Ccs2Keeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
