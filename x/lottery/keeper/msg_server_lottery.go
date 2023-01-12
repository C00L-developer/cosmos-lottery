package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"lottery/x/lottery/types"
)

func (k msgServer) CreateLottery(goCtx context.Context, msg *types.MsgCreateLottery) (*types.MsgCreateLotteryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetLottery(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var lottery = types.Lottery{
		Creator: msg.Creator,
		Index:   msg.Index,
		Winner:  msg.Winner,
	}

	k.SetLottery(
		ctx,
		lottery,
	)
	return &types.MsgCreateLotteryResponse{}, nil
}

func (k msgServer) UpdateLottery(goCtx context.Context, msg *types.MsgUpdateLottery) (*types.MsgUpdateLotteryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetLottery(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var lottery = types.Lottery{
		Creator: msg.Creator,
		Index:   msg.Index,
		Winner:  msg.Winner,
	}

	k.SetLottery(ctx, lottery)

	return &types.MsgUpdateLotteryResponse{}, nil
}

func (k msgServer) DeleteLottery(goCtx context.Context, msg *types.MsgDeleteLottery) (*types.MsgDeleteLotteryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetLottery(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveLottery(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteLotteryResponse{}, nil
}
