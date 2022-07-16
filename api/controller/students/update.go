package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piovani/aula/api/controller"
	"github.com/piovani/aula/entites"
	"github.com/piovani/aula/entites/shared"
)

func Update(c *gin.Context) {
	var input Input
	var studenFound entites.Student
	var newStudents []entites.Student

	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	input.ID = c.Params.ByName("id")
	input.UUID, err = shared.GetUuidByString(input.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	for _, studentElemt := range entites.Students {
		if studentElemt.ID == input.UUID {
			studenFound = studentElemt
		}
	}

	if studenFound.ID == shared.GetUuidEmpty() {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Não foi possivel encontrar o estudante"))
		return
	}

	studenFound.FullName = input.FullName
	studenFound.Age = input.Age

	for _, studentElement := range entites.Students {
		if studenFound.ID == studentElement.ID {
			newStudents = append(newStudents, studenFound)
		} else {
			newStudents = append(newStudents, studentElement)
		}
	}

	entites.Students = newStudents

	c.JSON(http.StatusOK, studenFound)
}
