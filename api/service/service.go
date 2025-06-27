package service

import (
	"github.com/echo-webkom/ludo/api/config"
	"github.com/echo-webkom/ludo/api/database"
)

type Service struct {
	config *config.Config
	db     database.Database
}

func New(config *config.Config) *Service {
	return &Service{
		config,
		database.NewTursoDB(config),
	}
}
