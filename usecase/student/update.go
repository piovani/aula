package student

import (
	"errors"

	"github.com/google/uuid"
	"github.com/piovani/aula/entites"
	"github.com/piovani/aula/entites/shared"
)

func (su *StudentUsecase) Update(id uuid.UUID, fullName string, age int) (entites.Student, error) {
	student, err := su.Database.StudentRepository.FindByID(id)
	if err != nil {
		return student, err
	}

	if student.ID == shared.GetUuidEmpty() {
		return student, errors.New("Não foi possivel encontrar o estudante")
	}

	student.FullName = fullName
	student.Age = age

	err = su.Database.StudentRepository.Update(&student)

	return student, err
}
