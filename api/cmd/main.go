package main

import (
	"os"
	"syscall"

	"github.com/echo-webkom/ludo/api/config"
	"github.com/echo-webkom/ludo/api/server"
	"github.com/echo-webkom/ludo/pkg/service"
	"github.com/jesperkha/notifier"
)

func main() {
	config := config.Load()
	notif := notifier.New()
	// db := database.NewTurso(config.TursoURL, config.TursoToken)
	db := service.NewSQLiteService(config.DatabaseURL)
	server := server.New(config, db)

	go server.ListenAndServe(notif)

	notif.NotifyOnSignal(os.Interrupt, syscall.SIGTERM)
}
