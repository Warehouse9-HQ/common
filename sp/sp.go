package sp

type SP interface {
	GetType() ProviderType // Returns provider type

	GetAddress() string    // Returns address of SP
	GetTokensLocked() bool // Returns TokensLocked

	GetListOfAddresses() []string                 // Returns list of addresses tokens locked on
	GetAddressesTokensLockedOn() map[string]int64 // Returns map: Address -> amount of tokens
	GetAmountOfTokensLocked() int64               // Returns total amount of tokens locked

	GetIP() string          // Returns IP address of a node
	GetVisibleName() string // Returns visible name

	GetAvailableStorageMB() int64 // Returns available storage
	GetPricePerMB() int64         // Returns price per MB
}
