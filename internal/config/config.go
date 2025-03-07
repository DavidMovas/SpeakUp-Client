package config

var (
	AppConfig = &Config{
		APIServerURL: "localhost:8081",
	}
)

type Config struct {
	APIServerURL string
}
