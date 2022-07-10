package apiserver

// Config ...
type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
	SessionKey  string `toml:"session_key"`
}

// New Config
func NewConfig() *Config {
	return &Config{
		BindAddr:   ":8080",
		LogLevel:   "debug",
		SessionKey: "NoKey",
	}
}
