package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		BetList:     []Bet{},
		LotteryList: []Lottery{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in bet
	betIndexMap := make(map[string]struct{})

	for _, elem := range gs.BetList {
		index := string(BetKey(elem.Index))
		if _, ok := betIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for bet")
		}
		betIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in lottery
	lotteryIndexMap := make(map[string]struct{})

	for _, elem := range gs.LotteryList {
		index := string(LotteryKey(elem.Index))
		if _, ok := lotteryIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for lottery")
		}
		lotteryIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
