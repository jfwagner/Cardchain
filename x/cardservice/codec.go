package cardservice

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSetName{}, "cardservice/SetName", nil)
	cdc.RegisterConcrete(MsgBuyCardScheme{}, "cardservice/BuyCardScheme", nil)
	cdc.RegisterConcrete(MsgSaveCardContent{}, "cardservice/SaveCardContent", nil)
	cdc.RegisterConcrete(MsgVoteCard{}, "cardservice/VoteCard", nil)
	cdc.RegisterConcrete(Card{}, "cardservice/Card", nil)
}