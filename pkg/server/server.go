package server

import (
	"github.com/MotyaSS/DB_CW/pkg/handler"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	router *gin.Engine
}

func NewServer() *Server {
	return &Server{
		router: handler.InitRouter(),
	}
}

func (s *Server) Run(port string) error {
	s.port = port
	return s.router.Run(port)
}
