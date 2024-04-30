package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePeople{}

func NewMsgCreatePeople(creator string, address string, name string) *MsgCreatePeople {
	return &MsgCreatePeople{
		Creator: creator,
		Address: address,
		Name:    name,
	}
}

func (msg *MsgCreatePeople) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdatePeople{}

func NewMsgUpdatePeople(creator string, id uint64, address string, name string) *MsgUpdatePeople {
	return &MsgUpdatePeople{
		Id:      id,
		Creator: creator,
		Address: address,
		Name:    name,
	}
}

func (msg *MsgUpdatePeople) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeletePeople{}

func NewMsgDeletePeople(creator string, id uint64) *MsgDeletePeople {
	return &MsgDeletePeople{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeletePeople) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
