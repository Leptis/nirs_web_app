package server

import "github.com/sirupsen/logrus"

type Server struct {
	config *Config
	logger *logrus.Logger
}


func New(config *Config) *Server {
	return &Server {
		config: config,
	}
}

func (s *Server) Start() error {
	return nil
}

