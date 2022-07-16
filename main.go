package main

import (
	"github.com/piovani/aula/api"
)


// func routegettudent(c *gin.Context) {
// 	var student Student

// 	idString := c.Params.ByName("id")
// 	id, err := shared.GetUuidByString(idString)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message_error": "Problema com seu id",
// 		})
// 		return
// 	}

// 	for _, studentElement := range Students {
// 		if studentElement.ID == id {
// 			student = studentElement
// 		}
// 	}

// 	if student.ID == shared.GetUuidEmpty() {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message_error": "Não foi possivel encontrar o estudante",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, student)
// }

func main() {
	service := api.NewService()
	service.Start()

	// getRoutes(service)
}

// func getRoutes(c *gin.Engine) *gin.Engine {
// 	groupStudents.GET("/:id", routegettudent)

// 	return c
// }
