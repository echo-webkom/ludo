package server

import (
	"context"
	"log"
	"net/http"

	"github.com/echo-webkom/ludo/api/config"
	"github.com/echo-webkom/ludo/api/database"
	"github.com/echo-webkom/ludo/api/server/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jesperkha/notifier"
)

type Server struct {
	port   string
	router http.Handler
}

func New(config *config.Config, db *database.Database) *Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Mount("/users", routes.UsersHandler(db))
	r.Mount("/items", routes.ItemsHandler(db))
	r.Mount("/boards", routes.BoardsHandler(db))

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

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
