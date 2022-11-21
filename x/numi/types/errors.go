package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/numi module sentinel errors
var (
	ErrUserIdNotUUID                          = sdkerrors.Register(ModuleName, 1100, "userId is not a UUID")
	ErrFirstNameBlank                         = sdkerrors.Register(ModuleName, 1101, "first name is blank")
	ErrLastNameBlank                          = sdkerrors.Register(ModuleName, 1102, "last name is blank")
	ErrCountryCodeBlank                       = sdkerrors.Register(ModuleName, 1103, "country code is blank")
	ErrCountryCodeInvalid                     = sdkerrors.Register(ModuleName, 1104, "country code is blank")
	ErrSubnationalEntityBlank                 = sdkerrors.Register(ModuleName, 1105, "subnational entity is blank")
	ErrCityBlank                              = sdkerrors.Register(ModuleName, 1106, "city is blank")
	ErrBioBlank                               = sdkerrors.Register(ModuleName, 1107, "bio is blank")
	ErrCreatorNotAuthorizedToVerifyIdentities = sdkerrors.Register(ModuleName, 1108, "creator is not authorized to verify identities")
	ErrUserIdAlreadyExists                    = sdkerrors.Register(ModuleName, 1109, "user with that userId already exists")
	ErrAccountAddressAlreadyExists            = sdkerrors.Register(ModuleName, 1110, "user with that account address already exists")
)
