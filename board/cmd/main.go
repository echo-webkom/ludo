package main

import (
	"github.com/echo-webkom/ludo/board/config"
	"github.com/echo-webkom/ludo/board/server"
)

func main() {
	config := config.Load()
	_ = server.New(config)
}
