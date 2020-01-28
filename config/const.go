package config

import (
	"time"
)

// Core server commands
const (
	CMDCoreRegisterService = "__core__.register_service"
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

// Server Configuration
const (
	LimitSize = 128
)

const ReceiverMethodName = "Service.Receive"

const (
	MinConns        = 10
	MaxConns        = 1000
	IdleConnTimeout = 10 * time.Millisecond
	WaitConnTimeout = 100 * time.Millisecond
	ClearPeriod     = time.Second
)
