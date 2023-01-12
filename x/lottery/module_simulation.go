package lottery

import (
	"math/rand"

	"lottery/testutil/sample"
	lotterysimulation "lottery/x/lottery/simulation"
	"lottery/x/lottery/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = lotterysimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgAddBet = "op_weight_msg_add_bet"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddBet int = 100

	opWeightMsgRevealBet = "op_weight_msg_reveal_bet"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRevealBet int = 100

	opWeightMsgCreateLottery = "op_weight_msg_lottery"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateLottery int = 100

	opWeightMsgUpdateLottery = "op_weight_msg_lottery"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateLottery int = 100

	opWeightMsgDeleteLottery = "op_weight_msg_lottery"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteLottery int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	lotteryGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		LotteryList: []types.Lottery{
			{
				Creator: sample.AccAddress(),
				Index:   0,
			},
			{
				Creator: sample.AccAddress(),
				Index:   1,
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&lotteryGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgAddBet int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddBet, &weightMsgAddBet, nil,
		func(_ *rand.Rand) {
			weightMsgAddBet = defaultWeightMsgAddBet
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddBet,
		lotterysimulation.SimulateMsgAddBet(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRevealBet int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRevealBet, &weightMsgRevealBet, nil,
		func(_ *rand.Rand) {
			weightMsgRevealBet = defaultWeightMsgRevealBet
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRevealBet,
		lotterysimulation.SimulateMsgRevealBet(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateLottery int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateLottery, &weightMsgCreateLottery, nil,
		func(_ *rand.Rand) {
			weightMsgCreateLottery = defaultWeightMsgCreateLottery
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateLottery,
		lotterysimulation.SimulateMsgCreateLottery(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateLottery int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateLottery, &weightMsgUpdateLottery, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateLottery = defaultWeightMsgUpdateLottery
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateLottery,
		lotterysimulation.SimulateMsgUpdateLottery(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteLottery int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteLottery, &weightMsgDeleteLottery, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteLottery = defaultWeightMsgDeleteLottery
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteLottery,
		lotterysimulation.SimulateMsgDeleteLottery(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
