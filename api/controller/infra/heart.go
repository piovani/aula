package infra

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piovani/aula/api/controller"
)

func Heart(c *gin.Context) {
	c.JSON(http.StatusOK, controller.NewResponseMessage("ok"))
}
