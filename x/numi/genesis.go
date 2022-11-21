package numi

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/johnreitano/numi/x/numi/keeper"
	"github.com/johnreitano/numi/x/numi/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the user
	for _, elem := range genState.UserList {
		k.SetUser(ctx, elem)
	}
	// Set all the userAccountAddress
	for _, elem := range genState.UserAccountAddressList {
		k.SetUserAccountAddress(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.UserList = k.GetAllUser(ctx)
	genesis.UserAccountAddressList = k.GetAllUserAccountAddress(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
