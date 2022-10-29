package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piovani/aula/api/controller"
	"github.com/piovani/aula/entites/shared"
)

func (sc *StudentController) Delete(c *gin.Context) {
	var input Input
	var err error

	input.ID = c.Params.ByName("id")
	input.UUID, err = shared.GetUuidByString(input.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Problema com seu id"))
		return
	}

	if err = sc.StudentUsecase.Delete(input.UUID); err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, controller.NewResponseMessage("Estudante exluido com sucesso"))
}
