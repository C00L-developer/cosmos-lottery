package keeper

import (
	"context"
	"fmt"

	"lottery/x/lottery/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddBet(goCtx context.Context, msg *types.MsgAddBet) (*types.MsgAddBetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	lottery := k.GetOpenLottery(ctx)
	if lottery.Winner != "" {
		return nil, fmt.Errorf("the lottery %d is locked, cannot bet at this moment", lottery.Index)
	}
	// create a new bet
	betID := types.LotteryKey(lottery.Index)
	betID = append(betID, []byte(msg.GetCreator())...)
	betKey := string(betID)
	bet := types.Bet{
		Index:      betKey,
		Amount:     msg.Amount,
		SuffixHash: msg.GetSuffixHash(),
		Player:     msg.GetCreator(),
	}

	// send tokens from the player to the module account
	player, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	amount, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return nil, err
	}
	// check if the bet amount is valid
	params := k.GetParams(ctx)
	fee, _ := sdk.ParseCoinsNormalized(params.LotteryFee)
	minBet, _ := sdk.ParseCoinsNormalized(params.MinBetAmount)
	maxBet, _ := sdk.ParseCoinsNormalized(params.MaxBetAmount)
	if !amount.IsAllGTE(fee.Add(minBet...)) {
		return nil, fmt.Errorf("the bet amount %v is not enough", amount)
	}
	if amount.IsAllGT(fee.Add(maxBet...)) {
		return nil, fmt.Errorf("the bet amount %v is too large", amount)
	}
	// send coins to the lottery pool
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, player, types.ModuleName, amount); err != nil {
		return nil, err
	}

	k.SetBet(ctx, bet)

	return &types.MsgAddBetResponse{}, nil
}
