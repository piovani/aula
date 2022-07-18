package student

import (
	"github.com/google/uuid"
	"github.com/piovani/aula/entites"
)

func Delete(id uuid.UUID) (err error) {
	var newStudents []entites.Student

	for _, studentElement := range entites.Students {
		if studentElement.ID != id {
			newStudents = append(newStudents, studentElement)
		}
	}

	entites.Students = newStudents

	return err
}
