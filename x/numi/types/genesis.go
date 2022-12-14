package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		UserList:               []User{},
		UserAccountAddressList: []UserAccountAddress{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in user
	userIndexMap := make(map[string]struct{})

	for _, elem := range gs.UserList {
		index := string(UserKey(elem.UserId))
		if _, ok := userIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for user")
		}
		userIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in userAccountAddress
	userAccountAddressIndexMap := make(map[string]struct{})

	for _, elem := range gs.UserAccountAddressList {
		index := string(UserAccountAddressKey(elem.AccountAddress))
		if _, ok := userAccountAddressIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for userAccountAddress")
		}
		userAccountAddressIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
