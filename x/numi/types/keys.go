package types

const (
	// ModuleName defines the module name
	ModuleName = "numi"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_numi"
)

const (
	UserCreatedAndVerifiedEventType              = "user-created-and-verified"
	UserCreatedAndVerifiedEventUserId            = "user-id"
	UserCreatedAndVerifiedEventFirstName         = "first-name"
	UserCreatedAndVerifiedEventLastName          = "last-name"
	UserCreatedAndVerifiedEventCountryCode       = "country-code"
	UserCreatedAndVerifiedEventSubnationalEntity = "subnational-entity"
	UserCreatedAndVerifiedEventCity              = "city"
	UserCreatedAndVerifiedEventBio               = "bio"
	UserCreatedAndVerifiedEventCreator           = "creator"
	UserCreatedAndVerifiedEventReferrer          = "referrer"
	UserCreatedAndVerifiedEventAccountAddress    = "account-address"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
