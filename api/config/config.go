package config

import (
	"log"
	"os"
	"strings"

	"github.com/echo-webkom/cenv"
)

type Config struct {
	Port            string
	DatabaseURL     string
	DatabaseToken   string
	IsDev           bool
	DBFile          string
	GitHubAuthToken string
}

func Load() *Config {
	if err := cenv.Load(); err != nil {
		log.Fatal(err)
	}

	return &Config{
		Port:            toGoPort(os.Getenv("PORT")),
		DatabaseToken:   os.Getenv("DATABASE_TOKEN"),
		DatabaseURL:     os.Getenv("DATABASE_URL"),
		GitHubAuthToken: os.Getenv("GITHUB_AUTH_TOKEN"),
	}
}

func toGoPort(port string) string {
	if !strings.HasPrefix(port, ":") {
		return ":" + port
	}
	return port
}
