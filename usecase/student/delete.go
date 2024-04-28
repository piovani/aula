package student

import (
	"errors"

	"github.com/google/uuid"
	"github.com/piovani/aula/entites/shared"
)

func (su *StudentUsecase) Delete(id uuid.UUID) error {
	student, err := su.findStudent(id)
	if err != nil {
		return err
	}

	if student.ID == shared.GetUuidEmpty() {
		return errors.New("Não foi possivel encontrar o estudante")
	}

	return su.deleteStudent(id)
}
