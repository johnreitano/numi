package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func SetAddressPrefixesInSDKConfig() *sdk.Config {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(AccountAddressPrefix, AccountAddressPrefix+"pub")
	config.SetBech32PrefixForValidator(AccountAddressPrefix+"valoper", AccountAddressPrefix+"valoperpub")
	config.SetBech32PrefixForConsensusNode(AccountAddressPrefix+"valcons", AccountAddressPrefix+"valconspub")
	return config
}
