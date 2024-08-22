package server

import (
	"distributed-systems-chatbot/internal/config"
	"distributed-systems-chatbot/internal/handlers"
	"net/http"
)

type Server struct {
	config *config.Config
}

func New(cfg *config.Config) *Server {
	return &Server{config: cfg}
}

func (s *Server) Start() error {
	s.routes()
	return http.ListenAndServe(s.config.ServerAddress, nil)
}

func (s *Server) routes() {
	http.HandleFunc("/ws", handlers.HandleConnections)
}
