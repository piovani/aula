package student

import (
	"github.com/piovani/aula/entites"
)

func Create(fullName string, age int) (student entites.Student, err error) {
	pontStudent := entites.NewStudent(fullName, age)
	student = *pontStudent
	entites.Students = append(entites.Students, student)

	return student, err
}
