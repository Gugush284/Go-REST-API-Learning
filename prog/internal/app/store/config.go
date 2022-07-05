package store

// Config for store ...
type Config struct {
	DatabaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{}
}
