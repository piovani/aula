package students

import "github.com/google/uuid"

type Input struct {
	ID       string
	UUID     uuid.UUID
	FullName string `json:"full_name"`
	Age      int    `json:"age"`
}
