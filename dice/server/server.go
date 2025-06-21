package server

import (
	"context"
	"log"
	"net/http"

	"github.com/echo-webkom/ludo/dice/config"
	"github.com/echo-webkom/ludo/dice/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jesperkha/notifier"
)

type Server struct {
	port   string
	router http.Handler
}

func New(config *config.Config, service *service.Service) *Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ping", pingHandler())

	r.Delete("/{id}", deleteHandler())
	r.Post("/create/{where}", createHandler())
	r.Patch("/move/{where}/{id}", moveHandler())

	return &Server{
		router: r,
		port:   config.Port,
	}
}

func (s *Server) ListenAndServe(notif *notifier.Notifier) {
	done, finish := notif.Register()

	server := &http.Server{
		Handler: s.router,
		Addr:    s.port,
	}

	go func() {
		<-done
		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatal(err)
		}
		finish()
	}()

	log.Println("listening at port " + s.port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
