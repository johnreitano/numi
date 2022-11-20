package keeper

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/johnreitano/numi/x/numi/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.IdentityVerifiers(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// IdentityVerifiers returns the IdentityVerifiers param
func (k Keeper) IdentityVerifiers(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyIdentityVerifiers, &res)
	return
}

// IdentityVerifierAddresses returns the bech32 addresses associated with the IdentityVerifiers param
func (k Keeper) IdentityVerifierAddresses(ctx sdk.Context) []sdk.AccAddress {
	verifiers := strings.Split(k.IdentityVerifiers(ctx), ",")
	verifierAddrs := []sdk.AccAddress{}
	for _, verifier := range verifiers {
		verifierAddr, err := sdk.AccAddressFromBech32(verifier)
		if err != nil {
			k.Logger(ctx).Info(fmt.Sprintf("identity verifier address %s is invalid - skipping: %s", verifier, err))
			continue
		}
		verifierAddrs = append(verifierAddrs, verifierAddr)
	}

	return verifierAddrs
}

// IsIdentityVerifier returns true iff the subject address matches on of the items in the IdentityVerifiers param
func (k Keeper) IsIdentityVerifier(ctx sdk.Context, subject sdk.AccAddress) bool {
	for _, verifierAddr := range k.IdentityVerifierAddresses(ctx) {
		if verifierAddr.Equals(subject) {
			return true
		}
	}
	return false
}
