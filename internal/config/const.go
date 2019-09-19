package config

// Core server commands
const (
	CMDCoreRegisterProcessors = "__core__.register_processors"
)

// Service server commands
const (
	CMDServiceUpdateRegistries = "__service__.update_registries"
	CMDServiceHealthCheck      = "__service__.health_check"
)

// Error
const (
	ErrorServiceNotAvailable = "the internal service is not available"
	ErrorNotFound            = "not found"
)
