package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/piovani/aula/infra/config"
)

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
	return s.Engine.Run(fmt.Sprintf(":%d", config.Env.ApiPort))
}
