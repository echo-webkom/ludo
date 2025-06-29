package main

import (
	"os"
	"syscall"

	"github.com/echo-webkom/ludo/api/config"
	"github.com/echo-webkom/ludo/api/database"
	"github.com/echo-webkom/ludo/api/server"
	"github.com/jesperkha/notifier"
)

func main() {
	config := config.Load()
	notif := notifier.New()
	// db := database.NewTurso(config.TursoURL, config.TursoToken)
	db := database.NewSQLite(config.DatabaseURL)
	server := server.New(config, db)

	go server.ListenAndServe(notif)

	notif.NotifyOnSignal(os.Interrupt, syscall.SIGTERM)
}
