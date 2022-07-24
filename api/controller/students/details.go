package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piovani/aula/api/controller"
	"github.com/piovani/aula/entites"
	"github.com/piovani/aula/entites/shared"
	student_usecase "github.com/piovani/aula/usecase/student"
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

	studentFound, err = student_usecase.SearchByID(input.UUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, studentFound)
}
