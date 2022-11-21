package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/johnreitano/numi/testutil/keeper"
	"github.com/johnreitano/numi/testutil/nullify"
	"github.com/johnreitano/numi/x/numi/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestUserAccountAddressQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.NumiKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNUserAccountAddress(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetUserAccountAddressRequest
		response *types.QueryGetUserAccountAddressResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetUserAccountAddressRequest{
				AccountAddress: msgs[0].AccountAddress,
			},
			response: &types.QueryGetUserAccountAddressResponse{UserAccountAddress: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetUserAccountAddressRequest{
				AccountAddress: msgs[1].AccountAddress,
			},
			response: &types.QueryGetUserAccountAddressResponse{UserAccountAddress: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetUserAccountAddressRequest{
				AccountAddress: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.UserAccountAddress(wctx, tc.request)
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

func TestUserAccountAddressQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.NumiKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNUserAccountAddress(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllUserAccountAddressRequest {
		return &types.QueryAllUserAccountAddressRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.UserAccountAddressAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.UserAccountAddress), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.UserAccountAddress),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.UserAccountAddressAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.UserAccountAddress), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.UserAccountAddress),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.UserAccountAddressAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.UserAccountAddress),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.UserAccountAddressAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
