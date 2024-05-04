package student

import (
	"github.com/piovani/aula/entites"
)

func (su *StudentUsecase) Create(fullName string, age int) (entites.Student, error) {
	student := entites.NewStudent(fullName, age)

	err := su.saveDB(student)

	su.saveCache(student)

	return *student, err
}

func (su *StudentUsecase) saveDB(student *entites.Student) error {
	return su.Database.StudentRepository.Create(student)
}
