package main

import (
	"os"
	"syscall"

	"github.com/echo-webkom/ludo/api/config"
	"github.com/echo-webkom/ludo/api/server"
	"github.com/echo-webkom/ludo/api/service"
	"github.com/jesperkha/notifier"
)

func main() {
	config := config.Load()
	notif := notifier.New()
	service := service.New(config)
	server := server.New(config, service)

	go server.ListenAndServe(notif)

	notif.NotifyOnSignal(os.Interrupt, syscall.SIGTERM)
}
