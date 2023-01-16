package keeper

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"

	"lottery/x/lottery/types"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		bankKeeper types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

	bankKeeper types.BankKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		bankKeeper: bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Drawlots determine the winner for the open lottery.
func (k Keeper) Drawlots(ctx sdk.Context) {
	lottery := k.GetOpenLottery(ctx)
	bets := k.GetBets(ctx, lottery.Index)
	if lottery.Winner == "" {
		// check if be able to close the current lottery
		if len(bets) >= int(k.GetParams(ctx).BetThresCount) {
			lottery.Winner = lottery.Creator
			k.SetLottery(ctx, lottery)
		}
	} else {
		// check if all bets are revealed
		suffix := ""
		hi, lo := 0, 0
		rewards := sdk.Coins{}
		for i, bet := range bets {
			if bet.Suffix == "" {
				return
			}
			suffix += bet.Suffix
			amount, err := sdk.ParseCoinsNormalized(bet.Amount)
			if err != nil {
				panic(err)
			}
			rewards = rewards.Add(amount...)
			hiAmount, _ := sdk.ParseCoinsNormalized(bets[hi].Amount)
			loAmount, _ := sdk.ParseCoinsNormalized(bets[lo].Amount)
			if amount.IsAllGT(hiAmount) {
				hi = i
			}
			if amount.IsAllLT(loAmount) {
				lo = i
			}
		}
		suffixHash := sha256.Sum256([]byte(fmt.Sprintf("%s%d", suffix, lottery.Index)))
		winner := int(binary.BigEndian.Uint32(suffixHash[:4]) % uint32(len(bets)))
		lottery.Winner = bets[winner].Player

		if winner != lo {
			if winner != hi {
				fee, err := sdk.ParseCoinsNormalized(k.GetParams(ctx).LotteryFee)
				if err != nil {
					panic(err)
				}
				rewards = rewards.Sub(fee.MulInt(math.NewInt(int64(len(bets))))...)
			}
			player, _ := sdk.AccAddressFromBech32(lottery.Winner)
			k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, player, rewards)
		}
		k.SetLottery(ctx, lottery)
	}
}
