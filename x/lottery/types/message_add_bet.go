package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddBet = "add_bet"

var _ sdk.Msg = &MsgAddBet{}

func NewMsgAddBet(creator string, amount string, suffixHash string) *MsgAddBet {
	return &MsgAddBet{
		Creator:    creator,
		Amount:     amount,
		SuffixHash: suffixHash,
	}
}

func (msg *MsgAddBet) Route() string {
	return RouterKey
}

func (msg *MsgAddBet) Type() string {
	return TypeMsgAddBet
}

func (msg *MsgAddBet) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddBet) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddBet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
