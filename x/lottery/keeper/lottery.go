package keeper

import (
	"lottery/x/lottery/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetLottery set a specific lottery in the store from its index
func (k Keeper) SetLottery(ctx sdk.Context, lottery types.Lottery) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryKeyPrefix))
	b := k.cdc.MustMarshal(&lottery)
	store.Set(types.LotteryKey(
		lottery.Index,
	), b)
}

// GetLottery returns a lottery from its index
func (k Keeper) GetLottery(
	ctx sdk.Context,
	index uint64,
) (val types.Lottery, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryKeyPrefix))

	b := store.Get(types.LotteryKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLottery removes a lottery from the store
func (k Keeper) RemoveLottery(
	ctx sdk.Context,
	index uint64,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryKeyPrefix))
	store.Delete(types.LotteryKey(
		index,
	))
}

// GetAllLottery returns all lottery
func (k Keeper) GetAllLottery(ctx sdk.Context) (list []types.Lottery) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Lottery
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetOpenLottery returns the current open lottery
// if not exist, returns the created one
func (k Keeper) GetOpenLottery(ctx sdk.Context) types.Lottery {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryKeyPrefix))
	iterator := sdk.KVStoreReversePrefixIterator(store, []byte{})
	var val types.Lottery
	var lastLotteryID uint64
	if iterator.Valid() {
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Winner == "" || val.Winner == types.ModuleName {
			return val
		} else {
			lastLotteryID = val.Index
		}
	}
	newLottery := types.Lottery{
		Index:   lastLotteryID + 1,
		Creator: types.ModuleName,
	}
	k.SetLottery(ctx, newLottery)

	return newLottery
}
