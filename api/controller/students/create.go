package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piovani/aula/api/controller"
	"github.com/piovani/aula/entites"
)

func Create(c *gin.Context) {
	var input Input
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	student := entites.NewStudent(input.FullName, input.Age)
	entites.Students = append(entites.Students, *student)

	c.JSON(http.StatusAccepted, student)
}
