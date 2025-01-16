package config

type Config struct {
	Server     ServerConfig
	Pocketbase PocketbaseConfig
}

type ServerConfig struct {
	Debug bool
	Port  int
}

type PocketbaseConfig struct {
	Url      string
	Email    string
	Password string
}
