package keeper_test

import (
	"context"
	"testing"

	keepertest "ccs2/testutil/keeper"
	"ccs2/x/ccs2/keeper"
	"ccs2/x/ccs2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.Ccs2Keeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
