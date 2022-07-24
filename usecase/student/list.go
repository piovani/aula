package student

import (
	"github.com/piovani/aula/entites"
)

func List() (students []entites.Student, err error) {
	students = entites.Students

	return students, err
}
