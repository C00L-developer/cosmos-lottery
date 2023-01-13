package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"lottery/x/lottery/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RevealBet(goCtx context.Context, msg *types.MsgRevealBet) (*types.MsgRevealBetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	lottery := k.GetOpenLottery(ctx)

	// reveal the existing bet
	betID := types.LotteryKey(lottery.Index)
	betID = append(betID, []byte(msg.GetCreator())...)
	betKey := string(betID)
	bet, found := k.GetBet(ctx, betKey)
	if !found {
		return nil, fmt.Errorf("the given bet not exist in the current lottery %d", lottery.Index)
	}

	suffixHash := sha256.Sum256([]byte(bet.Amount + msg.Suffix))
	if bet.SuffixHash != hex.EncodeToString(suffixHash[:]) {
		return nil, fmt.Errorf("the suffix hash %s is not matched with suffix %s", bet.SuffixHash, msg.Suffix)
	}
	bet.Suffix = msg.Suffix

	k.SetBet(ctx, bet)

	return &types.MsgRevealBetResponse{}, nil
}
