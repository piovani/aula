package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/piovani/aula/infra/config"
	"github.com/piovani/aula/infra/database"
)

type Service struct {
	Engine *gin.Engine

	Database *database.Database
}

func NewService(db *database.Database) *Service {
	return &Service{
		Engine:   gin.Default(),
		Database: db,
	}
}

func (s *Service) Start() error {
	s.GetRoutes()
	return s.Engine.Run(fmt.Sprintf(":%d", config.Env.ApiPort))
}
