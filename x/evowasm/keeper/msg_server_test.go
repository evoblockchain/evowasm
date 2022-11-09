package keeper_test

import (
	"context"
	"testing"

	keepertest "evowasm/testutil/keeper"
	"evowasm/x/evowasm/keeper"
	"evowasm/x/evowasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.EvowasmKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
