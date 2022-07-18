package student

import (
	"errors"

	"github.com/google/uuid"
	"github.com/piovani/aula/entites"
	"github.com/piovani/aula/entites/shared"
)

func Update(id uuid.UUID, fullName string, age int) (student entites.Student, err error) {
	var newStudents []entites.Student

	for _, studentElemt := range entites.Students {
		if studentElemt.ID == id {
			student = studentElemt
		}
	}

	if student.ID == shared.GetUuidEmpty() {
		return student, errors.New("Não foi possivel encontrar o estudante")
	}

	student.FullName = fullName
	student.Age = age

	for _, studentElement := range entites.Students {
		if student.ID == studentElement.ID {
			newStudents = append(newStudents, student)
		} else {
			newStudents = append(newStudents, studentElement)
		}
	}

	entites.Students = newStudents

	return student, err
}
