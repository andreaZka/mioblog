package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/andreaZka/mioblog/testutil/keeper"
	"github.com/andreaZka/mioblog/testutil/nullify"
	"github.com/andreaZka/mioblog/x/mioblog/keeper"
	"github.com/andreaZka/mioblog/x/mioblog/types"
)

func createTestNextPost(keeper *keeper.Keeper, ctx sdk.Context) types.NextPost {
	item := types.NextPost{}
	keeper.SetNextPost(ctx, item)
	return item
}

func TestNextPostGet(t *testing.T) {
	keeper, ctx := keepertest.MioblogKeeper(t)
	item := createTestNextPost(keeper, ctx)
	rst, found := keeper.GetNextPost(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestNextPostRemove(t *testing.T) {
	keeper, ctx := keepertest.MioblogKeeper(t)
	createTestNextPost(keeper, ctx)
	keeper.RemoveNextPost(ctx)
	_, found := keeper.GetNextPost(ctx)
	require.False(t, found)
}
