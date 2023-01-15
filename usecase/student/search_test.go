package student

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/piovani/aula/entites"
	"github.com/stretchr/testify/assert"
)

func (s *StudentUsecaseSuite) TestSearchShoudReturnSuccess() {
	id := uuid.New()

	s.StudentRepository.EXPECT().FindByID(id).Return(entites.Student{ID: id}, nil)

	student, err := s.StudentUsecase.SearchByID(id)

	assert.NotEmpty(s.T(), student)
	assert.EqualValues(s.T(), student.ID, id)
	assert.Nil(s.T(), err)
}

func (s *StudentUsecaseSuite) TestSearchShoudReturnErr() {
	id := uuid.New()
	errExpect := fmt.Errorf("error expected")

	s.StudentRepository.EXPECT().FindByID(id).Return(entites.Student{ID: id}, errExpect)

	student, err := s.StudentUsecase.SearchByID(id)

	assert.NotEmpty(s.T(), student)
	assert.EqualValues(s.T(), student.ID, id)
	assert.Error(s.T(), err, errExpect.Error())
}
