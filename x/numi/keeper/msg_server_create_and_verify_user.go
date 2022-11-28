package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/johnreitano/numi/x/numi/types"
)

const verificationRewardAmount = int64(10)

func (k msgServer) CreateAndVerifyUser(goCtx context.Context, msg *types.MsgCreateAndVerifyUser) (*types.MsgCreateAndVerifyUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	log := k.Keeper.Logger(ctx)

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
		log.Info("ValidateUserBasic returned error %s", err.Error())
		return nil, err
	}

	if !k.IsIdentityVerifier(ctx, user.Creator) {
		log.Info("Creator is not authorized: %s", types.ErrCreatorNotAuthorizedToVerifyIdentities.Error())
		return nil, types.ErrCreatorNotAuthorizedToVerifyIdentities
	}

	if _, found := k.GetUser(ctx, msg.UserId); found {
		log.Info("Creator is not authorized: %s", types.ErrUserIdAlreadyExists.Error())
		return nil, types.ErrUserIdAlreadyExists
	}

	if _, found := k.GetUserAccountAddress(ctx, msg.AccountAddress); found {
		log.Info("User already exists: %s", types.ErrAccountAddressAlreadyExists.Error())
		return nil, types.ErrAccountAddressAlreadyExists
	}

	k.Keeper.SetUser(ctx, user)
	k.Keeper.SetUserAccountAddress(ctx, types.UserAccountAddress{
		AccountAddress: msg.AccountAddress,
		UserId:         msg.UserId,
	})

	k.grantVerificationReward(ctx, msg.Creator)

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

	log.Info("user successfully created and verified")
	return &types.MsgCreateAndVerifyUserResponse{}, nil
}

func (k msgServer) grantVerificationReward(ctx sdk.Context, verifier string) error {
	r := k.verificationReward(ctx)
	return k.mintTo(ctx, r, verifier)
}

func (k msgServer) verificationReward(ctx sdk.Context) sdk.Coin {
	p := k.mintKeeper.GetParams(ctx)
	return sdk.Coin{
		Denom:  p.MintDenom,
		Amount: sdk.NewInt(verificationRewardAmount),
	}
}

func (k msgServer) mintTo(ctx sdk.Context, amount sdk.Coin, mintTo string) error {
	if err := k.Keeper.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(amount)); err != nil {
		return err
	}

	addr, err := sdk.AccAddressFromBech32(mintTo)
	if err != nil {
		return err
	}

	return k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		types.ModuleName,
		addr,
		sdk.NewCoins(amount),
	)
}
