package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"lottery/x/lottery/types"
)

func (k msgServer) AddBet(goCtx context.Context, msg *types.MsgAddBet) (*types.MsgAddBetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAddBetResponse{}, nil
}
