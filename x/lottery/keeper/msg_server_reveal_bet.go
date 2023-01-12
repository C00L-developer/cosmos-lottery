package keeper

import (
	"context"

	"lottery/x/lottery/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RevealBet(goCtx context.Context, msg *types.MsgRevealBet) (*types.MsgRevealBetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRevealBetResponse{}, nil
}
