package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/andreaZka/mioblog/testutil/keeper"
	"github.com/andreaZka/mioblog/x/mioblog/keeper"
	"github.com/andreaZka/mioblog/x/mioblog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MioblogKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
