package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"lottery/x/lottery/types"
)

func (k msgServer) RevealBet(goCtx context.Context, msg *types.MsgRevealBet) (*types.MsgRevealBetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRevealBetResponse{}, nil
}
