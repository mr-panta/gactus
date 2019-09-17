package gactus

// Environment core variables
const (
	CoreHTTPAddrVar = "GACTUS_CORE_HTTP_ADDR"
	CoreTCPAddrVar  = "GACTUS_CORE_TCP_ADDR"
)

// Environment service variables
const (
	ServiceCoreAddrVar = "GACTUS_SERVICE_CORE_ADDR"
	ServiceTCPAddrVar  = "GACTUS_SERVICE_TCP_ADDR"
)

// Default core addresses
const (
	DefaultCoreHTTPAddr = ":8000"
	DefaultCoreTCPAddr  = ":8001"
)

// Default command
const (
	CMDRegisterProcessors = "core.register_processors"
)
