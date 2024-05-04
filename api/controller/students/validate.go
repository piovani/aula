package students

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gookit/validate"
	"github.com/piovani/aula/entites"
	"github.com/piovani/aula/entites/shared"
)

func getInputBody(c *gin.Context) (input InputBody, err error) {
	err = c.Bind(&input)
	if err != nil {
		return input, err
	}

	validation := validate.Struct(input)
	if !validation.Validate() {
		return input, validation.Errors
	}

	return input, err
}

func getInputID(c *gin.Context) (id uuid.UUID, err error) {
	inputID := c.Params.ByName("id")

	id, err = shared.GetUuidByString(inputID)
	if err != nil {
		return id, err
	}

	return id, nil
}

func getOutputListStudents(students []entites.Student) (output OutputStudents, err error) {
	for _, s := range students {
		outputStudent, err := getOutputStudent(&s)
		if err != nil {
			return output, err
		}

		output.Students = append(output.Students, outputStudent)
	}

	return output, err
}

func getOutputStudent(student *entites.Student) (output OutputStudent, err error) {
	return OutputStudent{
		ID:       student.ID,
		FullName: student.FullName,
		Age:      student.Age,
	}, err
}
