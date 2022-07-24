package api

import (
	infra_controller "github.com/piovani/aula/api/controller/infra"
	student_controller "github.com/piovani/aula/api/controller/students"
)

func (s *Service) GetRoutes() {
	s.Engine.GET("/heart", infra_controller.Heart)

	groupStudent := s.Engine.Group("students")
	groupStudent.GET("/", student_controller.List)
	groupStudent.POST("/", student_controller.Create)
	groupStudent.PUT("/:id", student_controller.Update)
	groupStudent.DELETE("/:id", student_controller.Delete)
	groupStudent.GET("/:id", student_controller.Details)
}
