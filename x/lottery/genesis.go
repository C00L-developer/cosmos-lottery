package lottery

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"lottery/x/lottery/keeper"
	"lottery/x/lottery/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the bet
	for _, elem := range genState.BetList {
		k.SetBet(ctx, elem)
	}
	// Set all the lottery
	for _, elem := range genState.LotteryList {
		k.SetLottery(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.BetList = k.GetAllBet(ctx)
	genesis.LotteryList = k.GetAllLottery(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
