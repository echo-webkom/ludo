package main

import (
	"net/http"

	"github.com/echo-webkom/ludo/board/config"
	"github.com/echo-webkom/ludo/board/server"
)

func main() {
	config := config.Load()
	_ = server.New(config)

	http.ListenAndServe(config.Port, http.FileServer(http.Dir("web")))
}
