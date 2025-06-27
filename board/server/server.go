package server

import "github.com/echo-webkom/ludo/board/config"

type Server struct {
}

func New(config *config.Config) *Server {
	return &Server{}
}
