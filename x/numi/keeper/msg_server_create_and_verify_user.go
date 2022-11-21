package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/johnreitano/numi/x/numi/types"
)

func (k msgServer) CreateAndVerifyUser(goCtx context.Context, msg *types.MsgCreateAndVerifyUser) (*types.MsgCreateAndVerifyUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user := types.User{
		Creator:           msg.Creator,
		UserId:            msg.UserId,
		FirstName:         msg.FirstName,
		LastName:          msg.LastName,
		CountryCode:       msg.CountryCode,
		SubnationalEntity: msg.SubnationalEntity,
		City:              msg.City,
		Bio:               msg.Bio,
		Referrer:          msg.Referrer,
		AccountAddress:    msg.AccountAddress,
	}
	err := types.ValidateUserBasic(&user)
	if err != nil {
		return nil, err
	}

	if !k.IsIdentityVerifier(ctx, user.Creator) {
		return nil, types.ErrCreatorNotAuthorizedToVerifyIdentities
	}

	if _, found := k.GetUser(ctx, msg.UserId); found {
		return nil, types.ErrUserIdAlreadyExists
	}

	if _, found := k.GetUserAccountAddress(ctx, msg.AccountAddress); found {
		return nil, types.ErrAccountAddressAlreadyExists
	}

	k.Keeper.SetUser(ctx, user)
	k.Keeper.SetUserAccountAddress(ctx, types.UserAccountAddress{
		AccountAddress: msg.AccountAddress,
		UserId:         msg.UserId,
	})

	// TODO: add msg to upcoming auctions

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.UserCreatedAndVerifiedEventType,
			sdk.NewAttribute(types.UserCreatedAndVerifiedEventUserId, msg.UserId),
			sdk.NewAttribute(types.UserCreatedAndVerifiedEventFirstName, msg.FirstName),
			sdk.NewAttribute(types.UserCreatedAndVerifiedEventLastName, msg.LastName),
			sdk.NewAttribute(types.UserCreatedAndVerifiedEventCountryCode, msg.CountryCode),
			sdk.NewAttribute(types.UserCreatedAndVerifiedEventSubnationalEntity, msg.SubnationalEntity),
			sdk.NewAttribute(types.UserCreatedAndVerifiedEventCity, msg.City),
			sdk.NewAttribute(types.UserCreatedAndVerifiedEventBio, msg.Bio),
			sdk.NewAttribute(types.UserCreatedAndVerifiedEventCreator, msg.Creator),
			sdk.NewAttribute(types.UserCreatedAndVerifiedEventReferrer, msg.Referrer),
			sdk.NewAttribute(types.UserCreatedAndVerifiedEventAccountAddress, msg.AccountAddress),
		),
	)

	return &types.MsgCreateAndVerifyUserResponse{}, nil
}
