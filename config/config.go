package config

type Config struct {
	Server     ServerConfig
	Pocketbase PocketbaseConfig
	Chatgpt    ChatgptConfig
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

type ChatgptConfig struct {
	Key         string
	Url         string
	Proxy       string
	Instruction string
}
