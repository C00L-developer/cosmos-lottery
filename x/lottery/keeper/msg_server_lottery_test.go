package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "lottery/testutil/keeper"
	"lottery/x/lottery/keeper"
	"lottery/x/lottery/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestLotteryMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.LotteryKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateLottery{
			Creator: creator,
			Index:   uint64(i),
		}
		_, err := srv.CreateLottery(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetLottery(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestLotteryMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateLottery
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateLottery{
				Creator: creator,
				Index:   uint64(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateLottery{
				Creator: "B",
				Index:   uint64(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateLottery{
				Creator: creator,
				Index:   uint64(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.LotteryKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateLottery{
				Creator: creator,
				Index:   uint64(0),
			}
			_, err := srv.CreateLottery(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateLottery(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetLottery(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestLotteryMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteLottery
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteLottery{
				Creator: creator,
				Index:   uint64(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteLottery{
				Creator: "B",
				Index:   uint64(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteLottery{
				Creator: creator,
				Index:   uint64(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.LotteryKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateLottery(wctx, &types.MsgCreateLottery{
				Creator: creator,
				Index:   uint64(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteLottery(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetLottery(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
