package keeper_test

import (
	"context"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jinzhu/copier"
	"github.com/johnreitano/numi/app"
	keepertest "github.com/johnreitano/numi/testutil/keeper"
	"github.com/johnreitano/numi/x/numi"
	"github.com/johnreitano/numi/x/numi/keeper"
	"github.com/johnreitano/numi/x/numi/types"
	"github.com/stretchr/testify/require"
)

const (
	alice  = "numi12utcwzwmrwzsjj362qsht6zapq9cd5q0gh8pgy"
	bob    = "numi1hq9ujvq7yl2krs54xez67pfhfkpg8nlpqa75u2"
	carol  = "numi1staes5penzmhsxk3vmh4uwf483yejnvvx2ljwe"
	oliver = "numi13crpqdukn5l3gr4pzzcjzcl6fpx7rhay8uvy44"
	olivia = "numi1tsacr4aqrrerakdlcmlzl7daplle54fj874w2s"
)

func setupMsgServerCreateAndVerifyUser(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	app.SetAddressPrefixesInSDKConfig()

	k, ctx := keepertest.NumiKeeper(t)

	numi.InitGenesis(ctx, *k, *types.DefaultGenesis())

	verifiers := fmt.Sprintf("%s,%s", oliver, olivia)
	k.SetParams(ctx, types.NewParams(verifiers))

	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}

func TestCreateAndVerifyUser_ResponseIsAsExpected(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateAndVerifyUser(t)
	createResponse, err := msgServer.CreateAndVerifyUser(context, validMessage())
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateAndVerifyUserResponse{}, *createResponse)
}

func TestCreateAndVerifyUser_UserIsSaved(t *testing.T) {
	msgServer, keeper, context := setupMsgServerCreateAndVerifyUser(t)

	message := validMessage()
	_, err := msgServer.CreateAndVerifyUser(context, message)
	require.Nil(t, err)

	var expectedUser types.User
	copier.Copy(&expectedUser, message)
	actualUser, found := keeper.GetUser(sdk.UnwrapSDKContext(context), message.UserId)
	require.True(t, found)
	require.EqualValues(t, actualUser, expectedUser)
}

func TestCreateAndVerifyUser_FailsIfMsgFieldInvalid(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateAndVerifyUser(t)

	message := validMessage()
	message.FirstName = ""
	_, err := msgServer.CreateAndVerifyUser(context, message)
	require.ErrorIs(t, err, types.ErrFirstNameBlank)
}

func TestCreateAndVerifyUser_FailsIfCreatorNotIdentityProvider(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateAndVerifyUser(t)

	message := validMessage()
	message.Creator = alice
	_, err := msgServer.CreateAndVerifyUser(context, message)
	require.ErrorIs(t, err, types.ErrCreatorNotAuthorizedToVerifyIdentities)
}

func TestCreateAndVerifyUser_FailsIfUserIdAlreadyInUse(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateAndVerifyUser(t)

	message := validMessage()
	message.AccountAddress = bob
	_, err := msgServer.CreateAndVerifyUser(context, message)
	require.Nil(t, err)

	_, err = msgServer.CreateAndVerifyUser(context, message)
	require.ErrorIs(t, err, types.ErrUserIdAlreadyExists)
}

// TODO: add test for ErrAccountAddressAlreadyExists
// func TestCreateAndVerifyUser_FailsIfAccountAddressAlreadyInUse(t *testing.T) {
// 	msgServer, _, context := setupMsgServerCreateAndVerifyUser(t)

// 	message := validMessage()
// 	message.UserId = bob
// 	_, err := msgServer.CreateAndVerifyUser(context, message)
// 	require.Nil(t, err)

// 	_, err = msgServer.CreateAndVerifyUser(context, message)
// 	require.ErrorIs(t, err, types.ErrAccountAddressAlreadyExists)
// }

func validMessage() *types.MsgCreateAndVerifyUser {
	return &types.MsgCreateAndVerifyUser{
		Creator:           olivia,
		UserId:            "1bc3e020-2b02-40a7-abd8-eadc9b4250c5",
		FirstName:         "John",
		LastName:          "Doe",
		CountryCode:       "USA",
		SubnationalEntity: "California",
		City:              "San Diego",
		Bio:               "a serious man",
		Referrer:          alice,
		AccountAddress:    bob,
	}
}
