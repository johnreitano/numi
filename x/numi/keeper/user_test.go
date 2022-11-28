package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/johnreitano/numi/app"
	keepertest "github.com/johnreitano/numi/testutil/keeper"
	"github.com/johnreitano/numi/testutil/nullify"
	typestest "github.com/johnreitano/numi/testutil/types"
	"github.com/johnreitano/numi/x/numi/keeper"
	"github.com/johnreitano/numi/x/numi/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNUser(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.User {
	items := make([]types.User, n)
	for i := range items {
		items[i].UserId = strconv.Itoa(i)

		keeper.SetUser(ctx, items[i])
	}
	return items
}

func TestUserGet(t *testing.T) {
	keeper, ctx := keepertest.NumiKeeperWithMocks(t)
	items := createNUser(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetUser(ctx,
			item.UserId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestUserRemove(t *testing.T) {
	keeper, ctx := keepertest.NumiKeeperWithMocks(t)
	items := createNUser(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUser(ctx,
			item.UserId,
		)
		_, found := keeper.GetUser(ctx,
			item.UserId,
		)
		require.False(t, found)
	}
}

func TestUserGetAll(t *testing.T) {
	keeper, ctx := keepertest.NumiKeeperWithMocks(t)
	items := createNUser(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllUser(ctx)),
	)
}

func TestValidateBasic(t *testing.T) {
	app.SetAddressPrefixesInSDKConfig()

	validUser := typestest.ValidUser()
	for _, tc := range []struct {
		desc  string
		user  func() *types.User
		valid bool
	}{
		{
			desc:  "user with all fields valid",
			user:  func() *types.User { return validUser },
			valid: true,
		},
		{
			desc: "missing creator",
			user: func() *types.User {
				user := *validUser
				user.Creator = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "invalid creator",
			user: func() *types.User {
				user := *validUser
				user.Creator = "xxx"
				return &user
			},
			valid: false,
		},
		{
			desc: "missing user id",
			user: func() *types.User {
				user := *validUser
				user.UserId = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "invalid user id",
			user: func() *types.User {
				user := *validUser
				user.UserId = "xxx"
				return &user
			},
			valid: false,
		},
		{
			desc: "missing first name",
			user: func() *types.User {
				user := *validUser
				user.FirstName = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "missing last name",
			user: func() *types.User {
				user := *validUser
				user.LastName = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "missing country code",
			user: func() *types.User {
				user := *validUser
				user.CountryCode = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "2-digit country code",
			user: func() *types.User {
				user := *validUser
				user.CountryCode = "US"
				return &user
			},
			valid: false,
		},
		{
			desc: "invalid 3-digit country code",
			user: func() *types.User {
				user := *validUser
				user.CountryCode = "XXX"
				return &user
			},
			valid: false,
		},
		{
			desc: "missing subnational entity",
			user: func() *types.User {
				user := *validUser
				user.SubnationalEntity = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "missing city",
			user: func() *types.User {
				user := *validUser
				user.City = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "missing bio",
			user: func() *types.User {
				user := *validUser
				user.Bio = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "missing referrer",
			user: func() *types.User {
				user := *validUser
				user.Referrer = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "invalid referrer",
			user: func() *types.User {
				user := *validUser
				user.Referrer = "xxx"
				return &user
			},
			valid: false,
		},
		{
			desc: "missing account address",
			user: func() *types.User {
				user := *validUser
				user.AccountAddress = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "invalid account address",
			user: func() *types.User {
				user := *validUser
				user.AccountAddress = "xxx"
				return &user
			},
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := types.ValidateUserBasic(tc.user())
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
