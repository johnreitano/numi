package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/johnreitano/numi/x/numi/types"
)

func (k msgServer) CreateAndVerifyUser(goCtx context.Context, msg *types.MsgCreateAndVerifyUser) (*types.MsgCreateAndVerifyUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateAndVerifyUserResponse{}, nil
}
