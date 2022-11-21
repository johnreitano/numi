package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/johnreitano/numi/x/numi/types"
)

// SetUserAccountAddress set a specific userAccountAddress in the store from its index
func (k Keeper) SetUserAccountAddress(ctx sdk.Context, userAccountAddress types.UserAccountAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserAccountAddressKeyPrefix))
	b := k.cdc.MustMarshal(&userAccountAddress)
	store.Set(types.UserAccountAddressKey(
		userAccountAddress.AccountAddress,
	), b)
}

// GetUserAccountAddress returns a userAccountAddress from its index
func (k Keeper) GetUserAccountAddress(
	ctx sdk.Context,
	accountAddress string,

) (val types.UserAccountAddress, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserAccountAddressKeyPrefix))

	b := store.Get(types.UserAccountAddressKey(
		accountAddress,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUserAccountAddress removes a userAccountAddress from the store
func (k Keeper) RemoveUserAccountAddress(
	ctx sdk.Context,
	accountAddress string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserAccountAddressKeyPrefix))
	store.Delete(types.UserAccountAddressKey(
		accountAddress,
	))
}

// GetAllUserAccountAddress returns all userAccountAddress
func (k Keeper) GetAllUserAccountAddress(ctx sdk.Context) (list []types.UserAccountAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserAccountAddressKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UserAccountAddress
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
