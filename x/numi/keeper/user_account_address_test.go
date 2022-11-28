package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/johnreitano/numi/testutil/keeper"
	"github.com/johnreitano/numi/testutil/nullify"
	"github.com/johnreitano/numi/x/numi/keeper"
	"github.com/johnreitano/numi/x/numi/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNUserAccountAddress(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.UserAccountAddress {
	items := make([]types.UserAccountAddress, n)
	for i := range items {
		items[i].AccountAddress = strconv.Itoa(i)

		keeper.SetUserAccountAddress(ctx, items[i])
	}
	return items
}

func TestUserAccountAddressGet(t *testing.T) {
	keeper, ctx := keepertest.NumiKeeperWithMocks(t)
	items := createNUserAccountAddress(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetUserAccountAddress(ctx,
			item.AccountAddress,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestUserAccountAddressRemove(t *testing.T) {
	keeper, ctx := keepertest.NumiKeeperWithMocks(t)
	items := createNUserAccountAddress(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUserAccountAddress(ctx,
			item.AccountAddress,
		)
		_, found := keeper.GetUserAccountAddress(ctx,
			item.AccountAddress,
		)
		require.False(t, found)
	}
}

func TestUserAccountAddressGetAll(t *testing.T) {
	keeper, ctx := keepertest.NumiKeeperWithMocks(t)
	items := createNUserAccountAddress(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllUserAccountAddress(ctx)),
	)
}
