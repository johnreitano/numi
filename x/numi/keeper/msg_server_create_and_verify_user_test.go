package keeper_test

import (
	"context"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/copier"
	"github.com/johnreitano/numi/app"
	keepertest "github.com/johnreitano/numi/testutil/keeper"
	"github.com/johnreitano/numi/testutil/mock_types"
	typestest "github.com/johnreitano/numi/testutil/types"

	"github.com/johnreitano/numi/x/numi"
	"github.com/johnreitano/numi/x/numi/keeper"
	"github.com/johnreitano/numi/x/numi/types"
	"github.com/stretchr/testify/require"
)

type coinMatcher struct {
	sdk.Coin
}

func (m coinMatcher) Matches(x interface{}) bool {
	arg, ok := x.(sdk.Coins)
	if !ok {
		return false
	}
	return arg[0].Denom == m.Denom && arg[0].Amount.Equal(m.Amount)
}

func (m coinMatcher) String() string {
	return fmt.Sprintf("matches demon %s and amount %d", m.Denom, m.Amount)
}

func FirstCoinMatches(coin sdk.Coin) gomock.Matcher {
	return coinMatcher{coin}
}

type addressMatcher string

func (m addressMatcher) Matches(x interface{}) bool {
	arg, ok := x.(sdk.AccAddress)
	if !ok {
		return false
	}

	expectedAddress := string(m)
	actualAddress := arg.String()

	return expectedAddress == actualAddress
}

func (m addressMatcher) String() string {
	return fmt.Sprintf("matches address %s", string(m))
}

func AddressMatches(expectedRecipient string) gomock.Matcher {
	return addressMatcher(expectedRecipient)
}

func setupMsgServerCreateAndVerifyUserWithMocks(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	ctrl := gomock.NewController(t)
	bankKeeperMock := mock_types.NewMockBankKeeper(ctrl)
	mintKeeperMock := mock_types.NewMockMintKeeper(ctrl)
	msgServer, keeper, context := setupMsgServerCreateAndVerifyUser(t, bankKeeperMock, mintKeeperMock)
	bankKeeperMock.ExpectAny(context)
	mintKeeperMock.ExpectAny(context)
	return msgServer, keeper, context
}

func setupMsgServerCreateAndVerifyUser(t testing.TB, bankKeeper types.BankKeeper, mintKeeper types.MintKeeper) (types.MsgServer, keeper.Keeper, context.Context) {
	app.SetAddressPrefixesInSDKConfig()

	k, ctx := keepertest.NumiKeeper(t, bankKeeper, mintKeeper)
	numi.InitGenesis(ctx, *k, *types.DefaultGenesis())

	verifiers := fmt.Sprintf("%s,%s", typestest.Oliver, typestest.Olivia)
	k.SetParams(ctx, types.NewParams(verifiers))

	c := sdk.WrapSDKContext(ctx)
	return keeper.NewMsgServerImpl(*k), *k, c
}

// TODO: if appropriate for additional tests, add IntegrationTestSuite as in b9 app and use real BankKeeper and MintKeeper instead of mocks

func TestCreateAndVerifyUser_ResponseIsAsExpected(t *testing.T) {
	ctrl := gomock.NewController(t)
	bankKeeperMock := mock_types.NewMockBankKeeper(ctrl)
	mintKeeperMock := mock_types.NewMockMintKeeper(ctrl)
	msgServer, _, context := setupMsgServerCreateAndVerifyUser(t, bankKeeperMock, mintKeeperMock)

	bankKeeperMock.ExpectAny(context)
	mintKeeperMock.ExpectAny(context)
	msg := typestest.ValidMsgCreateAndVerifyUser()
	createResponse, err := msgServer.CreateAndVerifyUser(context, msg)
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateAndVerifyUserResponse{}, *createResponse)
}

func TestCreateAndVerifyUser_UserIsSaved(t *testing.T) {
	msgServer, keeper, context := setupMsgServerCreateAndVerifyUserWithMocks(t)

	message := typestest.ValidMsgCreateAndVerifyUser()
	_, err := msgServer.CreateAndVerifyUser(context, message)
	require.Nil(t, err)

	var expectedUser types.User
	copier.Copy(&expectedUser, message)
	actualUser, found := keeper.GetUser(sdk.UnwrapSDKContext(context), message.UserId)
	require.True(t, found)
	require.EqualValues(t, actualUser, expectedUser)
}

