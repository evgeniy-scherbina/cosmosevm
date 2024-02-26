package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/evgeniy-scherbina/cosmosevm/testutil/keeper"
	"github.com/evgeniy-scherbina/cosmosevm/testutil/nullify"
	"github.com/evgeniy-scherbina/cosmosevm/x/monobank/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestBalanceQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.MonobankKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNBalance(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetBalanceRequest
		response *types.QueryGetBalanceResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetBalanceRequest{
				Address: msgs[0].Address,
			},
			response: &types.QueryGetBalanceResponse{Balance: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetBalanceRequest{
				Address: msgs[1].Address,
			},
			response: &types.QueryGetBalanceResponse{Balance: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetBalanceRequest{
				Address: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Balance(wctx, tc.request)
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

func TestBalanceQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.MonobankKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNBalance(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllBalanceRequest {
		return &types.QueryAllBalanceRequest{
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
			resp, err := keeper.BalanceAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Balance), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Balance),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.BalanceAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Balance), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Balance),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.BalanceAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Balance),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.BalanceAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
