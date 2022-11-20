package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// UserKeyPrefix is the prefix to retrieve all User
	UserKeyPrefix = "User/value/"
)

// UserKey returns the store key to retrieve a User from the index fields
func UserKey(
	userId string,
) []byte {
	var key []byte

	userIdBytes := []byte(userId)
	key = append(key, userIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
