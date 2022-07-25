package api

import "github.com/gin-gonic/gin"

type Service struct {
	*gin.Engine
}

func NewService() *Service {
	return &Service{
		gin.Default(),
	}
}

func (s *Service) Start() error {
	s.GetRoutes()
	return s.Engine.Run()
}
