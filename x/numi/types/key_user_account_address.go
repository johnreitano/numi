package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// UserAccountAddressKeyPrefix is the prefix to retrieve all UserAccountAddress
	UserAccountAddressKeyPrefix = "UserAccountAddress/value/"
)

// UserAccountAddressKey returns the store key to retrieve a UserAccountAddress from the index fields
func UserAccountAddressKey(
	accountAddress string,
) []byte {
	var key []byte

	accountAddressBytes := []byte(accountAddress)
	key = append(key, accountAddressBytes...)
	key = append(key, []byte("/")...)

	return key
}
