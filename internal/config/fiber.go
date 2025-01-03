package config

type Fiber struct {
	Listen                string `env:"CORE_API_FIBER_LISTEN" envDefault:":8080"`
	RequestTimeoutMs      int    `env:"CORE_API_FIBER_REQUEST_TIMEOUT_MS" envDefault:"3000"`
	DisableStartupMessage bool   `env:"CORE_API_FIBER_DISABLE_STARTUP_MESSAGE" envDefault:"true"`
}
