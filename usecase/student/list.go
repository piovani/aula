package student

import (
	"github.com/piovani/aula/entites"
)

func (su *StudentUsecase) List() ([]entites.Student, error) {
	return su.Database.StudentRepository.List()
}
