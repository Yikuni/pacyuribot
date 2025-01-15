package config

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Debug bool
	Port  int
}
