package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgCreateAndVerifyUser = "create_and_verify_user"

var _ sdk.Msg = &MsgCreateAndVerifyUser{}

func NewMsgCreateAndVerifyUser(creator string, userId string, firstName string, lastName string, countryCode string, subnationalEntity string, city string, bio string, referrer string, accountAddress string) *MsgCreateAndVerifyUser {
	return &MsgCreateAndVerifyUser{
		Creator:           creator,
		UserId:            userId,
		FirstName:         firstName,
		LastName:          lastName,
		CountryCode:       countryCode,
		SubnationalEntity: subnationalEntity,
		City:              city,
		Bio:               bio,
		Referrer:          referrer,
		AccountAddress:    accountAddress,
	}
}

func (msg *MsgCreateAndVerifyUser) Route() string {
	return RouterKey
}

func (msg *MsgCreateAndVerifyUser) Type() string {
	return TypeMsgCreateAndVerifyUser
}

func (msg *MsgCreateAndVerifyUser) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAndVerifyUser) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAndVerifyUser) ValidateBasic() error {
	return ValidateUserBasic(msg)
}
