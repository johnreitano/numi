package numi

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/johnreitano/numi/testutil/sample"
	numisimulation "github.com/johnreitano/numi/x/numi/simulation"
	"github.com/johnreitano/numi/x/numi/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = numisimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateAndVerifyUser = "op_weight_msg_create_and_verify_user"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAndVerifyUser int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	numiGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&numiGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	numiParams := types.DefaultParams()
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyIdentityVerifiers), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(numiParams.IdentityVerifiers))
		}),
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateAndVerifyUser int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateAndVerifyUser, &weightMsgCreateAndVerifyUser, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAndVerifyUser = defaultWeightMsgCreateAndVerifyUser
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAndVerifyUser,
		numisimulation.SimulateMsgCreateAndVerifyUser(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
