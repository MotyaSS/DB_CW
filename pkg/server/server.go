package server

import (
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		slog.Info("Interrupted by Ctrl+C")
		if err := s.Shutdown(); err != nil {
			slog.Error("Server Closed:", "msg", err.Error())
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
