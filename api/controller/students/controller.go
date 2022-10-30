package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piovani/aula/api/controller"
	"github.com/piovani/aula/entites"
)

type StudentController struct {
	StudentUsecase entites.StudentUsecaseContract
}

func NewStudentController(su entites.StudentUsecaseContract) *StudentController {
	return &StudentController{
		StudentUsecase: su,
	}
}

func (sc *StudentController) Create(c *gin.Context) {
	input, err := getInputBody(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	student, err := sc.StudentUsecase.Create(input.FullName, input.Age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, student)
}

func (sc *StudentController) Delete(c *gin.Context) {
	id, err := getInputID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	if err = sc.StudentUsecase.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, controller.NewResponseMessage("Estudante exluido com sucesso"))
}

func (sc *StudentController) Details(c *gin.Context) {
	id, err := getInputID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Problema com seu id"))
		return
	}

	studentFound, err := sc.StudentUsecase.SearchByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	output, err := getOutputStudent(studentFound)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, output)
}

func (sc *StudentController) List(c *gin.Context) {
	students, err := sc.StudentUsecase.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
	}

	output, err := getOutputListStudents(students)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
	}

	c.JSON(http.StatusOK, output)
}

func (sc *StudentController) Update(c *gin.Context) {
	id, err := getInputID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	input, err := getInputBody(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	student, err := sc.StudentUsecase.Update(id, input.FullName, input.Age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	output, err := getOutputStudent(student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, output)
}
