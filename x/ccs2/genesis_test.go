package ccs2_test

import (
	"testing"

	keepertest "ccs2/testutil/keeper"
	"ccs2/testutil/nullify"
	"ccs2/x/ccs2"
	"ccs2/x/ccs2/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Ccs2Keeper(t)
	ccs2.InitGenesis(ctx, *k, genesisState)
	got := ccs2.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
