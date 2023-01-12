package lottery_test

import (
	"testing"

	keepertest "lottery/testutil/keeper"
	"lottery/testutil/nullify"
	"lottery/x/lottery"
	"lottery/x/lottery/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		BetList: []types.Bet{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		LotteryList: []types.Lottery{
			{
				Index: 0,
			},
			{
				Index: 1,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LotteryKeeper(t)
	lottery.InitGenesis(ctx, *k, genesisState)
	got := lottery.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.BetList, got.BetList)
	require.ElementsMatch(t, genesisState.LotteryList, got.LotteryList)
	// this line is used by starport scaffolding # genesis/test/assert
}
