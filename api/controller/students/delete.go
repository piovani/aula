package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piovani/aula/api/controller"
	"github.com/piovani/aula/entites"
	"github.com/piovani/aula/entites/shared"
)

func Delete(c *gin.Context) {
	var input Input
	var newStudents []entites.Student
	var err error

	input.ID = c.Params.ByName("id")
	input.UUID, err = shared.GetUuidByString(input.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Problema com seu id"))
		return
	}

	for _, studentElement := range entites.Students {
		if studentElement.ID != input.UUID {
			newStudents = append(newStudents, studentElement)
		}
	}

	entites.Students = newStudents

	c.JSON(http.StatusOK, controller.NewResponseMessage("Estudante exluido com sucesso"))
}
