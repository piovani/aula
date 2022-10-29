package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/piovani/aula/api/controller/students"
	"github.com/piovani/aula/infra/config"
	"github.com/piovani/aula/infra/database"

	student_usecase "github.com/piovani/aula/usecase/student"
)

type Service struct {
	Engine *gin.Engine

	Database *database.Database

	StudentController *students.StudentController
}

func NewService(db *database.Database) *Service {
	return &Service{
		Engine:   gin.Default(),
		Database: db,
	}
}

func (s *Service) GetControllers() {
	studentUsecase := student_usecase.NewStudentUsecase(s.Database)

	s.StudentController = students.NewStudentController(studentUsecase)
}

func (s *Service) Start() error {
	s.GetControllers()
	s.GetRoutes()

	return s.Engine.Run(fmt.Sprintf(":%d", config.Env.ApiPort))
}