func TestCreateAndVerifyUser_MintsAndTransfersRewardToCreator(t *testing.T) {
	ctrl := gomock.NewController(t)
	bankKeeperMock := mock_types.NewMockBankKeeper(ctrl)
	mintKeeperMock := mock_types.NewMockMintKeeper(ctrl)
	msgServer, _, context := setupMsgServerCreateAndVerifyUser(t, bankKeeperMock, mintKeeperMock)

	mintKeeperMock.EXPECT().GetParams(sdk.UnwrapSDKContext(context)).Return(minttypes.Params{MintDenom: "unumi"}).AnyTimes()

	msg := typestest.ValidMsgCreateAndVerifyUser()

	expectedReward := sdk.Coin{Denom: "unumi", Amount: sdk.NewInt(10)}
	bankKeeperMock.EXPECT().MintCoins(sdk.UnwrapSDKContext(context), "numi", FirstCoinMatches(expectedReward)).Times(1)

	bankKeeperMock.EXPECT().SendCoinsFromModuleToAccount(sdk.UnwrapSDKContext(context), "numi", AddressMatches(msg.Creator), FirstCoinMatches(expectedReward)).Times(1)

	_, err := msgServer.CreateAndVerifyUser(context, msg)
	require.Nil(t, err)
}

func TestCreateAndVerifyUser_Create1GameEmitted(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateAndVerifyUserWithMocks(t)
	ctx := sdk.UnwrapSDKContext(context)
	message := typestest.ValidMsgCreateAndVerifyUser()
	_, err := msgServer.CreateAndVerifyUser(context, message)
	require.Nil(t, err)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 1)
	event := events[0]
	require.EqualValues(t, sdk.StringEvent{
		Type: "user-created-and-verified",
		Attributes: []sdk.Attribute{
			{Key: "user-id", Value: message.UserId},
			{Key: "first-name", Value: message.FirstName},
			{Key: "last-name", Value: message.LastName},
			{Key: "country-code", Value: message.CountryCode},
			{Key: "subnational-entity", Value: message.SubnationalEntity},
			{Key: "city", Value: message.City},
			{Key: "bio", Value: message.Bio},
			{Key: "creator", Value: message.Creator},
			{Key: "referrer", Value: message.Referrer},
			{Key: "account-address", Value: message.AccountAddress},
		},
	}, event)
}

func TestCreateAndVerifyUser_FailsIfValidateBasicFailsForUser(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateAndVerifyUserWithMocks(t)
	message := typestest.ValidMsgCreateAndVerifyUser()
	message.FirstName = ""
	_, err := msgServer.CreateAndVerifyUser(context, message)
	require.NotNil(t, err)

}

func TestCreateAndVerifyUser_FailsIfCreatorNotIdentityProvider(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateAndVerifyUserWithMocks(t)

	message := typestest.ValidMsgCreateAndVerifyUser()
	message.Creator = typestest.Alice
	_, err := msgServer.CreateAndVerifyUser(context, message)
	require.ErrorIs(t, err, types.ErrCreatorNotAuthorizedToVerifyIdentities)
}

func TestCreateAndVerifyUser_FailsIfUserIdAlreadyInUse(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateAndVerifyUserWithMocks(t)

	message := typestest.ValidMsgCreateAndVerifyUser()
	_, err := msgServer.CreateAndVerifyUser(context, message)
	require.Nil(t, err)

	message.AccountAddress = typestest.Bob
	_, err = msgServer.CreateAndVerifyUser(context, message)
	require.ErrorIs(t, err, types.ErrUserIdAlreadyExists)
}

func TestCreateAndVerifyUser_FailsIfAccountAddressAlreadyInUse(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateAndVerifyUserWithMocks(t)

	message := typestest.ValidMsgCreateAndVerifyUser()
	_, err := msgServer.CreateAndVerifyUser(context, message)
	require.Nil(t, err)

	newUserId := "74b76d9e-bf21-4b50-93b7-2bdbb9dcf926"
	require.NotEqual(t, message.UserId, newUserId)
	message.UserId = newUserId
	_, err = msgServer.CreateAndVerifyUser(context, message)
	require.ErrorIs(t, err, types.ErrAccountAddressAlreadyExists)
}
