package mioblog_test

import (
	"testing"

	keepertest "github.com/andreaZka/mioblog/testutil/keeper"
	"github.com/andreaZka/mioblog/testutil/nullify"
	"github.com/andreaZka/mioblog/x/mioblog"
	"github.com/andreaZka/mioblog/x/mioblog/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MioblogKeeper(t)
	mioblog.InitGenesis(ctx, *k, genesisState)
	got := mioblog.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
