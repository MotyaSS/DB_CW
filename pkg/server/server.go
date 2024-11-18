package server

import (
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/MotyaSS/DB_CW/pkg/config"
)

type Server struct {
	httpServer *http.Server
}

func New(cfg *config.HttpServer, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         cfg.Address,
			ReadTimeout:  cfg.Timeout,
			WriteTimeout: cfg.Timeout,
			Handler:      handler,
		},
	}
}

func (s *Server) Run() {
	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit)
		<-quit
		log.Println("Interrupted by Ctrl+C")
		if err := s.Shutdown(); err != nil {
			log.Fatal("Server Closed:", err)
		}
	}()
	if err := s.httpServer.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			log.Println("Server closed under request")
		} else {
			log.Fatal("Server closed unexpect")
		}
	}
}

func (s *Server) Shutdown() error {
	return s.httpServer.Close()
}
