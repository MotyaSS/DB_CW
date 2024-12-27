package server

import (
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func New(addr string, timeout time.Duration, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         addr,
			ReadTimeout:  timeout,
			WriteTimeout: timeout,
			Handler:      handler,
		},
	}
}

func (s *Server) Run() {
	slog.Info("Server started at", "address", s.httpServer.Addr)
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
