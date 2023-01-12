package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgAddBet{}, "lottery/AddBet", nil)
	cdc.RegisterConcrete(&MsgRevealBet{}, "lottery/RevealBet", nil)
	cdc.RegisterConcrete(&MsgCreateLottery{}, "lottery/CreateLottery", nil)
	cdc.RegisterConcrete(&MsgUpdateLottery{}, "lottery/UpdateLottery", nil)
	cdc.RegisterConcrete(&MsgDeleteLottery{}, "lottery/DeleteLottery", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddBet{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRevealBet{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateLottery{},
		&MsgUpdateLottery{},
		&MsgDeleteLottery{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
