package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgDeposit{}, "jinx/MsgDeposit", nil)
	cdc.RegisterConcrete(&MsgWithdraw{}, "jinx/MsgWithdraw", nil)
	cdc.RegisterConcrete(&MsgBorrow{}, "jinx/MsgBorrow", nil)
	cdc.RegisterConcrete(&MsgLiquidate{}, "jinx/MsgLiquidate", nil)
	cdc.RegisterConcrete(&MsgRepay{}, "jinx/MsgRepay", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeposit{},
		&MsgWithdraw{},
		&MsgBorrow{},
		&MsgLiquidate{},
		&MsgRepay{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)

	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
}
