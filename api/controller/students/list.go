package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piovani/aula/api/controller"
	student_usecase "github.com/piovani/aula/usecase/student"
)

func List(c *gin.Context) {
	students, err := student_usecase.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
	}

	c.JSON(http.StatusOK, students)
}
