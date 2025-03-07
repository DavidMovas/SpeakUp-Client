package config

var (
	AppConfig = &Config{
		APIServerTCPURL: "127.0.0.1:8081",
	}
)

type Config struct {
	APIServerTCPURL string
}
