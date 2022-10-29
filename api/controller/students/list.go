package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piovani/aula/api/controller"
)

func (sc *StudentController) List(c *gin.Context) {
	students, err := sc.StudentUsecase.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
	}

	c.JSON(http.StatusOK, students)
}
