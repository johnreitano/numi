package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/johnreitano/numi/x/numi/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) UserAccountAddressAll(c context.Context, req *types.QueryAllUserAccountAddressRequest) (*types.QueryAllUserAccountAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var userAccountAddresss []types.UserAccountAddress
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	userAccountAddressStore := prefix.NewStore(store, types.KeyPrefix(types.UserAccountAddressKeyPrefix))

	pageRes, err := query.Paginate(userAccountAddressStore, req.Pagination, func(key []byte, value []byte) error {
		var userAccountAddress types.UserAccountAddress
		if err := k.cdc.Unmarshal(value, &userAccountAddress); err != nil {
			return err
		}

		userAccountAddresss = append(userAccountAddresss, userAccountAddress)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUserAccountAddressResponse{UserAccountAddress: userAccountAddresss, Pagination: pageRes}, nil
}

func (k Keeper) UserAccountAddress(c context.Context, req *types.QueryGetUserAccountAddressRequest) (*types.QueryGetUserAccountAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetUserAccountAddress(
		ctx,
		req.AccountAddress,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetUserAccountAddressResponse{UserAccountAddress: val}, nil
}
