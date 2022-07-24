package mioblog

import (
	"github.com/andreaZka/mioblog/x/mioblog/keeper"
	"github.com/andreaZka/mioblog/x/mioblog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the post
	for _, elem := range genState.PostList {
		k.SetPost(ctx, elem)
	}
	// Set if defined
	if genState.NextPost != nil {
		k.SetNextPost(ctx, *genState.NextPost)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PostList = k.GetAllPost(ctx)
	// Get all nextPost
	nextPost, found := k.GetNextPost(ctx)
	if found {
		genesis.NextPost = &nextPost
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
