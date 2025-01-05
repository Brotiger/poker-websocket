package config

type MongoDB struct {
	Uri              string `env:"WEBSOCKET_MONGODB_URI" envDefault:"mongodb://127.0.0.1:27017/"`
	Username         string `env:"WEBSOCKET_MONGODB_USERNAME"`
	Password         string `env:"WEBSOCKET_MONGODB_PASSWORD"`
	Database         string `env:"WEBSOCKET_MONGODB_DATABASE" envDefault:"poker"`
	ConnectTimeoutMs int    `env:"WEBSOCKET_MONGODB_CONNECT_TIMEOUT_MS" envDefault:"30000"`

	Table struct {
		User         string `env:"WEBSOCKET_MONGODB_TABLE_USER" envDefault:"user"`
		Game         string `env:"WEBSOCKET_MONGODB_TABLE_GAME" envDefault:"game"`
		ConnectToken string `env:"WEBSOCKET_MONGODB_TABLE_CONNECT_TOKEN" envDefault:"connect_token"`
	}
}
