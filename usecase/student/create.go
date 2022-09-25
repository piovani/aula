package student

import (
	"github.com/piovani/aula/entites"
)

func (su *StudentUsecase) Create(fullName string, age int) (entites.Student, error) {
	student := entites.NewStudent(fullName, age)

	err := su.Database.StudentRepository.Create(student)

	return *student, err
}
