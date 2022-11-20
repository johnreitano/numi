package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyIdentityVerifiers            = []byte("IdentityVerifiers")
	DefaultIdentityVerifiers string = ""
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	identityVerifiers string,
) Params {
	return Params{
		IdentityVerifiers: identityVerifiers,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultIdentityVerifiers,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyIdentityVerifiers, &p.IdentityVerifiers, validateIdentityVerifiers),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateIdentityVerifiers(p.IdentityVerifiers); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateIdentityVerifiers validates the IdentityVerifiers param
func validateIdentityVerifiers(v interface{}) error {
	identityVerifiers, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}
	for _, verifier := range strings.Split(identityVerifiers, ",") {
		if verifier == "" {
			continue
		}
		_, err := sdk.AccAddressFromBech32(verifier)
		if err != nil {
			return fmt.Errorf("identity verifier address %s is invalid: %w", identityVerifiers, err)
		}
	}
	return nil
}
