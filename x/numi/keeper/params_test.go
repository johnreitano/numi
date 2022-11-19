package keeper_test

import (
	"testing"

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
