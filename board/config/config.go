package config

import (
	"log"
	"os"
	"strings"

	"github.com/echo-webkom/cenv"
)

type Config struct {
	Port         string
	ApiUrl       string
	ApiAuthToken string
}

func Load() *Config {
	if err := cenv.Load(); err != nil {
		log.Fatal(err)
	}

	return &Config{
		Port:         toGoPort(os.Getenv("PORT")),
		ApiUrl:       os.Getenv("API_URL"),
		ApiAuthToken: os.Getenv("API_AUTH_TOKEN"),
	}
}

func toGoPort(port string) string {
	if !strings.HasPrefix(port, ":") {
		return ":" + port
	}
	return port
}
