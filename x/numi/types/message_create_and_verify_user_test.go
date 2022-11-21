package types_test

import (
	"testing"

	"github.com/johnreitano/numi/app"
	typestest "github.com/johnreitano/numi/testutil/types"
	"github.com/johnreitano/numi/x/numi/types"
	"github.com/stretchr/testify/require"
)

func TestValidateBasic(t *testing.T) {
	app.SetAddressPrefixesInSDKConfig()

	validMsg := typestest.ValidMsgCreateAndVerifyUser()
	for _, tc := range []struct {
		desc  string
		msg   func() *types.MsgCreateAndVerifyUser
		valid bool
	}{
		{
			desc:  "user with all fields valid",
			msg:   func() *types.MsgCreateAndVerifyUser { return validMsg },
			valid: true,
		},
		{
			desc: "missing creator",
			msg: func() *types.MsgCreateAndVerifyUser {
				msg := *validMsg
				msg.Creator = ""
				return &msg
			},
			valid: false,
		},
		{
			desc: "invalid creator",
			msg: func() *types.MsgCreateAndVerifyUser {
				user := *validMsg
				user.Creator = "xxx"
				return &user
			},
			valid: false,
		},
		{
			desc: "missing user id",
			msg: func() *types.MsgCreateAndVerifyUser {
				user := *validMsg
				user.UserId = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "invalid user id",
			msg: func() *types.MsgCreateAndVerifyUser {
				user := *validMsg
				user.UserId = "xxx"
				return &user
			},
			valid: false,
		},
		{
			desc: "missing first name",
			msg: func() *types.MsgCreateAndVerifyUser {
				user := *validMsg
				user.FirstName = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "missing last name",
			msg: func() *types.MsgCreateAndVerifyUser {
				user := *validMsg
				user.LastName = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "missing country code",
			msg: func() *types.MsgCreateAndVerifyUser {
				user := *validMsg
				user.CountryCode = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "2-digit country code",
			msg: func() *types.MsgCreateAndVerifyUser {
				user := *validMsg
				user.CountryCode = "US"
				return &user
			},
			valid: false,
		},
		{
			desc: "invalid 3-digit country code",
			msg: func() *types.MsgCreateAndVerifyUser {
				user := *validMsg
				user.CountryCode = "XXX"
				return &user
			},
			valid: false,
		},
		{
			desc: "missing subnational entity",
			msg: func() *types.MsgCreateAndVerifyUser {
				user := *validMsg
				user.SubnationalEntity = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "missing city",
			msg: func() *types.MsgCreateAndVerifyUser {
				user := *validMsg
				user.City = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "missing bio",
			msg: func() *types.MsgCreateAndVerifyUser {
				user := *validMsg
				user.Bio = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "missing referrer",
			msg: func() *types.MsgCreateAndVerifyUser {
				user := *validMsg
				user.Referrer = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "invalid referrer",
			msg: func() *types.MsgCreateAndVerifyUser {
				user := *validMsg
				user.Referrer = "xxx"
				return &user
			},
			valid: false,
		},
		{
			desc: "missing account address",
			msg: func() *types.MsgCreateAndVerifyUser {
				user := *validMsg
				user.AccountAddress = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "invalid account address",
			msg: func() *types.MsgCreateAndVerifyUser {
				user := *validMsg
				user.AccountAddress = "xxx"
				return &user
			},
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := types.ValidateUserBasic(tc.msg())
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
