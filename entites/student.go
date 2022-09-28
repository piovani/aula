package entites

import (
	"github.com/google/uuid"
	"github.com/piovani/aula/entites/shared"
)

type Student struct {
	ID       uuid.UUID `json:"id" bson:"_id"`
	FullName string    `json:"full_name" bson:"full_name"`
	Age      int       `json:"age" bson:"age"`
}

func NewStudent(fullName string, age int) *Student {
	return &Student{
		ID:       shared.GetUuid(),
		FullName: fullName,
		Age:      age,
	}
}

type StudentRepository interface {
	Create(student *Student) error
	List() ([]Student, error)
	FindByID(id uuid.UUID) (Student, error)
	Update(student *Student) error
	Delete(id uuid.UUID) error
}
