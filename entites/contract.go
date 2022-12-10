package entites

import "github.com/google/uuid"

type StudentUsecaseContract interface {
	Create(fullName string, age int) (Student, error)
	Delete(id uuid.UUID) error
	List() ([]Student, error)
	SearchByID(id uuid.UUID) (Student, error)
	Update(id uuid.UUID, fullName string, age int) (Student, error)
}
