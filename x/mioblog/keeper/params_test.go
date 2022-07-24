package keeper_test

import (
	"testing"

	testkeeper "github.com/andreaZka/mioblog/testutil/keeper"
	"github.com/andreaZka/mioblog/x/mioblog/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MioblogKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
