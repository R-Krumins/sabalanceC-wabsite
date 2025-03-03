package config

type Config struct {
	ServerPort int
}

func LoadConfig() Config {
	return Config{
		ServerPort: 3000,
	}
}
