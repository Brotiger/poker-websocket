package config

type App struct {
	GracefulShutdownTimeoutMS int `env:"CORE_API_APP_GRACEFUL_SHUTDOWN_TIMEOUT" envDefault:"10000"`
	Jwt                       struct {
		Secret string `env:"CORE_API_APP_JWT_SECRET"`
	}
}
