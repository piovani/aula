package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piovani/aula/api/controller"
	"github.com/piovani/aula/entites"
	"github.com/piovani/aula/entites/shared"
)

func Details(c *gin.Context) {
	var input Input
	var err error
	var studentFound entites.Student

	input.ID = c.Params.ByName("id")
	input.UUID, err = shared.GetUuidByString(input.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Problema com seu id"))
		return
	}

	for _, studentElement := range entites.Students {
		if studentElement.ID == input.UUID {
			studentFound = studentElement
		}
	}

	if studentFound.ID == shared.GetUuidEmpty() {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Não foi possivel encontrar o estudante"))
		return
	}

	c.JSON(http.StatusOK, studentFound)
}
