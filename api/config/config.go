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
		DatabaseToken:   os.Getenv("TURSO_AUTH_TOKEN"),
		DatabaseURL:     os.Getenv("TURSO_DATABASE_URL"),
		IsDev:           os.Getenv("MODE") == "dev",
		DBFile:          os.Getenv("DB_FILE"),
		GitHubAuthToken: os.Getenv("GITHUB_AUTH_TOKEN"),
	}
}

func toGoPort(port string) string {
	if !strings.HasPrefix(port, ":") {
		return ":" + port
	}
	return port
}
