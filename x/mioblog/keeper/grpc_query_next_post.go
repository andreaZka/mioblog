package keeper

import (
	"context"

	"github.com/andreaZka/mioblog/x/mioblog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NextPost(c context.Context, req *types.QueryGetNextPostRequest) (*types.QueryGetNextPostResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNextPost(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNextPostResponse{NextPost: val}, nil
}
