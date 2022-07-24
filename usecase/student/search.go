package student

import (
	"errors"

	"github.com/google/uuid"
	"github.com/piovani/aula/entites"
	"github.com/piovani/aula/entites/shared"
)

func SearchByID(id uuid.UUID) (student entites.Student, err error) {
	for _, studentElement := range entites.Students {
		if studentElement.ID == id {
			student = studentElement
		}
	}

	if student.ID == shared.GetUuidEmpty() {
		return student, errors.New("Não foi possivel encontrar o estudante")
	}

	return student, err
}
