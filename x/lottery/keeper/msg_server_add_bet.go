package keeper

import (
	"context"

	"lottery/x/lottery/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddBet(goCtx context.Context, msg *types.MsgAddBet) (*types.MsgAddBetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// k.GetLottery()
	// create a new bet
	// _ := types.Bet{}
	_ = ctx

	return &types.MsgAddBetResponse{}, nil
}
