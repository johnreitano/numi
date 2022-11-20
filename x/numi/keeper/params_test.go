package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/johnreitano/numi/app"
	testkeeper "github.com/johnreitano/numi/testutil/keeper"
	"github.com/johnreitano/numi/x/numi/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NumiKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.IdentityVerifiers, k.IdentityVerifiers(ctx))
}

func setSDKPrefixes() {
	accountPubKeyPrefix := app.AccountAddressPrefix + "pub"
	validatorAddressPrefix := app.AccountAddressPrefix + "valoper"
	validatorPubKeyPrefix := app.AccountAddressPrefix + "valoperpub"
	consNodeAddressPrefix := app.AccountAddressPrefix + "valcons"
	consNodePubKeyPrefix := app.AccountAddressPrefix + "valconspub"

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(app.AccountAddressPrefix, accountPubKeyPrefix)
	config.SetBech32PrefixForValidator(validatorAddressPrefix, validatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(consNodeAddressPrefix, consNodePubKeyPrefix)
	config.Seal()
}

func TestIdentityVerifierAddresses(t *testing.T) {
	setSDKPrefixes()

	k, ctx := testkeeper.NumiKeeper(t)
	verifier0 := "numi17jmfn9c6x7k0uem9hndf9808u0ufx24zjlqyke"
	verifier1 := "numi1v030en3xa6azvyc477mhs3jh46xg9hcjg6x333"
	verifiers := fmt.Sprintf("%s,%s", verifier0, verifier1)
	k.SetParams(ctx, types.NewParams(verifiers))

	addrs := k.IdentityVerifierAddresses(ctx)
	require.EqualValues(t, len(addrs), 2)
	require.EqualValues(t, addrs[0].String(), verifier0)
	require.EqualValues(t, addrs[1].String(), verifier1)
}

func TestIsIdentityVerifier(t *testing.T) {
	setSDKPrefixes()

	k, ctx := testkeeper.NumiKeeper(t)
	verifier0 := "numi17jmfn9c6x7k0uem9hndf9808u0ufx24zjlqyke"
	verifier1 := "numi1v030en3xa6azvyc477mhs3jh46xg9hcjg6x333"
	verifiers := fmt.Sprintf("%s,%s", verifier0, verifier1)
	k.SetParams(ctx, types.NewParams(verifiers))

	addr0 := sdk.MustAccAddressFromBech32(verifier0)
	require.True(t, k.IsIdentityVerifier(ctx, addr0))

	addr1 := sdk.MustAccAddressFromBech32(verifier1)
	require.True(t, k.IsIdentityVerifier(ctx, addr1))

	verifier2 := "numi1hlmu6pw6ff9tqx6zplzrhsszv6xh7c8stj3w6k"
	addr2 := sdk.MustAccAddressFromBech32(verifier2)
	require.False(t, k.IsIdentityVerifier(ctx, addr2))
}
