package config

type JWT struct {
	Secret string `env:"WEBSOCKET_JWT_SECRET"`
}
