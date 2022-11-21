package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/johnreitano/numi/x/numi/types"
)

func (k msgServer) CreateAndVerifyUser(goCtx context.Context, msg *types.MsgCreateAndVerifyUser) (*types.MsgCreateAndVerifyUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx

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

	k.Keeper.SetUser(ctx, user)

	// TODO: add to upcoming auctions

	// TODO: emit an event
	// ctx.EventManager().EmitEvent(
	// 	sdk.NewEvent(types.UserCreatedEventType,
	// 		sdk.NewAttribute(types.UserCreatedEventUserId, msg.UserId),
	// 		sdk.NewAttribute(types.UserCreatedEventFirstName, msg.FirstName),
	// 		sdk.NewAttribute(types.UserCreatedEventLastName, msg.LastName),
	// 		sdk.NewAttribute(types.UserCreatedEventReferrer, msg.Referrer),
	// 		sdk.NewAttribute(types.UserCreatedEventCreator, msg.Creator),
	// 	),
	// )}

	return &types.MsgCreateAndVerifyUserResponse{}, nil
}
