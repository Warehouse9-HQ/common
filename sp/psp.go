package sp

// Permanent Storage Provider
type PSP struct {
	WalletAddress           string           `json:"wallet_address"` // Unique id which also can be used to verify digital signatures
	TokensLocked            bool             `json:"tokens_locked"`  // Are there any tokens locked?
	AddressesTokensLockedOn map[string]int64 // Map: Address -> amount of tokens locked

	IP          string `json:"ip"`           // IP address to reach the target
	VisibleName string `json:"visible_name"` // Visible name broadcasted to the network

	StorageAvailableMB int64 `json:"storage_available_mb"` // Amount of the storage available (in MB)
	PricePerMB         int64 `json:"price_per_mb"`         // Price per MB
}

func (psp *PSP) GetType() ProviderType {
	return PermanentStorageProvider
}

func (psp *PSP) GetAddress() string {
	return psp.WalletAddress
}

func (psp *PSP) GetTokensLocked() bool {
	return psp.TokensLocked
}

func (psp *PSP) GetListOfAddresses() []string {
	var result []string
	for k := range psp.AddressesTokensLockedOn {
		result = append(result, k)
	}
	return result
}

func (psp *PSP) GetAddressesTokensLockedOn() map[string]int64 {
	return psp.AddressesTokensLockedOn
}

func (psp *PSP) GetAmountOfTokensLocked() int64 {
	var result int64
	for _, k := range psp.AddressesTokensLockedOn {
		result += k
	}
	return result
}

func (psp *PSP) GetIP() string {
	return psp.IP
}

func (psp *PSP) GetVisibleName() string {
	return psp.VisibleName
}

func (psp *PSP) GetAvailableStorageMB() int64 {
	return psp.StorageAvailableMB
}

func (psp *PSP) GetPricePerMB() int64 {
	return psp.PricePerMB
}
