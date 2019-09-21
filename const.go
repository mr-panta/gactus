package gactus

// Environment core variables
const (
	CoreHTTPPortVar                   = "GACTUS_CORE_HTTP_PORT"
	CoreTCPPortVar                    = "GACTUS_CORE_TCP_PORT"
	CoreDefaultHealthCheckIntervalVar = "GACTUS_CORE_HEALTH_CHECK_INTERVAL"
)

// Environment service variables
const (
	ServiceCoreAddrVar = "GACTUS_SERVICE_CORE_ADDR"
	ServiceTCPPortVar  = "GACTUS_SERVICE_TCP_PORT"
)

// Default core addresses
const (
	DefaultCoreHTTPPort        = 8000
	DefaultCoreTCPPort         = 1739
	DefaultHealthCheckInterval = 300
)
