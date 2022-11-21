package types_test

import (
	"testing"

	"github.com/johnreitano/numi/app"
	"github.com/johnreitano/numi/x/numi/types"
	"github.com/stretchr/testify/require"
)

func TestValidatableAsUser_ValidateUserBasic(t *testing.T) {
	app.SetAddressPrefixesInSDKConfig()

	validUser := types.User{
		Creator:           "numi17jmfn9c6x7k0uem9hndf9808u0ufx24zjlqyke",
		UserId:            "1bc3e020-2b02-40a7-abd8-eadc9b4250c5",
		FirstName:         "John",
		LastName:          "Doe",
		CountryCode:       "USA",
		SubnationalEntity: "California",
		City:              "San Diego",
		Bio:               "a serious man",
		Referrer:          "numi1staes5penzmhsxk3vmh4uwf483yejnvvx2ljwe",
		AccountAddress:    "numi1tsacr4aqrrerakdlcmlzl7daplle54fj874w2s",
	}
	for _, tc := range []struct {
		desc  string
		user  func() *types.User
		valid bool
	}{
		{
			desc:  "user with all fields valid",
			user:  func() *types.User { return &validUser },
			valid: true,
		},
		{
			desc: "missing creator",
			user: func() *types.User {
				user := validUser
				user.Creator = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "invalid creator",
			user: func() *types.User {
				user := validUser
				user.Creator = "xxx"
				return &user
			},
			valid: false,
		},
		{
			desc: "missing user id",
			user: func() *types.User {
				user := validUser
				user.UserId = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "invalid user id",
			user: func() *types.User {
				user := validUser
				user.UserId = "xxx"
				return &user
			},
			valid: false,
		},
		{
			desc: "missing first name",
			user: func() *types.User {
				user := validUser
				user.FirstName = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "missing last name",
			user: func() *types.User {
				user := validUser
				user.LastName = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "missing country code",
			user: func() *types.User {
				user := validUser
				user.CountryCode = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "2-digit country code",
			user: func() *types.User {
				user := validUser
				user.CountryCode = "US"
				return &user
			},
			valid: false,
		},
		{
			desc: "invalid 3-digit country code",
			user: func() *types.User {
				user := validUser
				user.CountryCode = "XXX"
				return &user
			},
			valid: false,
		},
		{
			desc: "missing subnational entity",
			user: func() *types.User {
				user := validUser
				user.SubnationalEntity = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "missing city",
			user: func() *types.User {
				user := validUser
				user.City = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "missing bio",
			user: func() *types.User {
				user := validUser
				user.Bio = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "missing referrer",
			user: func() *types.User {
				user := validUser
				user.Referrer = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "invalid referrer",
			user: func() *types.User {
				user := validUser
				user.Referrer = "xxx"
				return &user
			},
			valid: false,
		},
		{
			desc: "missing account address",
			user: func() *types.User {
				user := validUser
				user.AccountAddress = ""
				return &user
			},
			valid: false,
		},
		{
			desc: "invalid account address",
			user: func() *types.User {
				user := validUser
				user.AccountAddress = "xxx"
				return &user
			},
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := types.ValidateUserBasic(tc.user())
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
