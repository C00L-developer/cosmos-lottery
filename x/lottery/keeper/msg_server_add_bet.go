package keeper

import (
	"context"

	"lottery/x/lottery/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddBet(goCtx context.Context, msg *types.MsgAddBet) (*types.MsgAddBetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	lottery := k.GetOpenLottery(ctx)

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
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, player, types.ModuleName, amount); err != nil {
		return nil, err
	}

	k.SetBet(ctx, bet)

	return &types.MsgAddBetResponse{}, nil
}
