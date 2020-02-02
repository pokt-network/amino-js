package types

import (
	"sync"
)

// TmConfig is the structure that holds the SDK configuration parameters.
// This could be used to initialize certain configuration parameters for the SDK.
type Config struct {
	mtx             sync.RWMutex
	sealed          bool
	addressVerifier func([]byte) error
}

var (
	// Initializing an instance of TmConfig
	sdkConfig = &Config{
		sealed: false,
	}
)

// GetConfig returns the config instance for the SDK.
func GetConfig() *Config {
	return sdkConfig
}

// GetAddressVerifier returns the function to verify that addresses have the correct format
func (config *Config) GetAddressVerifier() func([]byte) error {
	return config.addressVerifier
}
