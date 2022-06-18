package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/piovani/aula/shared"
)

type Student struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"full_name"`
	Age      int       `json:"age"`
}

var Students = []Student{
	Student{ID: shared.GetUuid(), FullName: "Joao", Age: 18},
	Student{ID: shared.GetUuid(), FullName: "Gabriel", Age: 19},
}

func routeHearth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func routeGetStudents(c *gin.Context) {
	c.JSON(http.StatusOK, Students)
}

func routePosttudent(c *gin.Context) {
	var student Student

	err := c.Bind(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message_error": "Não foi possivel obter o payload",
		})
		return
	}

	student.ID = shared.GetUuid()
	Students = append(Students, student)

	c.JSON(http.StatusCreated, student)
}

func routePuttudent(c *gin.Context) {
	var studentPayload Student
	var studentLocal Student
	var newStudents []Student

	err := c.BindJSON(&studentPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message_error": "Não foi possivel obter o payload",
		})
		return
	}

	idString := c.Params.ByName("id")
	id, err := shared.GetUuidByString(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message_error": "Problema com seu id",
		})
		return
	}

	for _, studentElemt := range Students {
		if studentElemt.ID == id {
			studentLocal = studentElemt
		}
	}

	if studentLocal.ID == shared.GetUuidEmpty() {
		c.JSON(http.StatusBadRequest, gin.H{
			"message_error": "Não foi possivel encontrar o estudante",
		})
		return
	}

	studentLocal.FullName = studentPayload.FullName
	studentLocal.Age = studentPayload.Age

	for _, studentElement := range Students {
		if id == studentElement.ID {
			newStudents = append(newStudents, studentLocal)
		} else {
			newStudents = append(newStudents, studentElement)
		}
	}

	Students = newStudents

	c.JSON(http.StatusOK, studentLocal)
}

func routeDeletetudent(c *gin.Context) {
	var newStudents []Student

	idString := c.Params.ByName("id")
	id, err := shared.GetUuidByString(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message_error": "Problema com seu id",
		})
		return
	}

	for _, studentElement := range Students {
		if studentElement.ID != id {
			newStudents = append(newStudents, studentElement)
		}
	}

	Students = newStudents

	c.JSON(http.StatusOK, gin.H{
		"message": "Estudante exluido com sucesso",
	})
}

func routegettudent(c *gin.Context) {
	var student Student

	idString := c.Params.ByName("id")
	id, err := shared.GetUuidByString(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message_error": "Problema com seu id",
		})
		return
	}

	for _, studentElement := range Students {
		if studentElement.ID == id {
			student = studentElement
		}
	}

	if student.ID == shared.GetUuidEmpty() {
		c.JSON(http.StatusBadRequest, gin.H{
			"message_error": "Não foi possivel encontrar o estudante",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

func main() {
	service := gin.Default()

	getRoutes(service)

	service.Run()
}

func getRoutes(c *gin.Engine) *gin.Engine {
	c.GET("/heart", routeHearth)

	groupStudents := c.Group("/students")
	groupStudents.GET("/", routeGetStudents)
	groupStudents.POST("/", routePosttudent)
	groupStudents.PUT("/:id", routePuttudent)
	groupStudents.DELETE("/:id", routeDeletetudent)
	groupStudents.GET("/:id", routegettudent)

	return c
}
