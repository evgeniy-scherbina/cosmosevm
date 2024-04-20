package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEthereumTx = "ethereum_tx"

var _ sdk.Msg = &MsgEthereumTx{}

func NewMsgEthereumTx(data *codectypes.Any, size float64, hash string, from string) *MsgEthereumTx {
	return &MsgEthereumTx{
		Data:  data,
		Size_: size,
		Hash:  hash,
		From:  from,
	}
}

func (msg *MsgEthereumTx) Route() string {
	return RouterKey
}

func (msg *MsgEthereumTx) Type() string {
	return TypeMsgEthereumTx
}

func (msg *MsgEthereumTx) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEthereumTx) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEthereumTx) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
