package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piovani/aula/entites"
)

func List(c *gin.Context) {
	c.JSON(http.StatusOK, entites.Students)
}
