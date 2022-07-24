package keeper

import (
	"context"
	"strconv"

	"github.com/andreaZka/mioblog/x/mioblog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	nextPost, found := k.Keeper.GetNextPost(ctx)
	if !found {
		//	panic("NextGame not found")
	}
	newIndex := strconv.FormatUint(nextPost.IdValue, 10)

	// Create variable of type Post
	var post = types.Post{
		Index: newIndex,
		Title: msg.Title,
		Body:  msg.Body,
	}
	k.Keeper.SetPost(ctx, post)

	nextPost.IdValue++
	k.Keeper.SetNextPost(ctx, nextPost)

	return &types.MsgCreatePostResponse{IdValue: newIndex}, nil
}
