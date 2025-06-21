package main

import (
	"os"
	"syscall"

	"github.com/echo-webkom/ludo/dice/config"
	"github.com/echo-webkom/ludo/dice/server"
	"github.com/echo-webkom/ludo/dice/service"
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
