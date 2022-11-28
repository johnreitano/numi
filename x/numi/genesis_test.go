package numi_test

import (
	"testing"

	keepertest "github.com/johnreitano/numi/testutil/keeper"
	"github.com/johnreitano/numi/testutil/nullify"
	"github.com/johnreitano/numi/x/numi"
	"github.com/johnreitano/numi/x/numi/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		UserList: []types.User{
			{
				UserId: "0",
			},
			{
				UserId: "1",
			},
		},
		UserAccountAddressList: []types.UserAccountAddress{
			{
				AccountAddress: "0",
			},
			{
				AccountAddress: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NumiKeeperWithMocks(t)
	numi.InitGenesis(ctx, *k, genesisState)
	got := numi.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.UserList, got.UserList)
	require.ElementsMatch(t, genesisState.UserAccountAddressList, got.UserAccountAddressList)
	// this line is used by starport scaffolding # genesis/test/assert
}
