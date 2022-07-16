package entites

import (
	"github.com/google/uuid"
	"github.com/piovani/aula/entites/shared"
)

type Student struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"full_name"`
	Age      int       `json:"age"`
}

func NewStudent(fullName string, age int) *Student {
	return &Student{
		ID:       shared.GetUuid(),
		FullName: fullName,
		Age:      age,
	}
}

var Students = []Student{
	Student{ID: shared.GetUuid(), FullName: "Joao", Age: 18},
	Student{ID: shared.GetUuid(), FullName: "Gabriel", Age: 19},
}
