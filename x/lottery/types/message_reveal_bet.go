package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRevealBet = "reveal_bet"

var _ sdk.Msg = &MsgRevealBet{}

func NewMsgRevealBet(creator string, suffix string) *MsgRevealBet {
	return &MsgRevealBet{
		Creator: creator,
		Suffix:  suffix,
	}
}

func (msg *MsgRevealBet) Route() string {
	return RouterKey
}

func (msg *MsgRevealBet) Type() string {
	return TypeMsgRevealBet
}

func (msg *MsgRevealBet) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRevealBet) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRevealBet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
