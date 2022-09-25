package api

import (
	infra_controller "github.com/piovani/aula/api/controller/infra"
)

func (s *Service) GetRoutes() {
	s.Engine.GET("/heart", infra_controller.Heart)

	groupStudent := s.Engine.Group("students")
	groupStudent.GET("/", s.StudentController.List)
	groupStudent.POST("/", s.StudentController.Create)
	groupStudent.PUT("/:id", s.StudentController.Update)
	groupStudent.DELETE("/:id", s.StudentController.Delete)
	groupStudent.GET("/:id", s.StudentController.Details)
}
