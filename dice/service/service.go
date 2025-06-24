package service

import (
	"github.com/echo-webkom/ludo/dice/config"
	"github.com/echo-webkom/ludo/dice/database"
)

type Service struct {
	config *config.Config
	db     *database.TursoDB
}

func New(config *config.Config) *Service {
	return &Service{
		config,
		database.NewTursoDB(config),
	}
}
