package student

import (
	"github.com/google/uuid"
	"github.com/piovani/aula/entites"
)

func (su *StudentUsecase) SearchByID(id uuid.UUID) (entites.Student, error) {
	return su.Database.StudentRepository.FindByID(id)
}
