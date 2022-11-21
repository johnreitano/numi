package keeper_test

import (
	"context"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jinzhu/copier"
	"github.com/johnreitano/numi/app"
	keepertest "github.com/johnreitano/numi/testutil/keeper"
	typestest "github.com/johnreitano/numi/testutil/types"
	"github.com/johnreitano/numi/x/numi"
	"github.com/johnreitano/numi/x/numi/keeper"
	"github.com/johnreitano/numi/x/numi/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServerCreateAndVerifyUser(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	app.SetAddressPrefixesInSDKConfig()

	k, ctx := keepertest.NumiKeeper(t)

	numi.InitGenesis(ctx, *k, *types.DefaultGenesis())

	verifiers := fmt.Sprintf("%s,%s", typestest.Oliver, typestest.Olivia)
	k.SetParams(ctx, types.NewParams(verifiers))

	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}

func TestCreateAndVerifyUser_ResponseIsAsExpected(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateAndVerifyUser(t)
	createResponse, err := msgServer.CreateAndVerifyUser(context, typestest.ValidMsgCreateAndVerifyUser())
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateAndVerifyUserResponse{}, *createResponse)
}

func TestCreateAndVerifyUser_UserIsSaved(t *testing.T) {
	msgServer, keeper, context := setupMsgServerCreateAndVerifyUser(t)

	message := typestest.ValidMsgCreateAndVerifyUser()
	_, err := msgServer.CreateAndVerifyUser(context, message)
	require.Nil(t, err)

	var expectedUser types.User
	copier.Copy(&expectedUser, message)
	actualUser, found := keeper.GetUser(sdk.UnwrapSDKContext(context), message.UserId)
	require.True(t, found)
	require.EqualValues(t, actualUser, expectedUser)
}

func TestCreateAndVerifyUser_FailsIfValidateBasicFailsForUser(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateAndVerifyUser(t)
	message := typestest.ValidMsgCreateAndVerifyUser()
	message.FirstName = ""
	_, err := msgServer.CreateAndVerifyUser(context, message)
	require.NotNil(t, err)

}

func TestCreateAndVerifyUser_FailsIfCreatorNotIdentityProvider(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateAndVerifyUser(t)

	message := typestest.ValidMsgCreateAndVerifyUser()
	message.Creator = typestest.Alice
	_, err := msgServer.CreateAndVerifyUser(context, message)
	require.ErrorIs(t, err, types.ErrCreatorNotAuthorizedToVerifyIdentities)
}

func TestCreateAndVerifyUser_FailsIfUserIdAlreadyInUse(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateAndVerifyUser(t)

	message := typestest.ValidMsgCreateAndVerifyUser()
	_, err := msgServer.CreateAndVerifyUser(context, message)
	require.Nil(t, err)

	message.AccountAddress = typestest.Bob
	_, err = msgServer.CreateAndVerifyUser(context, message)
	require.ErrorIs(t, err, types.ErrUserIdAlreadyExists)
}

func TestCreateAndVerifyUser_FailsIfAccountAddressAlreadyInUse(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateAndVerifyUser(t)

	message := typestest.ValidMsgCreateAndVerifyUser()
	_, err := msgServer.CreateAndVerifyUser(context, message)
	require.Nil(t, err)

	newUserId := "74b76d9e-bf21-4b50-93b7-2bdbb9dcf926"
	require.NotEqual(t, message.UserId, newUserId)
	message.UserId = newUserId
	_, err = msgServer.CreateAndVerifyUser(context, message)
	require.ErrorIs(t, err, types.ErrAccountAddressAlreadyExists)
}

func TestCreate1GameEmitted(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateAndVerifyUser(t)
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
