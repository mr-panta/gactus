package gactus

// Environment core variables
const (
	CoreHTTPPortVar                   = "GACTUS_CORE_HTTP_PORT"
	CoreTCPPortVar                    = "GACTUS_CORE_TCP_PORT"
	CoreDefaultHealthCheckIntervalVar = "GACTUS_CORE_HEALTH_CHECK_INTERVAL"
)

// Environment service variables
const (
	ServiceNameVar            = "GACTUS_SERVICE_NAME"
	ServiceCoreAddrVar        = "GACTUS_SERVICE_CORE_ADDR"
	ServiceTCPPortVar         = "GACTUS_SERVICE_TCP_PORT"
	ServiceMinConnsVar        = "GACTUS_SERVICE_MIN_CONNS"
	ServiceMaxConnsVar        = "GACTUS_SERVICE_MAX_CONNS"
	ServiceIdleConnTimeoutVar = "GACTUS_SERVICE_IDLE_CONN_TIMEOUT"
	ServiceWaitConnTimeoutVar = "GACTUS_SERVICE_WAIT_CONN_TIMEOUT"
	ServiceClearPeriodVar     = "GACTUS_SERVICE_CLEAR_PERIOD"
	ServiceReadTimeoutVar     = "GACTUS_SERVICE_READ_TIMEOUT"
)

// Default core variables
const (
	DefaultCoreHTTPPort        = 8000
	DefaultCoreTCPPort         = 1739
	DefaultHealthCheckInterval = 300
)

// Default service variables
const (
	DefaultServiceTCPPort         = 3000
	DefaultServiceMinConns        = 10
	DefaultServiceMaxConns        = 100
	DefaultServiceIdleConnTimeout = 100
	DefaultServiceWaitConnTimeout = 10
	DefaultServiceClearPeriod     = 1000
	DefaultServiceReadTimeout     = 30000
)
