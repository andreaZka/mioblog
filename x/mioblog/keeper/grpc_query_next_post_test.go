package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/andreaZka/mioblog/testutil/keeper"
	"github.com/andreaZka/mioblog/testutil/nullify"
	"github.com/andreaZka/mioblog/x/mioblog/types"
)

func TestNextPostQuery(t *testing.T) {
	keeper, ctx := keepertest.MioblogKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestNextPost(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetNextPostRequest
		response *types.QueryGetNextPostResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetNextPostRequest{},
			response: &types.QueryGetNextPostResponse{NextPost: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.NextPost(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
