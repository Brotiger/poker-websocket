package config

type App struct {
	GracefulShutdownTimeoutMS int `env:"WEBSOCKET_APP_GRACEFUL_SHUTDOWN_TIMEOUT" envDefault:"10000"`
}
